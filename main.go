package main

import (
	"log"
	"net/http"

	proto "github.com/and07/rss2website/proto"
	"github.com/julienschmidt/httprouter"
)

type IRss interface {
	GetRssData(string) proto.PostPageData
	Name() string
}

type Rss []IRss

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)

	tpl := NewTemplate("tpl/layouts/", "tpl/", `{{define "main" }} {{ template "base" . }} {{ end }}`)
	tpl.Init()

	foxnews := Foxnews{}
	rss := []IRss{&foxnews}
	rssList := getRSSList()
	list := rssList.GetList()
	posts := make(map[string]proto.PostPageData)
	for _, v := range rss {
		log.Println(list[v.Name()])
		posts[v.Name()] = v.GetRssData(list[v.Name()])
		///fmt.Print("%#v \n", feed)
	}

	log.Println()
	log.Println()
	log.Println()

	r := httprouter.New()

	r.ServeFiles("/static/*filepath", http.Dir("static"))

	r.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		tpl.RenderTemplate(w, "list.html", posts)
	})

	r.GET("/post/:rss/:slug", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		slug := ps.ByName("slug")
		rss := ps.ByName("rss")

		tpl.RenderTemplate(w, "detail.html", posts[rss].Pages[slug])
	})

	http.ListenAndServe(":3000", r)

}
