package main // import "github.com/SCP-028/vuelog"

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// in case Nginx stops working (very unlikely)
	r.Use(static.Serve("/static", static.LocalFile("./dist/static", true)))
	r.LoadHTMLGlob("./dist/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// group: api
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	r.Run(":9587")
}
