package http

import (
	"net/http"
	"oneclick/config"
	"os"
	"os/signal"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/form/v4"
	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

var env config.Env

var (
	conform   = modifiers.New()
	decoder   = form.NewDecoder()
	validate  = validator.New()
	pathRegex = regexp.MustCompile(`^\/?([A-z\d]+)/([A-z\d]+)/([A-z\d]+)\/?$`)
)

type Server struct {
	l          *zap.SugaredLogger
	engine     *gin.Engine
	httpServer *http.Server
	services   map[string]interface{}
}


func NewServer() (*Server, error) {
	var corsEnable bool
	if env.EnableCors == "1" {
		corsEnable = true
	} else {
		corsEnable = false
	}
	api, err := New(corsEnable)
	if err != nil {
		return nil, err
	}

	srv := http.Server{
		Addr:    ":" + env.HttpPort,
		Handler: api,
	}

	s := &Server{
		l:          zap.S(),
		engine:     api,
		httpServer: &srv,
		services:   make(map[string]interface{}),
	}

	return s, nil
}

func (s *Server) Start() {
	s.l.Info("Starting API server...")
	go func() {
		if err := s.httpServer.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	s.l.Infof("Listening on %s\n", s.httpServer.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	s.l.Infof("Shutting down server... Reason: %v", sig)
	// teardown logic...

	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	s.l.Infof("Server gracefully stopped")
}

