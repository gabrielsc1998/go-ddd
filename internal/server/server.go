package webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]map[string]http.HandlerFunc // [path][method
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, method string, handler http.HandlerFunc) {
	if _, ok := s.Handlers[path]; !ok {
		s.Handlers[path] = make(map[string]http.HandlerFunc)
	}
	s.Handlers[path][method] = handler
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	methods := map[string]interface{}{
		"GET":    s.Router.Get,
		"POST":   s.Router.Post,
		"PUT":    s.Router.Put,
		"DELETE": s.Router.Delete,
	}
	for path, handler := range s.Handlers {
		for method, handler := range handler {
			methods[method].(func(string, http.HandlerFunc))(path, handler)
		}
	}
	fmt.Println("Starting web server on port", s.WebServerPort)
	http.ListenAndServe(":"+s.WebServerPort, s.Router)
}
