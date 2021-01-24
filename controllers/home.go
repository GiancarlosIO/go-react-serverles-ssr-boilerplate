package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// HomeHandler is the handler for the homepage
func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, pageVariables PageVariables) {
	tpl, err := template.ParseFiles(pageVariables.TemplatePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error", err)
	}

	err = tpl.Execute(w, pageVariables)
	if err != nil {
		log.Print("Template execute error", err)
	}
}
