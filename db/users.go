package db

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type (
	userTmpl struct {
		User        string `form:"username" json:"username" binding:"required"`
		Password    string `form:"password" json:"password" binding:"required,min=6"`
		ConfirmPass string `form:"cofirmPass" json:"cofirmPass" binding:"omitempty,eqfield=Password"`
		Email       string `form:"email" json:"email" binding:"omitempty,email"`
	}
	alterTmpl struct {
		User     string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required,min=6"`
		NewPass  string `form:"newPass" json:"newPass" binding:"required,min=6,nefield=Password"`
	}
)

// CreateUser INSERT INTO users table. Does nothing if username already exists.
func CreateUser(c *gin.Context) {
	var json userTmpl
	if errs := c.ShouldBind(&json); errs == nil {
		db := Init()
		defer db.Close()

		hash, _ := hashPassword(json.Password)
		user := Users{
			Name:     json.User,
			Password: hash,
			Email:    json.Email,
		}
		if err := db.Create(&user).Error; err == nil {
			// Create new user
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
				"err":    err,
			})
		}
	} else {
		// Form validation errors
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "",
			"err":    errs,
		})
	}
}

// DeleteUser actually a soft delete as there's the "deleted_at" field.
func DeleteUser(c *gin.Context) {
	var json userTmpl
	db := Init()
	defer db.Close()

	if errs := c.ShouldBind(&json); errs == nil {
		var user Users
		if err := db.Where("name = ?", json.User).First(&user).Error; gorm.IsRecordNotFoundError(err) {
			// User not found
			c.JSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"msg":    "",
				"err":    err,
			})
		} else {
			if checkPasswordHash(json.Password, user.Password) {
				// Soft delete user
				db.Delete(&user)
				c.JSON(http.StatusOK, gin.H{
					"status": http.StatusOK,
					"msg":    fmt.Sprintf("User %s deleted.", json.User),
					"err":    "",
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": http.StatusBadRequest,
					"msg":    "",
					"err":    fmt.Sprintf("Wrong password for user %s", user.Name),
				})
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "",
			"err":    errs,
		})
	}
}

// ChangePassword ...
func ChangePassword(c *gin.Context) {
	var json alterTmpl
	if errs := c.ShouldBind(&json); errs == nil {
		db := Init()
		defer db.Close()
		var user Users

		if err := db.Where("name = ?", json.User).First(&user).Error; err != nil {
			// User not found
			c.JSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"msg":    "",
				"err":    err,
			})
		} else {
			if checkPasswordHash(json.Password, user.Password) {
				newHash, _ := hashPassword(json.NewPass)
				db.Model(&user).Update("password", newHash)
				c.JSON(http.StatusOK, gin.H{
					"status": http.StatusOK,
					"msg":    fmt.Sprintf("Successfully updated password for user %s", user.Name),
					"err":    "",
				})
			} else {
				// Old password doesn't match
				c.JSON(http.StatusBadRequest, gin.H{
					"status": http.StatusBadRequest,
					"msg":    "",
					"err":    fmt.Sprintf("Wrong password for user %s.", user.Name),
				})
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "",
			"err":    errs,
		})
	}
}

// FetchAllUsers lists all users in the database.
func FetchAllUsers(c *gin.Context) {
	db := Init()
	defer db.Close()

	var users []Users
	db.Select("uid, name, email, created_at, updated_at").
		Where("deleted_at is NULL").
		Find(&users)
	userNum := len(users)
	if userNum <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"msg":    "",
			"err":    "There aren't any users in the database.",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
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
