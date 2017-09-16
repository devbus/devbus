package controllers

import (
	"github.com/devbus/devbus/models"
	"github.com/astaxie/beego/logs"
	"github.com/devbus/devbus/services"
)

type RegisterController struct {
	jsonApiController
}

func (c *RegisterController) Register() {
	var service services.UserService
	user := &models.User{}
	err := c.parseJson(user)
	if err != nil {
		logs.Debug("failed to parse request, error: %v", err)
		goto FAIL
	}
	service = services.GetUserService()
	if err := service.Register(user); err != nil {
		goto FAIL
	}
	c.renderData(nil)
	return
FAIL:
	c.renderError(err)
}

func (c *RegisterController) Activate() {

}