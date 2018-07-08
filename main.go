package main // import "github.com/SCP-028/vuelog"

import (
	"net/http"

	"github.com/SCP-028/vuelog/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.UpdateSchema()

	api := r.Group("/api")
	{
		// Test if the server is working
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		/*** Users API ***/
		api.POST("/user/signup", db.CreateUser)
		api.POST("/user/delete", db.DeleteUser)
		api.PUT("/user/modify", db.ChangePassword)
		api.GET("/users", db.FetchAllUsers)
	}
	r.Run(":9587")
}
