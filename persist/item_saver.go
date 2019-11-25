package persist

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"spider/engine"
)

func ItemSaver(index string) (chan *engine.Item, error) {
	ch := make(chan *engine.Item)

	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}

	go func() {
		itemCount := 0
		for {
			item := <-ch
			log.Printf("Got item#%d: %+v", itemCount, item)
			itemCount++

			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item: %v: %v", item, err)
			}
		}
	}()

	return ch, nil
}

func Save(client *elastic.Client, index string, item *engine.Item) error {
	if item.Type == "" {
		return fmt.Errorf("Must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
