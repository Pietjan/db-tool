package server

import (
	"db-tool/pkg/app"
	"io/fs"
	"log"
	"net/http"

	"github.com/pietjan/template"

	"github.com/julienschmidt/httprouter"
)

type Option = func(*server)

type Server interface {
	Serve() error
}

func New(options ...Option) Server {
	server := &server{
		router: httprouter.New(),
	}

	for _, fn := range options {
		fn(server)
	}

	routes(server)

	return server
}

func Static(static fs.FS) Option {
	return func(s *server) {
		s.static = static
	}
}

func Template(template template.Template) Option {
	return func(s *server) {
		s.template = template
	}
}

func Application(app app.Application) Option {
	return func(s *server) {
		s.app = app
	}
}

type server struct {
	app      app.Application
	static   fs.FS
	router   *httprouter.Router
	template template.Template
}

// Serve implements Server.
func (s server) Serve() error {
	return http.ListenAndServe(`:8080`, s.router)
}

func routes(s *server) {
	s.router.ServeFiles(`/static/*filepath`, http.FS(s.static))

	s.router.GET(`/`, index(s))
}

func index(s *server) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if err := s.template.Render(w, `page/index`, nil); err != nil {
			log.Println(err)
		}

		return
	}
}
