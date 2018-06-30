package main // import "github.com/SCP-028/vuelog"

import (
	"net/http"

	"database/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.Init()
	// group: api
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	r.Run(":9587")
}
