package main

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gosimple/slug"
	"github.com/mmcdole/gofeed"
	proto "gitlab.com/and07/rss2website/proto"
)

type Cnn struct {
	ticker *time.Ticker
	mx     sync.RWMutex
	exit   chan struct{}
	data   map[string]*proto.Post
	urlRss string
	title  string
	image  string
}

func NewCnn(list map[string]string) *Cnn {

	c := Cnn{}
	urlRss := list[c.Name()]
	c.urlRss = urlRss
	c.data = make(map[string]*proto.Post)
	c.exit = make(chan struct{})
	c.data = c.getData()
	c.Start()
	return &c
}

func (c *Cnn) Start() {

	c.ticker = time.NewTicker(30 * time.Minute)
	go func() {

		for {
			select {
			case <-c.ticker.C:
				go func() {
					c.mx.Lock()
					defer c.mx.Unlock()
					defer log.Println("Tick")

					c.data = c.getData()

					//log.Printf("coin pair = %#v\n", marketName)
					//log.Printf("book = %#v\n", book)
					//log.Printf("history = %#v\n", history)
				}()
			case <-c.exit:
				log.Println("Exit")
				c.ticker.Stop()
				return
			}
		}
	}()
}

func (c *Cnn) Closed() {
	c.exit <- struct{}{}
	close(c.exit)
}

func (c *Cnn) Name() string {
	return "cnn"
}

func (c *Cnn) GetRssData() proto.PostPageData {
	c.mx.RLock()
	defer c.mx.RUnlock()

	return proto.PostPageData{
		PageTitle: c.title,
		PageImage: c.image,
		Pages:     c.data,
	}
}

func (c *Cnn) getData() map[string]*proto.Post {
	fp := gofeed.NewParser()

	data := make(map[string]*proto.Post)

	feed, errFeed := fp.ParseURL(c.urlRss)
	if errFeed != nil {
		log.Println(errFeed)
		return data
	}
	log.Println(feed.Title)
	log.Println(feed.Image.URL)

	c.title = feed.Title
	c.image = feed.Image.URL

	for _, v := range feed.Items {

		if v.Extensions["media"]["thumbnail"] != nil {
			//log.Printf("%#v \n", v.Extensions["media"]["thumbnail"][0].Attrs["url"])

			if _, ok := data[slug.Make(v.Title)]; !ok {

				t1, _ := time.Parse(time.RFC1123, v.Published)

				img := strings.Replace(v.Extensions["media"]["thumbnail"][0].Attrs["url"], "http://", "//", -1)

				data[slug.Make(v.Title)] = &proto.Post{
					Published:   t1.Unix(),
					Categories:  v.Categories,
					Title:       v.Title,
					Slug:        slug.Make(v.Title),
					Link:        strings.Replace(v.Link, "http://", "//", -1),
					Description: strings.Replace(v.Description, "http://", "//", -1),
					Image:       img,
					SourceImage: "//i.cdn.turner.com/money/.element/cnnm-3.0/img/logo/cnnmoney_blue.svg",
					SourceTitle: feed.Title,
				}

			}

		}

	}
	return data

}
