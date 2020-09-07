package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golly/api/handler"
	"log"
)

func route(c *gin.Engine) {
	api := c.Group("/api")
	{
		api.GET("/tasks", handler.ListTask)
		api.GET("/task/:id", handler.GetTask)
		api.POST("/task", handler.CreateTask)
		api.PUT("/task", handler.UpdateTask)
		api.DELETE("/task/:id", handler.DeleteTask)
		api.POST("/task/:id/run", handler.RunTask)
		api.POST("/task/:id/abort", handler.AbortTask)
		api.GET("/task/:id/log", handler.LogTask)
	}
}

func InitAPI() {
	g := gin.Default()
	g.Use(cors.Default())
	route(g)
	log.Fatal(g.Run(":8090"))
}
