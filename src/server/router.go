package server

import (
	"api/src/server/routes"
	"fmt"
	"net/http"
)

type Server interface {
	Config()
	Start()
}

type standardServer struct {
	mux    *http.ServeMux
	routes []routes.Route
	server *http.Server
}

func (s *standardServer) Config() {

	s.routes = append(s.routes, routes.Users...)
	for _, route := range s.routes {
		s.mux.Handle(route.GetWildCard(), route.Handler)
	}
}

func (s *standardServer) Start() {
	s.server.ListenAndServe()
}

func NewServer(url, port string) Server {
	s := new(standardServer)
	s.mux = http.NewServeMux()
	s.server = &http.Server{
		Addr: fmt.Sprintf("%s:%s", url, port),
	}
	s.Config()
	return s

}
