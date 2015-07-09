package controllers

import (
	"html/template"
	"net/http"

	"github.com/partyschaum/go-web-app/controllers/util"
	"github.com/partyschaum/go-web-app/viewmodels"
)

type profileController struct {
	template *template.Template
}

func (p *profileController) handle(w http.ResponseWriter, req *http.Request) {
	gzrw := util.GetResponseWriter(w, req)
	defer gzrw.Close()

	vm := viewmodels.GetProfile()

	if req.Method == "POST" {
		vm.User.Email = req.FormValue("email")
		vm.User.FirstName = req.FormValue("firstName")
		vm.User.LastName = req.FormValue("lastName")
		vm.User.Stand.Address = req.FormValue("standAddress")
		vm.User.Stand.City = req.FormValue("standCity")
		vm.User.Stand.State = req.FormValue("standState")
		vm.User.Stand.Zip = req.FormValue("standZip")
	}

	gzrw.Header().Add("Content Type", "text/html")
	p.template.Execute(gzrw, vm)
}
