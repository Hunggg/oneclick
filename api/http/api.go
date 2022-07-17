package http

import (
	"net/http"
	"time"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New(enableCORS bool) (*gin.Engine, error) {
	engine := gin.Default()
	deployTime := time.Now()

	engine.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	engine.Use(func(ctx *gin.Context) {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.Scope().SetTag("rawUrl", ctx.Request.Host+ctx.Request.RequestURI)
		}
		ctx.Next()
	})

	if enableCORS {
		engine.Use(cors.New(cors.Config{
			AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			AllowAllOrigins: true,
		}))
	}

	engine.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"timestamp":  time.Now().Unix(),
			"now":        time.Now(),
			"deployTime": deployTime,
			"version":    "0.0.1",
		})
	})

	return engine, nil
}