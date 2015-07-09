package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/partyschaum/go-web-app/controllers/util"
	"github.com/partyschaum/go-web-app/converters"
	"github.com/partyschaum/go-web-app/models"
	"github.com/partyschaum/go-web-app/viewmodels"
)

type categoriesController struct {
	template *template.Template
}

func (c *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	categories := models.GetCategories()

	categoriesVM := []viewmodels.Category{}
	for _, category := range categories {
		categoriesVM = append(categoriesVM, converters.ConvertCategoryToViewModel(category))
	}

	vm := viewmodels.GetCategories()
	vm.Categories = categoriesVM

	w.Header().Add("Content Type", "text/html")

	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	c.template.Execute(responseWriter, vm)
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
