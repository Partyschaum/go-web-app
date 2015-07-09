package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/partyschaum/go-web-app/controllers/util"
	"github.com/partyschaum/go-web-app/viewmodels"
)

type standLocatorController struct {
	template *template.Template
}

func (s *standLocatorController) get(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vm := viewmodels.GetStandLocator()
	responseWriter.Header().Add("Content Type", "text/html")

	s.template.Execute(responseWriter, vm)
}

func (s *standLocatorController) apiSearch(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vm := viewmodels.GetStandLocations()
	responseWriter.Header().Add("Content Type", "application/json")
	data, err := json.Marshal(vm)
	if err == nil {
		responseWriter.Write(data)
	} else {
		responseWriter.WriteHeader(http.StatusNotFound)
	}
}
