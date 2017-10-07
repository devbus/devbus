package apis

import (
	"github.com/devbus/devbus/common"
	"github.com/devbus/devbus/models"
	"github.com/devbus/devbus/services"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var service services.UserService
	user := &models.User{}
	err := context.BindJSON(user)
	if err != nil {
		log.Debug("failed to parse request data, error: %+v", err)
		goto FAIL
	}
	service = services.GetUserService()
	if service.Auth(user.Email, user.Password) {
		session := sessions.Default(context)
		session.Save()
		renderData(context, nil)
		log.Debugf("user login successfully, user: %s", user.Email)
		return
	}
FAIL:
	renderErrorCode(context, common.ErrPasswrod)
}

func Register(context *gin.Context) {
	var service services.UserService
	user := &models.User{}
	err := context.BindJSON(user)
	if err != nil {
		log.Debug("failed to parse request, error: %v", err)
		goto FAIL
	}
	service = services.GetUserService()
	if err = service.Register(user); err != nil {
		goto FAIL
	}
	renderData(context, nil)
	return
FAIL:
	renderError(context, err)
}

func Activate(context *gin.Context) {

}

