package main

import (
	"mrn-portfolio/controllers"
	"mrn-portfolio/database"
	"mrn-portfolio/utils"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
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
		DB:     dbconn,
		Router: router,
	}

	s.Router.ServeFiles("/static/*filepath", http.Dir("static"))
	s.Router.GET("/", s.CreateHandler("homepage", s.HomeHandler))

	err := http.ListenAndServe(":"+port, s.Router)
	utils.HandleError(err)
}
