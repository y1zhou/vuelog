package main // import "github.com/SCP-028/vuelog"

import (
	"net/http"

	"github.com/SCP-028/vuelog/db"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.UpdateSchema() // Pull up Postgres and auto migrate schema
	store := db.InitRedis()
	r.Use(sessions.Sessions("y1zhou", store))

	api := r.Group("/api")
	{
		// Test if the server is working
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": {
					"gin-server": "pong",
					"redis": "",
					"postgres": ""
			})
		})

		/*** Users API ***/
		api.GET("/users", db.FetchAllUsers)
		api.POST("/user/signup", db.CreateUser)
		api.POST("/user/delete", db.DeleteUser)
		api.PUT("/user/update", db.UpdateUser)
		api.POST("/user/login", db.LoginUser)
		api.GET("/user/auth", db.AuthUser)
	}
	r.Run(":9587")
}
