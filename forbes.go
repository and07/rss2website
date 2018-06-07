package main

import (
	"log"
	"sync"
	"time"

	proto "github.com/and07/rss2website/proto"
	"github.com/gosimple/slug"
	"github.com/mmcdole/gofeed"
)

type Forbes struct {
	ticker *time.Ticker
	mx     sync.RWMutex
	exit   chan struct{}
	data   map[string]*proto.Post
	urlRss string
	title  string
	image  string
}

func NewForbes(list map[string]string) *Forbes {

	f := Forbes{}
	urlRss := list[f.Name()]
	f.urlRss = urlRss
	f.data = make(map[string]*proto.Post)
	f.exit = make(chan struct{})
	f.getData()
	f.Start()
	return &f
}

func (c *Forbes) Start() {

	c.ticker = time.NewTicker(30 * time.Minute)
	go func() {

		for {
			select {
			case <-c.ticker.C:
				go func() {
					c.mx.Lock()
					defer c.mx.Unlock()
					defer log.Println("Tick")

					c.getData()

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

func (c *Forbes) Closed() {
	c.exit <- struct{}{}
	close(c.exit)
}

func (c *Forbes) Name() string {
	return "forbes"
}

func (c *Forbes) GetRssData() proto.PostPageData {
	c.mx.RLock()
	defer c.mx.RUnlock()

	return proto.PostPageData{
		PageTitle: c.title,
		PageImage: c.image,
		Pages:     c.data,
	}
}

func (c *Forbes) getData() {
	fp := gofeed.NewParser()

	feed, errFeed := fp.ParseURL(c.urlRss)
	if errFeed != nil {
		log.Println(errFeed)
		return
	}
	log.Println(feed.Title)

	c.title = feed.Title

	log.Printf("%#v \n", feed.Items[0])

	for _, v := range feed.Items {

		if v.Extensions["media"]["content"] != nil {
			//log.Printf("%#v \n", v.Extensions["media"]["thumbnail"][0].Attrs["url"])

			if _, ok := c.data[slug.Make(v.Title)]; !ok {

				//t1, _ := time.Parse(time.RFC1123Z, v.Published)

				c.data[slug.Make(v.Title)] = &proto.Post{
					Published:   v.PublishedParsed.Unix(),
					Categories:  v.Categories,
					Title:       v.Title,
					Slug:        slug.Make(v.Title),
					Link:        v.Link,
					Description: v.Description,
					Image:       v.Extensions["media"]["content"][0].Attrs["url"],
					SourceImage: "//cdn2.mhpbooks.com/2016/12/forbes_1200x1200-235x235.jpg",
				}

			}

		}

	}

}
