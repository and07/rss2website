package main

import (
	"log"
	"testing"

	proto "gitlab.com/and07/rss2website/proto"
)

func TestFoxnewsData(t *testing.T) {
	log.SetFlags(log.Flags() | log.Lshortfile)
	f := Foxnews{}
	urlRss := "http://feeds.foxnews.com/foxnews/latest"
	f.urlRss = urlRss
	f.data = make(map[string]*proto.Post)
	f.exit = make(chan struct{})
	f.getData()

}
