package main

import (
	"flag"
	"log"
	"time"

	"github.com/Hoi-Kun/letsgo-bongu.com/server"
	"github.com/Hoi-Kun/letsgo-bongu.com/siteconfig"
)

func init() {
	modeDev := flag.Bool("dev", false, "set http://localhost:port if true")
	flag.Parse()
	if *modeDev {
		siteconfig.Config.ModeDev = true
	}
	log.Printf("DEV     %t", *modeDev)
	log.Printf("HOST    %s", siteconfig.Host())

	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}
	time.Local = loc
}

func main() {
	s := server.New(uint32(8006))
	log.Fatal(s.Serve())
}
