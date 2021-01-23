package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)



func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("Homepage!!")
}