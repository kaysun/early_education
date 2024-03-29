package user

import (
	"crypto/md5"
	"fmt"
	"github.com/circle/early_education/errors"
	"github.com/circle/early_education/model"
	"github.com/circle/early_education/service/user"
	"github.com/circle/early_education/transport/user/proto"
	"github.com/circle/early_education/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var userService user.UserService

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
		"msg":  "register success",
		"data": proto.RegisterResp{
			UserID: userID,
		},
	})
}

func Login(context *gin.Context) {
	var loginReq proto.LoginRequest
	if err := context.ShouldBindJSON(&loginReq); err != nil {
		return
	}

	userID, err := userService.Login(model.User{
		Name:     loginReq.UserName,
		Password: loginReq.PassWord,
	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": errors.LoginError,
			"msg":  err.Error(),
		})
		return
	}

	if userID > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "login success",
			"data": proto.LoginResp{
				UserID: userID,
			},
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code": errors.LoginUserOrPasswordInvalid,
		"msg":  "login failed",
	})
}

// AdminUserLogin 管理页面登录
func AdminUserLogin(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.AdminUserLoginResp{},
	})
}
