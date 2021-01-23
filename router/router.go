package nrouter

import (
	"fmt"
	"net/http"
)


type route struct {
	method string
	segs []string
	handler http.Handler
	prefix bool
}

type Router struct {
	trees map[string]*node
	// Function to handle panics recovered from http handlers
	// It should be used to generate an error page and return the http error code 500 (Internal Server Error)
	// The handler can be used to keep your server from crashing because of
	// unrecovered panic
	PanicHandler func(w http.ResponseWriter, r *http.Request, err interface{})
	// TODO
	RedirectTrailingSlash bool
	// TODO
	RedirectFixedPath bool
	// TODO
	HandleMethodNotAllowed bool
	// TODO
	HandleOPTIONS bool

}
type Param struct {
	key string
	value string
}
type Params []Param

// Handler is a function that can be registered to a route to handle a http request
type Handler func(w http.ResponseWriter, r *http.Request, p *Params)

func New() *Router {
	return &Router{
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      true,
		HandleMethodNotAllowed: true,
		HandleOPTIONS:          true,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("serve http",  r.trees)
	if r.PanicHandler != nil {
		// recover the error inside the goroutine(deferred function)
		defer r.recv(w, req)
	}
	path := req.URL.Path

	if root := r.trees[path]; root != nil {
		ps := Params{
			Param{key: "a", value: "b"},
		}
		root.handler(w, req, &ps)
	}
}

func (r *Router) GET(path string, handle Handler) {
	r.Handle(http.MethodGet, path, handle)
}

// Handle registers a new request handle with the given path and method
// For GET, POST, PUT, PATCH and DELETE the respective shortcuts can be used.

// This function is intended for bulk loading and to allow the usage of less
// frequently used, non-standardized or custom methods (e.g. for internal
// communication with a proxy)
func (r *Router) Handle(method, path string, handler Handler) {
	// fmt.Printf("Router Register: Method: %s, Path: %s\n", method, path)
	// varscount := uint16(0)
	if method == "" {
		panic("method must not be empty")
	}
	// we need to use the '/' (simple quotation marks) because path[0] returns a byte not a string
	if len(path) < 1 || path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}
	if handler == nil {
		panic("handle must not be nil")
	}

	// if the handler doesn't exists initialize it with an empty map
	if r.trees == nil {
		r.trees = make(map[string]*node)
	}

	// if the handler doesn't exists, save it to the trees map (old implementation)
	/*root := r.trees[path]
	if root == nil {
		root = &node{
			path:   path,
			handler: handle,
		}
		r.trees[path+method] = root
	}*/
	// each method will have a root node
	root := r.trees[method]
	if root == nil {
		root = new(node)
		r.trees[method] = root
	}
	root.addRoute(path, handler)
}

func (r *Router) recv(w http.ResponseWriter, req *http.Request) {
	// The recover built-in function allows a program to manage behavior of a
	// panicking goroutine. Executing a call to recover inside a deferred
	// function (but not any function called by it) stops the panicking sequence
	// by restoring normal execution and retrieves the error value passed to the
	// call of panic. If recover is called outside the deferred function it will
	// not stop a panicking sequence. In this case, or when the goroutine is not
	// panicking, or if the argument supplied to panic was nil, recover returns
	// nil. Thus the return value from recover reports whether the goroutine is
	// panicking.
	if rcv := recover(); rcv != nil {
		r.PanicHandler(w, req, rcv)
	}
}