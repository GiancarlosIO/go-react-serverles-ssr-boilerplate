package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

type BaseController struct {
}

type ssr struct {
	HTML     template.HTML
	CSS      template.HTML
	MetaTags template.HTML
}

type webpack struct {
	Entry      string
	StaticPath string
}

// PageVariables defines the data that the current page needs in order to render the ui
type PageVariables struct {
	SSR      ssr
	Webpack  webpack
	Data     interface{}
	Template *template.Template
}

// Handler defined the function that a controller should expose
type Handler func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, pageVariables PageVariables)

var client http.Client = http.Client{
	Timeout: 500 * time.Millisecond,
}

// CreateHandler creates a handler that calls the handler with the pageVariables struct
// it will use the webpackEntry value to fetch the ssr data and also to map the template and static paths needed for the current page
func (s *Server) CreateHandler(webpackEntry string, handler Handler) httprouter.Handle {
	ssrEndpoint := os.Getenv("SSR_ENDPOINT")
	if ssrEndpoint == "" {
		ssrEndpoint = "http://localhost:3000/dev"
	}
	ssrEndpoint = ssrEndpoint + "/" + webpackEntry

	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		start := time.Now()
		tpl, err := template.ParseFiles("frontend/dist/static/" + webpackEntry + ".gohtml")
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			log.Print("Template parsing error", err)
		}

		body := map[string]string{
			"url": r.URL.Path,
		}
		jsonBody, err := json.Marshal(body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			log.Printf("Failed to marshal the body for the WebpackEntry: %s", webpackEntry)
		}

		fmt.Println("> Making a Request to SSR Endpoint: ", ssrEndpoint)
		res, err := client.Post(ssrEndpoint, "application/json", bytes.NewBuffer(jsonBody))
		var html, css, metatags template.HTML
		if err == nil {
			var ssrData map[string]string
			json.NewDecoder(res.Body).Decode(&ssrData)
			// fmt.Println(ssrData)
			metatags = template.HTML(ssrData["metaTags"])
			html = template.HTML(ssrData["html"])
			css = template.HTML(ssrData["css"])
		} else {
			// http.Error(rw, err.Error(), http.StatusBadRequest)
			log.Printf("Failed to get the response from the serverless ssr. WebpackEntry: %s", webpackEntry)
		}

		pageVariables := PageVariables{
			SSR: ssr{
				MetaTags: metatags,
				HTML:     html,
				CSS:      css,
			},
			Webpack: webpack{
				Entry: webpackEntry,
			},
			Template: tpl,
		}

		handler(rw, r, p, pageVariables)

		// logging code
		duration := time.Since(start).Milliseconds()
		t, _ := time.Parse(time.RFC822Z, start.String())
		fmt.Printf("> %s %s: [%s] %s %vms\n", webpackEntry, t, r.Method, r.URL.Path, duration)
	}
}
