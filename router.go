package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	proto "gitlab.com/and07/rss2website/proto"
)

func Router(rss []IRss, posts map[string]IRss, tpl *Template) *httprouter.Router {

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

	r.GET("/api/v1/post/:rss/:slug", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Println("detail api v1")
		slug := ps.ByName("slug")
		rss := ps.ByName("rss")
		ir := posts[rss].GetRssData()
		tmpl := template.Must(template.New("detail.api.html").Funcs(tpl.Fmap()).ParseFiles("tpl/detail.api.html"))
		log.Println(tmpl)
		tmpl.Execute(w, ir.Pages[slug])
	})
	return r
}
