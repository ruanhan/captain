package main

import (
	"github.com/captain/deployment"
	"github.com/captain/lib"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()
	r.GET("/deployments", func(context *gin.Context) {
		list := lib.DataBuilder().SetTitle("deployments").SetData("list", deployment.ListAll("default"))
		context.JSON(200, gin.H{
			"data":    list,
			"message": "pong",
		})
	})
	r.Run(":8081")
}
