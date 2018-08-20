package db

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser INSERT INTO users table. Does nothing if username already exists.
func CreateUser(c *gin.Context) {
	var json userSignupForm
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
		hash, _ := hashPassword(json.Password)
		user := Users{
			Name:     json.Username,
			Password: hash,
			Email:    json.Email,
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

// DeleteUser actually a soft delete as there's the "deleted_at" field.
func DeleteUser(c *gin.Context) {
	var json userDeleteForm
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "",
			"err":    err,
		})
	} else {
		db := Init()
		defer db.Close()
		var user Users
		if err := db.Where("name = ?", json.Username).
			First(&user).Error; gorm.IsRecordNotFoundError(err) {
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
					"msg":    fmt.Sprintf("User %s deleted.", json.Username),
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
	}
}

// UpdateUser ...
func UpdateUser(c *gin.Context) {
	var json userUpdateForm
	if errs := c.ShouldBind(&json); errs == nil {
		db := Init()
		defer db.Close()
		var user Users

		if err := db.Where("name = ?", json.Username).First(&user).Error; err != nil {
			// User not found
			c.JSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"msg":    "",
				"err":    err,
			})
		} else {
			if checkPasswordHash(json.Password, user.Password) {
				newHash, _ := hashPassword(json.NewPass)
				db.Model(&user).
					Updates(map[string]interface{}{"password": newHash, "email": json.Email})
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

// QueryUser ...
func QueryUser(c *gin.Context) {
	var json userQueryForm
}

// LoginUser Start login session
func LoginUser(c *gin.Context) {
	var json userLoginForm
	session := sessions.Default(c)
	if errs := c.ShouldBind(&json); errs == nil {
		db := Init()
		defer db.Close()
		var user Users

		if err := db.Where("name = ?", json.Username).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"msg":    "",
				"err":    err,
			})
		} else {
			if checkPasswordHash(json.Password, user.Password) {
				session.Set("user", user.UID)
				session.Save()
				c.JSON(http.StatusOK, gin.H{
					"status": http.StatusOK,
					"msg":    fmt.Sprintf("User %s successfully logged in!", user.Name),
					"err":    "",
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": http.StatusBadRequest,
					"msg":    "",
					"err":    fmt.Sprintf("Wrong password for user %s.", user.Name),
				})
			}
		}
	}
}

// AuthUser Check if user session is logged in
func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user")
		if userID == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"msg":    "",
				"err":    "Invalid session token.",
			})
		} else {
			c.Next()
		}
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
