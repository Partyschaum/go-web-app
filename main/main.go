package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/partyschaum/go-web-app/viewmodels"
)

func main() {
	templates := populateTemplates()

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
		}

		if template != nil {
			template.Execute(w, context)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(http.StatusText(http.StatusNotFound)))
		}
	})

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)

	http.ListenAndServe(":8000", nil)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path

	log.Printf("Serving asset: %s", path)

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
		defer f.Close()
		w.Header().Add("Content Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
	}
}

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatesPathsRaw, _ := templateFolder.Readdir(-1)

	templatesPaths := new([]string)
	for _, pathInfo := range templatesPathsRaw {
		if !pathInfo.IsDir() && strings.HasSuffix(pathInfo.Name(), ".html") {
			log.Printf("Loading template %s", pathInfo.Name())
			*templatesPaths = append(*templatesPaths, basePath+"/"+pathInfo.Name())
		}
	}

	result.ParseFiles(*templatesPaths...)

	return result
}
