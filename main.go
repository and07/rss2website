package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
	"github.com/mmcdole/gofeed"
	"github.com/oxtoacart/bpool"
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

type TemplateConfig struct {
	TemplateLayoutPath  string
	TemplateIncludePath string
}

var templates map[string]*template.Template
var bufpool *bpool.BufferPool

var mainTmpl = `{{define "main" }} {{ template "base" . }} {{ end }}`

var templateConfig TemplateConfig

func loadConfiguration() {
	templateConfig.TemplateLayoutPath = "tpl/layouts/"
	templateConfig.TemplateIncludePath = "tpl/"
}

func loadTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layoutFiles, err := filepath.Glob(templateConfig.TemplateLayoutPath + "*.html")
	if err != nil {
		log.Fatal(err)
	}

	includeFiles, err := filepath.Glob(templateConfig.TemplateIncludePath + "*.html")
	if err != nil {
		log.Fatal(err)
	}

	mainTemplate := template.New("main")

	mainTemplate, err = mainTemplate.Parse(mainTmpl)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			log.Fatal(err)
		}
		templates[fileName] = template.Must(templates[fileName].ParseFiles(files...))
	}

	log.Println("templates loading successful")

	bufpool = bpool.NewBufferPool(64)
	log.Println("buffer allocation successful")
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, fmt.Sprintf("The template %s does not exist.", name),
			http.StatusInternalServerError)
	}

	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := tmpl.Execute(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
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

	loadConfiguration()
	loadTemplates()

	fp := gofeed.NewParser()

	var posts PostPageData
	for _, v := range getRSSList() {
		log.Println(v.Rss)
		feed, _ := fp.ParseURL(v.Rss)
		log.Println(feed.Title)
		log.Println(feed.Image.URL)

		pages := make(map[string]Post)
		for _, v := range feed.Items {
			if v.Extensions["media"]["group"] != nil {
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

	r := httprouter.New()

	r.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		renderTemplate(w, "list.html", posts)
	})

	r.GET("/:slug", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		slug := ps.ByName("slug")

		renderTemplate(w, "detail.html", posts.Pages[slug])
	})

	http.ListenAndServe(":3000", r)

}
