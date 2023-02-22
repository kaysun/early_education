package user

import (
	"github.com/circle/early_education/model"
	"github.com/circle/early_education/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Resister(context *gin.Context) {

}

type LoginRequest struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

func Login(context *gin.Context) {
	var loginReq LoginRequest
	if err := context.ShouldBindJSON(&loginReq); err != nil {
		return
	}

	userService := user.UserService{}
	flag, err := userService.Login(model.User{
		Name:     loginReq.UserName,
		Password: loginReq.PassWord,
	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": 10010,
			"msg":  err.Error(),
		})
		return
	}

	if flag {
		context.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "login success",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code": 20010,
		"msg":  "login failed",
	})
}
