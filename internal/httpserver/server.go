package httpserver

import (
	"fmt"
	"net/http"

	"github.com/anton-okolelov/json-app/internal/service"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	userService service.Service
	router      chi.Router
	port        int
}

func New(userService service.Service, port int) Server {
	s := Server{
		userService: userService,
		port:        port,
		router:      chi.NewRouter(),
	}

	s.router.Post("/users", s.createUser)
	s.router.Get("/users/{id:[0-9]+}", s.getUser)
	return s
}

func (s Server) Start() error {
	fmt.Printf("Start server at %d\n", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router)
}
