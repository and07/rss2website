package main

import (
	"log"
	"testing"

	proto "github.com/and07/rss2website/proto"
)

func TestForbesData(t *testing.T) {
	log.Println()
	log.Println()
	log.Println()
	log.Println()
	log.Println()
	log.SetFlags(log.Flags() | log.Lshortfile)
	c := Forbes{}
	urlRss := "https://www.forbes.com/real-time/feed2/"
	c.urlRss = urlRss
	c.data = make(map[string]*proto.Post)
	c.exit = make(chan struct{})
	c.getData()
	log.Println()
	log.Println()
	log.Println()
	log.Println()
	log.Println(c.data)

}
