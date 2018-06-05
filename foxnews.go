package main

import (
	"log"

	proto "github.com/and07/rss2website/proto"
	"github.com/gosimple/slug"
	"github.com/mmcdole/gofeed"
)

type Foxnews struct {
}

func (fn *Foxnews) Name() string {
	return "foxnews"
}

func (fn *Foxnews) GetRssData(urlRss string) proto.PostPageData {
	fp := gofeed.NewParser()

	var posts proto.PostPageData

	feed, _ := fp.ParseURL(urlRss)
	log.Println(feed.Title)
	log.Println(feed.Image.URL)

	pages := make(map[string]*proto.Post)
	for _, v := range feed.Items {
		if v.Extensions["media"]["group"] != nil {
			log.Printf("%#v \n", v.Extensions["media"]["group"][0].Children["content"][0].Attrs["url"])

			pages[slug.Make(v.Title)] = &proto.Post{
				Title:       v.Title,
				Slug:        slug.Make(v.Title),
				Link:        v.Link,
				Description: v.Description,
				Image:       v.Extensions["media"]["group"][0].Children["content"][0].Attrs["url"],
				SourceImage: "//global.fncstatic.com/static/orion/styles/img/fox-news/favicons/apple-touch-icon-60x60.png",
			}
		}

	}
	data := proto.PostPageData{
		PageTitle: feed.Title,
		PageImage: feed.Image.URL,
		Pages:     pages,
	}
	posts = data

	return posts
}
