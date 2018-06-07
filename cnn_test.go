package main

import (
	"log"
	"testing"

	proto "github.com/and07/rss2website/proto"
)

func TestCnnData(t *testing.T) {
	log.SetFlags(log.Flags() | log.Lshortfile)
	c := Cnn{}
	urlRss := "http://rss.cnn.com/rss/money_topstories.rss"
	c.urlRss = urlRss
	c.data = make(map[string]*proto.Post)
	c.exit = make(chan struct{})
	c.getData()

}
