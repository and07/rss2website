package main

import (
	"log"
	"net/http"

	proto "github.com/and07/rss2website/proto"
	"github.com/julienschmidt/httprouter"
)

type IRss interface {
	GetRssData() proto.PostPageData
	Name() string
}

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)

	tpl := NewTemplate("tpl/layouts/", "tpl/", `{{define "main" }} {{ template "base" . }} {{ end }}`)
	tpl.Init()

	rssList := getRSSList()
	list := rssList.GetList()

	foxnews := NewFoxnews(list)
	defer foxnews.Closed()

	cnn := NewCnn(list)
	defer cnn.Closed()

	forbes := NewForbes(list)
	defer forbes.Closed()

	rss := []IRss{foxnews, cnn, forbes}

	posts := make(map[string]IRss)
	for _, v := range rss {
		log.Println(list[v.Name()])
		posts[v.Name()] = v
		///fmt.Print("%#v \n", feed)
	}

	log.Println()
	log.Println()
	log.Println()

	r := httprouter.New()

	r.ServeFiles("/static/*filepath", http.Dir("static"))

	r.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		log.Println("list")
		p := make(map[string]proto.PostPageData)
		for _, v := range rss {
			//log.Println(list[v.Name()])
			p[v.Name()] = v.GetRssData()
			///fmt.Print("%#v \n", feed)
		}

		tpl.RenderTemplate(w, "list.html", p)
	})

	r.GET("/post/:rss/:slug", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Println("detail")
		slug := ps.ByName("slug")
		rss := ps.ByName("rss")
		ir := posts[rss].GetRssData()

		tpl.RenderTemplate(w, "detail.html", ir.Pages[slug])
	})

	http.ListenAndServe(":3000", r)

}
