package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/partyschaum/go-web-app/controllers"
)

func main() {
	templates := populateTemplates()

	controllers.Register(templates)

	http.ListenAndServe(":8000", nil)
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
