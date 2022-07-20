package http

import (
	"oneclick/api/http/controller"
	"oneclick/config"
	"oneclick/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	var env config.Env
	env.LoadConfig()
	server, err := NewServer()
	if err != nil {
		server.l.Info(err)
	}
	client := server.engine.Group("/api")
	{
		client.GET("/categories/:id", controller.GetCategoryById)
		client.GET("/list/categories", controller.GetListCategory)
		client.PATCH("/update/categories", controller.UpdateCategory)
		client.POST("/create/categories", controller.SaveCategory)
		client.DELETE("/delete/categories/:id", controller.DeleteCategoryById)
	}
	runSwagger(server.engine)
	server.engine.Run(":" + env.HttpPort)
}

func runSwagger(g *gin.Engine) {
	docs.SwaggerInfo.Title = "ONECLICK API"
	docs.SwaggerInfo.Description = "This is ONECLICK server."
	docs.SwaggerInfo.Version = "1.0"
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
