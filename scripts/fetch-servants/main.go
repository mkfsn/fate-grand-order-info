package main

import (
	"log"

	"fetch-servant/crawler"
	"fetch-servant/dumper"
	"fetch-servant/parser"
	"fetch-servant/processer"
)

func main() {
	b, err := crawler.New("https://grandorder.wiki/Servant_List").Crawl()
	if err != nil {
		log.Fatalln(err)
	}

	servants, commands, err := parser.New(b).Parse()
	if err != nil {
		log.Fatalln(err)
	}

	res, err := processer.New(servants, commands).Process()
	if err != nil {
		log.Fatalln(err)
	}

	if err := dumper.New(res).Dump(); err != nil {
		log.Fatalln(err)
	}
}
