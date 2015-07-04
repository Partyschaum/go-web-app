package controllers

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func Register(templates *template.Template) {
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		requestedFile := req.URL.Path[1:]
		template := templates.Lookup(requestedFile + ".html")

		log.Printf("Requesting '%s'", requestedFile)

		var context interface{} = nil

		switch requestedFile {
		case "home":
		context = viewmodels.GetHome()
		case "categories":
		context = viewmodels.GetCategories()
		case "products":
		context = viewmodels.GetProducts()
		case "product":
		context = viewmodels.GetProduct()
		}

		if template != nil {
		template.Execute(w, context)
		} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
		}
		})
	*/

	hc := new(homeController)
	hc.template = templates.Lookup("home.html")
	http.HandleFunc("/home", hc.get)

	cc := new(categoriesController)
	cc.template = templates.Lookup("categories.html")
	http.HandleFunc("/categories", cc.get)

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path

	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		log.Printf("Serving asset: %s", path)

		defer f.Close()
		w.Header().Add("Content Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		log.Printf("Missing asset: %s", path)

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
	}
}
