package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
	"github.com/mmcdole/gofeed"
)

type RssList []struct {
	URL string `json:"url"`
	Rss string `json:"rss"`
}

type Post struct {
	Title       string
	Slug        string
	Description string
	Link        string
	Image       string
	SourceImage string
	Done        bool
}

type PostPageData struct {
	PageTitle string
	PageImage string
	Pages     map[string]Post
}

func getRSSList() RssList {
	//https://gist.github.com/stungeye/fe88fc810651174d0d180a95d79a8d97
	j := `[
		{
			"url": "http://www.foxnews.com",
			"rss": "http://feeds.foxnews.com/foxnews/latest"
		}
	]`

	var list RssList
	if err := json.Unmarshal([]byte(j), &list); err != nil {
		log.Fatal(err)
	}
	return list

}

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)

	fp := gofeed.NewParser()

	var posts PostPageData
	for _, v := range getRSSList() {
		log.Println(v.Rss)
		feed, _ := fp.ParseURL(v.Rss)
		log.Println(feed.Title)
		log.Println(feed.Image.URL)

		pages := make(map[string]Post)
		for _, v := range feed.Items {
			log.Printf("%#v \n", v.Extensions["media"]["group"][0].Children["content"][0].Attrs["url"])

			pages[slug.Make(v.Title)] = Post{
				Title:       v.Title,
				Slug:        slug.Make(v.Title),
				Link:        v.Link,
				Description: v.Description,
				Image:       v.Extensions["media"]["group"][0].Children["content"][0].Attrs["url"],
				SourceImage: "//global.fncstatic.com/static/orion/styles/img/fox-news/favicons/apple-touch-icon-60x60.png",
				Done:        false,
			}
		}
		data := PostPageData{
			PageTitle: feed.Title,
			PageImage: feed.Image.URL,
			Pages:     pages,
		}
		posts = data
		///fmt.Print("%#v \n", feed)
	}

	log.Println()
	log.Println()
	log.Println()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("tpl/layout.html"))
		tmpl.Execute(w, posts)
	})

	r.HandleFunc("/{slug}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		slug := vars["slug"]

		tmpl := template.Must(template.ParseFiles("tpl/detail.html"))
		tmpl.Execute(w, posts.Pages[slug])
	})

	http.ListenAndServe(":3000", r)

}
