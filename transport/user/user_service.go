package user

import (
	"crypto/md5"
	"fmt"
	"github.com/circle/early_education/errors"
	"github.com/circle/early_education/infra"
	"github.com/circle/early_education/model"
	"github.com/circle/early_education/service/user"
	"github.com/circle/early_education/transport/user/proto"
	"github.com/circle/early_education/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var userService = user.NewUserService(infra.NewUserDAO())

func Resister(context *gin.Context) {
	var registerReq proto.RegisterReq
	if err := context.ShouldBindJSON(&registerReq); err != nil {
		return
	}

	secret := md5.Sum([]byte(registerReq.Password))
	now := time.Now()
	userID, err := userService.Register(model.User{
		Name:       registerReq.Name,
		Password:   fmt.Sprint(secret),
		Nick:       registerReq.Nick,
		Birthday:   util.StrToTime(registerReq.Birthday),
		CreateTime: now,
		UpdateTime: now,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": errors.RegisterError,
			"msg":  err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "login success",
		"data": userID,
	})
}

func Login(context *gin.Context) {
	var loginReq proto.LoginRequest
	if err := context.ShouldBindJSON(&loginReq); err != nil {
		return
	}

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
