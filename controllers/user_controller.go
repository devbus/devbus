package controllers

import (
	"github.com/devbus/devbus/common"
	"github.com/devbus/devbus/models"
	"github.com/devbus/devbus/services"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kuun/slog"
)

var userLog = slog.GetLogger()

func login(context *gin.Context) {
	var service services.UserService
	user := &models.User{}
	err := context.BindJSON(user)
	if err != nil {
		userLog.Debug("failed to parse request data, error: %+v", err)
		goto FAIL
	}
	if err != nil || user.Email == "" || user.Password == "" {
		goto FAIL
	}
	if service.Auth(user.Email, user.Password) {
		session := sessions.Default(context)
		session.Save()
		renderData(context, nil)
		return
	}
FAIL:
	renderErrorCode(context, common.ErrPasswrod)
}

func register(context *gin.Context) {
	var service services.UserService
	user := &models.User{}
	err := context.BindJSON(user)
	if err != nil {
		userLog.Debug("failed to parse request, error: %v", err)
		goto FAIL
	}
	service = services.GetUserService()
	if err := service.Register(user); err != nil {
		goto FAIL
	}
	renderData(context, nil)
	return
FAIL:
	renderError(context, err)
}

func activate(context *gin.Context) {

}

func modifyPassword(context *gin.Context) {

}

func init() {
	router := gin.Default().Group("/api/user")
	router.POST("/login", login)
	router.POST("/register", register)
	router.GET("/activate", activate)
	router.PUT("/password", modifyPassword)
}
