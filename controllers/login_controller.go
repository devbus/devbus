package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/devbus/devbus/common"
	"github.com/devbus/devbus/models"
	"github.com/devbus/devbus/services"
)


type LoginController struct {
	jsonApiController
}

// Post handles user's login request
func (c *LoginController) Post() {
	var service services.UserService
	user := &models.User{}
	err := c.parseJson(user)
	if err != nil {
		logs.Debug("failed to parse request data, error: %+v", err)
		goto FAIL
	}
	if err != nil || user.Email == "" || user.Password == "" {
		goto FAIL
	}
	if service.Auth(user.Email, user.Password) {
		c.StartSession()
		c.renderData(nil)
		return
	}
FAIL:
	c.renderErrorCode(common.ErrPasswrod)
}
