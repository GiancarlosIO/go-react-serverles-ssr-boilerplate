package main

import (
	"github.com/julienschmidt/httprouter"
	"mrn-portfolio/controllers"
	"mrn-portfolio/database"
	"mrn-portfolio/utils"
	"net/http"
	"os"
)


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db := database.DB{}
	dbconn := db.Open()

	router := httprouter.New()

	s := controllers.Server{
		DB: dbconn,
		Router: router,
	}
	//s.Router.GET("/", s.HomeHandler)
	s.Router.GET("/save", s.HomeHandler)
	s.Router.GET("/", s.HomeHandler)

	err := http.ListenAndServe(":"+port, s.Router)
	utils.HandleError(err)
}
