package controllers

import (
	"html/template"
	"net/http"

	"github.com/partyschaum/go-web-app/viewmodels"
)

type categoriesController struct {
	template *template.Template
}

func (c *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetCategories()

	w.Header().Add("Content Type", "text/html")
	c.template.Execute(w, vm)
}
