package http

import (
	"oneclick/api/http/controller"
	"oneclick/config"
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
		client.GET("/categories", controller.GetListCategory)
		client.POST("/categories", controller.SaveCategory)
		client.POST("/batch/categories", controller.SaveBatchCategory)
	}

	server.engine.Run(":" + env.HttpPort)
}
