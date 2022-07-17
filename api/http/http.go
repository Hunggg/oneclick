package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"oneclick/config"
	"os"
	"os/signal"
	"regexp"
	"strings"

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

type errorResponse struct {
	Timestamp int64  `json:"timestamp"`
	Error     string `json:"error"`
}

func parseJSON(input interface{}, c *gin.Context) error {
	defer c.Request.Body.Close()
	contentType := c.ContentType()
	switch contentType {
	case "application/json":
		decoder := json.NewDecoder(c.Request.Body)
		return decoder.Decode(input)
	default:
		return fmt.Errorf("Unspported contentType type: %s", contentType)
	}
}

func normalizeMethod(method string) string {
	if len(method) == 0 {
		return method
	}
	return strings.ToUpper(method[0:1]) + method[1:]
}

func getParamMap(params gin.Params) map[string][]string {
	res := make(map[string][]string)
	for _, p := range params {
		res[p.Key] = []string{p.Value}
	}
	return res
}


