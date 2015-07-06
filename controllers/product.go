package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/partyschaum/go-web-app/viewmodels"
)

type productController struct {
	template *template.Template
}

func (p *productController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, err := strconv.Atoi(vars["id"])
	if err == nil {
		vm := viewmodels.GetProduct(id)

		w.Header().Add("Content Type", "text/html")
		p.template.Execute(w, vm)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
	}
}
