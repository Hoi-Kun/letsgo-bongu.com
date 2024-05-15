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

	Title:       "홍대 가라오케 가즈아",
	Description: "홍대 가라오케 가즈아에서는 홍대 지역의 특색을 반영한 역동적이고 즐거운 분위기에서 친구들과 함께 신나는 노래방 경험을 즐길 수 있습니다. 가즈아는 젊은 층에게 인기 있는 장소로, 최신 곡들과 고품질의 사운드 시스템을 갖추고 있어 노래를 부르며 즐거운 시간을 보낼 수 있습니다. 활기찬 홍대의 밤문화를 대표하는 이곳에서 매력적인 음악과 함께 잊지 못할 추억을 만들어 보세요.",
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
