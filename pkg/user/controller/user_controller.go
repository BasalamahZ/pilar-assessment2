package controller

import (
	"assessment2/pkg/user/dto"
	"assessment2/pkg/user/usecase"
	jwt_local "assessment2/utils/helper/jwt"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserHTTPController struct {
	usecase usecase.UsecaseInterfaceUser
}

func InitControllerUser(uc usecase.UsecaseInterfaceUser) *UserHTTPController {
	return &UserHTTPController{
		usecase: uc,
	}
}

func (uc *UserHTTPController) Register(c *gin.Context) {
	var req dto.UserDTO
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.usecase.Register(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	access_token, err := jwt_local.GenerateNewToken(jwt.MapClaims{
		"username": req.Username,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"status":  http.StatusCreated,
		"data": gin.H{
			"access_token": access_token,
		},
	})
}

func (uc *UserHTTPController) Login(c *gin.Context) {
	var req dto.UserLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.usecase.CreateTokenUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	access_token, err := jwt_local.GenerateNewToken(jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refresh_token, err := jwt_local.GenerateNewToken(jwt.MapClaims{
		"username": user.Username,
		"id":       user.ID,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", access_token, 60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refresh_token, 60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", 60, "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  http.StatusOK,
		"data": gin.H{
			"access_token": access_token,
			"user":         user,
		},
	})
}

func (uc *UserHTTPController) RefreshAccessToken(c *gin.Context) {
	message := "could not refresh access token"

	cookie, err := c.Cookie("refresh_token")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
		return
	}

	sub, err := jwt_local.ParseToken(cookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// userInfo := c.MustGet("user_info").(jwt.MapClaims)
	// id := userInfo["id"].(int)
	newId, err := strconv.ParseInt(fmt.Sprint(sub), 10, 0)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	user, err := uc.usecase.GetUserById(int(newId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
		return
	}

	access_token, err := jwt_local.GenerateNewToken(jwt.MapClaims{
		"username": user.Username,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.SetCookie("access_token", access_token, 60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", 60, "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}

func (uc *UserHTTPController) ViewProfile(c *gin.Context) {
	userInfo := c.MustGet("user_info").(jwt.MapClaims)
	username := userInfo["username"].(string)
	user, err := uc.usecase.ViewProfile(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success get post",
		"data": gin.H{
			"user": user,
		},
	})
}
