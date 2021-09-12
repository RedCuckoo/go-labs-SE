package service

import (
	"github.com/redcuckoo/go-labs-SE/lab4/config"
	"net"
	"net/http"
)

type service struct {
	cfg config.Config
}

func (s *service) run() error {
	r := NewRouter(s.cfg)

	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		return err
	}

	return http.Serve(listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		cfg: cfg,
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}