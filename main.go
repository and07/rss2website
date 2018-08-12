package main

import (
	"log"
	"net/http"

	proto "gitlab.com/and07/rss2website/proto"
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

	http.ListenAndServe(":3000", Router(rss, posts, tpl))

}
