package siteconfig

import (
	"fmt"
	"os"
	"time"
)

type searchEngineConnection struct {
	Google string
}

type store struct {
	Location string
	Category string
	Name     string
	Address  string
	Tel      string
	Images   []string
}

func parseImages() []string {
	var ss []string
	entry, err := os.ReadDir("./static/img/index/store")
	if err != nil {
		panic(err)
	}
	for _, e := range entry {
		ss = append(ss, e.Name())
	}
	return ss
}

type config struct {
	ModeDev bool

	Port   uint32
	Domain string

	DatePublished time.Time
	DateModified  time.Time

	Title       string
	Description string
	Keywords    string
	Author      string

	Thumbnail string

	Store *store

	SearchEngineConnection *searchEngineConnection
}

var Config *config = &config{
	ModeDev: false,

	Port:   uint32(8006),
	Domain: "letsgo-bongu.com",

	DatePublished: time.Date(2024, time.Month(5), 14, 0, 0, 0, 0, time.Local),
	DateModified:  time.Date(2024, time.Month(5), 15, 0, 0, 0, 0, time.Local),

	Title:       "가즈아",
	Description: "채우자",
	Keywords:    "봉구실장, 가즈아 가라오케, 가즈아 봉구실장, 가라오케 봉구실장, 가라오케, 레츠고 봉구실장, 레츠고 가라오케",
	Author:      "봉구실장",

	Thumbnail: "/static/img/index/thumbnail.png",

	Store: &store{
		Location: "홍대",
		Category: "가라오케",
		Name:     "가즈아",
		Address:  "서울 마포구 양화로 지하 160",
		Tel:      "010-7676-3421",
		Images:   parseImages(),
	},
	SearchEngineConnection: &searchEngineConnection{
		Google: "vca6lAnDIG-285kKpSxmm47sVeAOO5U-IM-_fvTAsaU",
	},
}

func Host() string {
	if Config.ModeDev {
		return fmt.Sprintf("http://localhost:%d", Config.Port)
	}
	return fmt.Sprintf("https://%s", Config.Domain)
}
