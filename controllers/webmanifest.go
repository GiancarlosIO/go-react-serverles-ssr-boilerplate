package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// WebManifestHandler serves the site.webmanifest file
func (s *Server) WebManifestHandler(wr http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.ServeFile(wr, r, "frontend/dist/static/manifest.json")
}

// BrowserConfigFileHandler serves the browserconfig.xml file
func (s *Server) BrowserConfigFileHandler(wr http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.ServeFile(wr, r, "frontend/dist/static/browserconfig.xml")
}

// YanderBrowserFileHandler serves the yandex-browser-manifest.json file
func (s *Server) YanderBrowserFileHandler(wr http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.ServeFile(wr, r, "frontend/dist/static/yandex-browser-manifest.json")
}

// ManifestWebAppFileHandler serves the manifest.webapp file
func (s *Server) ManifestWebAppFileHandler(wr http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.ServeFile(wr, r, "frontend/dist/static/manifest.webapp")
}
