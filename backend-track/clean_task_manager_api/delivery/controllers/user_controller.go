package controllers

import (
	"clean-task-manager-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUseCaseInterface
}

func NewUserController(uu domain.UserUseCaseInterface) *UserController {
	return &UserController{UserUsecase: uu}
}

func (uc *UserController) Register(c *gin.Context) {

	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	err := uc.UserUsecase.Register(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User registered successfully"})
}

func (uc *UserController) Login(c *gin.Context) {
	var LoginRequest domain.LoginRequest

	if err := c.ShouldBindJSON(&LoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Payload"})
		return
	}

	loginResponse, err := uc.UserUsecase.Login(c, LoginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "User logged in successfully", "Response": loginResponse})
}

func (uc *UserController) Promote(c *gin.Context) {
	id := c.Param("id")

	err := uc.UserUsecase.Promote(c, id)

	if err != nil {

		if err.Error() == "invalid user ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Promoted successfully"})

}
