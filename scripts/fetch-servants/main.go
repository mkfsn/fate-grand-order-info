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

	staticImages := []dumper.Option{
		dumper.WithImage("https://grandorder.wiki/images/thumb/f/f2/Icon_Class_Beast_Gold.png/40px-Icon_Class_Beast_Gold.png"),
		dumper.WithImage("https://grandorder.wiki/images/thumb/f/f2/Icon_Class_Beast_Gold.png/60px-Icon_Class_Beast_Gold.png"),
		dumper.WithImage("https://grandorder.wiki/images/f/f2/Icon_Class_Beast_Gold.png"),
		dumper.WithImage("https://grandorder.wiki/images/thumb/7/7c/Icon_Class_Shielder_Gold.png/40px-Icon_Class_Shielder_Gold.png"),
		dumper.WithImage("https://grandorder.wiki/images/thumb/7/7c/Icon_Class_Shielder_Gold.png/60px-Icon_Class_Shielder_Gold.png"),
		dumper.WithImage("https://grandorder.wiki/images/7/7c/Icon_Class_Shielder_Gold.png"),
	}
	if err := dumper.New(res, staticImages...).Dump(); err != nil {
		log.Fatalln(err)
	}
}
