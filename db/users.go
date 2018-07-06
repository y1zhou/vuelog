package db

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser INSERT INTO users table. Does nothing if username already exists.
func CreateUser(c *gin.Context) {
	password := c.PostForm("password")
	if len(password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "",
			"err":    "Password should be at least 6 characters.",
		})
	} else {
		db := Init()
		defer db.Close()

		hash, _ := hashPassword(password)
		user := Users{
			Name:     c.PostForm("username"),
			Password: hash,
		}
		if db.NewRecord(user) {
			// Create new user
			db.Create(&user)
			c.JSON(http.StatusCreated, gin.H{
				"status": http.StatusCreated,
				"msg":    fmt.Sprintf("Successfully created user %s", user.Name),
				"err":    "",
			})
		} else {
			// User already exists
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"msg":    "",
				"err":    fmt.Sprintf("Username %s already exists.", user.Name),
			})
		}
	}
}

// DeleteUser actually a soft delete as there's the "deleted_at" field.
func DeleteUser(c *gin.Context) {
	db := Init()
	defer db.Close()

	username := c.PostForm("username")
	hash, _ := hashPassword(c.PostForm("password"))

	var user Users
	if err := db.Where("name = ?", username).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		// User not found
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"msg":    "",
			"err":    err,
		})
	} else {
		if user.Password == hash {
			// Soft delete user
			db.Delete(&user)
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"msg":    fmt.Sprintf("User %s deleted.", username),
				"err":    "",
			})
		}
	}
}

// ChangePassword ...
func ChangePassword(c *gin.Context) {
	db := Init()
	defer db.Close()

	username := c.PostForm("username")
	oldHash, _ := hashPassword(c.PostForm("old_password"))
	var user Users

	if err := db.Where("name = ?", username).First(&user).Error; err != nil {
		// User not found
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"msg":    "",
			"err":    err,
		})
	} else {
		if user.Password == oldHash {
			// Old password correct
			newHash, _ := hashPassword(c.PostForm("new_password"))
			user.Password = newHash
			db.Save(&user)
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"msg":    fmt.Sprintf("Successfully updated password for user %s", user.Name),
				"err":    "",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"msg":    "",
				"err":    fmt.Sprintf("Wrong password for user %s.", username),
			})
		}
	}
}

// FetchAllUsers lists all users in the database.
func FetchAllUsers(c *gin.Context) {
	db := Init()
	defer db.Close()

	var users []Users
	db.Select("uid, name, email, group, created_at, updated_at").Find(&users)
	userNum := len(users)
	if userNum <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"msg":    "",
			"err":    "There aren't any users in the database.",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusNotFound,
			"msg":    fmt.Sprintf("Returning %d users.", userNum),
			"err":    "",
			"data":   users,
		})
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
