package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const SERVER_PORT string = ":8080"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	albumRouter := r.Group("/album")
	{
		albumRouter.GET("/:albumId/comment", func(c *gin.Context) {
			albumId := c.Param("albumId")
			c.JSON(http.StatusOK, gin.H{
				"paramName": albumId,
			})
		})
	}

	return r
}

func main() {
	r := SetupRouter()
	r.Run(SERVER_PORT)
}
