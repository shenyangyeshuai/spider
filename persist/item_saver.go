package persist

import (
	"log"
)

func ItemSaver() chan interface{} {
	ch := make(chan interface{})

	go func() {
		itemCount := 0
		for {
			item := <-ch
			log.Printf("Got item#%d: %+v", itemCount, item)
			itemCount++
		}
	}()

	return ch
}
