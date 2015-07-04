package controllers

import (
	"html/template"
	"net/http"

	"github.com/partyschaum/go-web-app/viewmodels"
)

type homeController struct {
	template *template.Template
}

func (h *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()

	w.Header().Add("Content Type", "text/html")
	h.template.Execute(w, vm)
}
