package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/partyschaum/go-web-app/controllers/util"
	"github.com/partyschaum/go-web-app/models"
	"github.com/partyschaum/go-web-app/viewmodels"
)

type homeController struct {
	template      *template.Template
	loginTemplate *template.Template
}

func (h *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()

	w.Header().Add("Content Type", "text/html")

	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	h.template.Execute(responseWriter, vm)
}

func (h *homeController) login(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	w.Header().Add("Content Type", "text/html")

	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")

		member, err := models.GetMember(email, password)

		if err == nil {
			session, err := models.CreateSession(member)
			log.Printf("create session err: %v", err)
			if err == nil {
				var cookie http.Cookie
				cookie.Name = "sessionId"
				cookie.Value = session.SessionId()
				responseWriter.Header().Add("Set-Cookie", cookie.String())
			}
		} else {
			log.Print("User not found!")
		}
	}

	vm := viewmodels.GetLogin()

	h.loginTemplate.Execute(responseWriter, vm)
}
