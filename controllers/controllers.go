package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

type Handler func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, pageVariables PageVariables)

// CreateHandler creates a handler that calls the handler with the pageVariables struct
// it will use the webpackEntry value to fetch the ssr data and also to map the template and static paths needed for the current page
func (s *Server) CreateHandler(webpackEntry string, handler Handler) httprouter.Handle {

	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		start := time.Now()
		tpl, err := template.ParseFiles("frontend/dist/static/app.gohtml")
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			log.Print("Template parsing error", err)
		}

		pageVariables := PageVariables{
			SSR: ssr{
				MetaTags: template.HTML(`<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Mr N</title>`),
				HTML: template.HTML("<h1>hell asdasd asdasdo</h1>"),
				CSS:  template.HTML("<style>h1{color: red}</style>"),
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
