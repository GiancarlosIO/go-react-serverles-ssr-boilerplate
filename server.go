package main

import (
	"fmt"
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

	s.Router.ServeFiles("/static/*filepath", http.Dir("frontend/dist/static"))
	s.Router.GET("/manifest.json", s.WebManifestHandler)
	s.Router.GET("/browserconfig.xml", s.BrowserConfigFileHandler)
	s.Router.GET("/yandex-browser-manifest.json", s.YanderBrowserFileHandler)
	s.Router.GET("/manifest.webapp", s.ManifestWebAppFileHandler)

	s.Router.GET("/", s.CreateHandler("homepage", s.HomeHandler))

	err := http.ListenAndServe(":"+port, s.Router)
	fmt.Printf("> Server is running in http://localhost:%s/\n", port)
	utils.HandleError(err)
}
