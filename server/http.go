package server

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"ServiceStructure/config"
	"ServiceStructure/server/handler"
	"ServiceStructure/services/ping"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	chiTrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-chi/chi"
)

// Server represents a chi server mux
type Server struct {
	Router  *chi.Mux
	Address string
}

// New setups and returns a server
func New() *Server {
	r := chi.NewMux()
	addr := config.Config.Service.Host + ":" + strconv.Itoa(config.Config.Service.Port)
	server := Server{Router: r, Address: addr}
	server.SetupMiddleware()
	server.SetupRouter()
	return &server
}

// SetupMiddleware setups middleware
func (s Server) SetupMiddleware() {
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(chiTrace.Middleware(chiTrace.WithServiceName(config.Config.Service.ServiceName)))
	s.Router.Use(middleware.Timeout(time.Duration(config.Config.Service.ServerTimeout) * time.Minute))
}

// SetupRouter setups router
func (s Server) SetupRouter() {
	// Health check route
	s.Router.Route("/v1", func(r chi.Router) {
		r.Get("/ping", handler.Make(ping.Ping))
	})
}

// ServeHTTP starts http server
func (s Server) ServeHTTP() {
	log.Fatal(http.ListenAndServe(s.Address, s.Router))
}
