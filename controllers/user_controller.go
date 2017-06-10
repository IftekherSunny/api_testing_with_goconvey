package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iftekhersunny/api_testing_with_goconvey/models"
)

// User controller struct
type UserController struct {
	User models.User
}

// Create a new user
func (self *UserController) Create(c *gin.Context) {
	name := c.PostForm("name")
	age, _ := strconv.Atoi(c.PostForm("age"))

	c.JSON(200, gin.H{
		"message": self.User.Create(name, age),
	})
}

// Get all users
func (self *UserController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": self.User.All(),
	})
}
