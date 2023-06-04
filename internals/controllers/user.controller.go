package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morelmiles/go-redis-caching/internals/models"
	"github.com/morelmiles/go-redis-caching/pkg/database"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	err := database.DB.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userResponses []models.UserResponse
	for _, user := range users {
		userResponse := &models.UserResponse{
			ID:          user.ID,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		}
		userResponses = append(userResponses, *userResponse)
	}

	c.JSON(http.StatusOK, gin.H{"data": userResponses})
}

func CreateUser(c *gin.Context) {
	var input models.SignUpInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
	}

	user.Prepare()
	err := user.Validate("")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userResponse := &models.UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{"user": userResponse}})
}

func GetUserById(c *gin.Context) {
	var user models.User

	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found!"})
		return
	}

	userResponse := &models.UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	c.JSON(http.StatusOK, gin.H{"data": userResponse})
}

func UpdateUserById(c *gin.Context) {
	var user models.User

	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found!"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Prepare()
	err := user.Validate("")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&user).Updates(input)
	updatedUser := user

	c.JSON(http.StatusOK, gin.H{"data": updatedUser})
}

func DeleteUserById(c *gin.Context) {
	var user models.User

	if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found!"})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "user deleted"})
}
