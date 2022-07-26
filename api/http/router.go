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
	client := server.engine.Group("/v1")
	{
		// client.Use(adapter.Wrap(middleware.JWTValidator().CheckJWT))
		categories := client.Group("/categories")
		{
			categories.GET("/:id", controller.GetCategoryById)
			categories.GET("/list", controller.GetListCategory)
			categories.PATCH("/update", controller.UpdateCategory)
			categories.POST("/create", controller.SaveCategory)
			categories.DELETE("/delete/:id", controller.DeleteCategoryById)
		}

		client.POST("/register", controller.Register)
		// client.POST("/logn", controller.Login)
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
