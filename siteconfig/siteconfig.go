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

	Title:       "채우자",
	Description: "채우자",
	Keywords:    "채우자",
	Author:      "채우자",

	Thumbnail: "/static/img/index/thumbnail.png",

	Store: &store{
		Location: "채우자",
		Category: "채우자",
		Name:     "채우자",
		Address:  "채우자",
		Tel:      "010.7676.3421",
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
