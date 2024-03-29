package parser

import (
	"regexp"
	"spider/engine"
	"spider/model"
	"spider_dist/config"
	"strconv"
)

// 一共 8 个
var (
	// nameRe       = regexp.MustCompile(`<h1 data-v-.*="" class="nickName">([^<]+)</h1>`)
	genderRe     = regexp.MustCompile(`<div[^>]*class="m-title"[^>]*>(.)的动态</div>`)
	ageRe        = regexp.MustCompile(`<div[^>]*class="m-btn purple"[^>]*>(\d+)岁</div>`)
	heightRe     = regexp.MustCompile(`<div[^>]*class="m-btn purple"[^>]*>(\d+)cm</div>`)
	incomeRe     = regexp.MustCompile(`<div[^>]*class="m-btn purple"[^>]*>月收入:([^<]+)</div>`)
	marriageRe   = regexp.MustCompile(`<div[^>]*class="m-btn purple"[^>]*>([^<]+)</div>`)
	occupationRe = regexp.MustCompile(`<div[^>]*class="m-btn purple"[^>]*>工作地:([^<]+)</div>`)
	xingzuoRe    = regexp.MustCompile(`<div[^>]*class="m-btn purple"[^>]*>(..)座\(.+-.+\)</div>`)
	idURLRe      = regexp.MustCompile(`album.zhenai.com/u/(\d+)`)
)

func parseProfile(contents []byte, url string, name string) *engine.ParseResult {
	profile := model.Profile{}

	// 姓名(昵称)
	// profile.Name = extract(contents, nameRe)
	profile.Name = name

	// 性别
	ta := extract(contents, genderRe)
	if ta == "她" {
		ta = "女"
	} else {
		ta = "男"
	}
	profile.Gender = ta

	// 年龄
	age, _ := strconv.Atoi(extract(contents, ageRe))
	profile.Age = age

	// 身高
	height, _ := strconv.Atoi(extract(contents, heightRe))
	profile.Height = height

	// 收入
	profile.Income = extract(contents, incomeRe)

	// 婚况
	profile.Marriage = extract(contents, marriageRe)

	// 职业
	profile.Occupation = extract(contents, occupationRe)

	// 星座
	xingzuo := extract(contents, xingzuoRe)
	profile.Xingzuo = xingzuo + "座"

	// URL

	item := &engine.Item{
		URL:     url,
		Id:      extract([]byte(url), idURLRe),
		Type:    "zhenai",
		Payload: &profile,
	}
	// fmt.Println("############################################################")
	// fmt.Printf("URL: %v\n", url)
	// fmt.Printf("Id: %v\n", extract([]byte(url), idURLRe))
	// fmt.Printf("%+v\n", item.Payload)
	// fmt.Println("############################################################")

	return &engine.ParseResult{
		Items: []*engine.Item{item},
	}
}

func extract(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}

	return ""
}

type ProfileParser struct {
	userName string
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}

func (p *ProfileParser) Parse(contents []byte, url string) *engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ProfileParser, p.userName
}
