package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/oxtoacart/bpool"
)

type TemplateConfig struct {
	TemplateLayoutPath  string
	TemplateIncludePath string
}

type Template struct {
	templates map[string]*template.Template
	bufpool   *bpool.BufferPool

	mainTmpl string

	templateConfig TemplateConfig
}

func formatAsDate(t int64) string {
	tm := time.Unix(t, 0)
	year, month, day := tm.Date()
	return fmt.Sprintf("%d.%d.%d", day, month, year)
}
func safe(s string) template.HTML {
	return template.HTML(s)
}

func NewTemplate(templateLayoutPath, templateIncludePath, mainTmpl string) *Template {
	tc := TemplateConfig{
		TemplateLayoutPath:  templateLayoutPath,
		TemplateIncludePath: templateIncludePath,
	}

	return &Template{
		mainTmpl:       mainTmpl,
		templateConfig: tc,
	}

}

func (t *Template) Init() {
	t.loadTemplates()
}

func (t *Template) Fmap() template.FuncMap {
	fmap := template.FuncMap{
		"formatAsDate": formatAsDate,
		"safe":         safe,
	}
	return fmap
}

func (t *Template) loadTemplates() {

	if t.templates == nil {
		t.templates = make(map[string]*template.Template)
	}

	layoutFiles, err := filepath.Glob(t.templateConfig.TemplateLayoutPath + "*.html")
	if err != nil {
		log.Fatal(err)
	}

	includeFiles, err := filepath.Glob(t.templateConfig.TemplateIncludePath + "*.html")
	if err != nil {
		log.Fatal(err)
	}

	mainTemplate := template.New("main")

	mainTemplate, err = mainTemplate.Parse(t.mainTmpl)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		t.templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			log.Fatal(err)
		}
		t.templates[fileName] = template.Must(t.templates[fileName].Funcs(t.Fmap()).ParseFiles(files...))
	}

	log.Println("templates loading successful")

	t.bufpool = bpool.NewBufferPool(64)
	log.Println("buffer allocation successful")
}

func (t *Template) RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, ok := t.templates[name]
	if !ok {
		http.Error(w, fmt.Sprintf("The template %s does not exist.", name),
			http.StatusInternalServerError)
	}

	buf := t.bufpool.Get()
	defer t.bufpool.Put(buf)

	err := tmpl.Execute(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}
