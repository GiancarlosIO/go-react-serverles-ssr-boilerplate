package controllers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// BlogHandler is the handler for the homepage
func (s *Server) BlogHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, pageVariables PageVariables) {
	err := pageVariables.Template.Execute(w, pageVariables)
	if err != nil {
		log.Print("Template execute error", err)
	}
}
