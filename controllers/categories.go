package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/partyschaum/go-web-app/controllers/util"
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

type categoryController struct {
	template *template.Template
}

func (c *categoryController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, err := strconv.Atoi(vars["id"])

	if err == nil {
		vm := viewmodels.GetProducts(id)

		w.Header().Add("Content Type", "text/html")

		gzw := util.GetResponseWriter(w, req)
		defer gzw.Close()

		c.template.Execute(gzw, vm)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
	}

}
