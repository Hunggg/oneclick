package http

import (
	"encoding/json"
	"fmt"
	"oneclick/config"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/form/v4"
	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

var (
	conform   = modifiers.New()
	decoder   = form.NewDecoder()
	validate  = validator.New()
	pathRegex = regexp.MustCompile(`^\/?([A-z\d]+)/([A-z\d]+)/([A-z\d]+)\/?$`)
)

type Server struct {
	l          *zap.SugaredLogger
	engine     *gin.Engine
	services   map[string]interface{}
}

func NewServer() (*Server, error) {
	var env config.Env
	env.LoadConfig()
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
	s := &Server{
		l:          zap.S(),
		engine:     api,
		services:   make(map[string]interface{}),
	}
	return s, nil
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


