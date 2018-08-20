package db

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CreateContent INSERT INTO contents table.
func CreateContent(c *gin.Context) {
	var json contentCreateForm
	if err := c.ShouldBind(&json); err != nil {
		// Form validation errors
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "",
			"err":    err,
		})
	} else {
		db := Init()
		defer db.Close()
		var user []Users
		if err := db.Where("name = ?", json.Author).
		First(&user).Error; gorm.IsRecordNotFoundError(err) {

		}
		content := Contents{
			Title:    json.Title,
			Text:     json.Text,
			AuthorID: user.UID,
			Status:   json.Status,
			Slug:     ""
		}
		if err := db.Create(&user).Error; err != nil {
			// User already exists
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"msg":    "",
				"err":    err,
			})
		} else {
			// Create new user
			c.JSON(http.StatusCreated, gin.H{
				"status": http.StatusCreated,
				"msg":    fmt.Sprintf("Successfully created user %s", user.Name),
				"err":    "",
			})
		}
	}
}
