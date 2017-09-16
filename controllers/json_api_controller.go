package controllers

import (
	"github.com/astaxie/beego"
	"github.com/devbus/devbus/common"
	"github.com/astaxie/beego/logs"
	"encoding/json"
)

type jsonApiController struct {
	beego.Controller
}

func (c *jsonApiController) parseJson(data interface{}) error {
	return json.Unmarshal(c.Ctx.Input.RequestBody, data)
}

func (c *jsonApiController) renderError(err error) {
	var opErr common.OpError
	var ok bool

	if opErr, ok = err.(common.OpError); !ok {
		opErr = common.OpError{common.ErrUnknown}
		logs.Error("handle unknown error: %+v", err)
	} else {
		logs.Debug("handle error: %s", err)
	}
	result := &common.RestResult{}
	result.SetError(opErr)
	c.Data["json"] = result
	c.Ctx.Output.SetStatus(500)
	c.ServeJSON(true)
}

func (c *jsonApiController) renderErrorCode(code common.ErrorCode) {
	result := &common.RestResult{}
	result.SetErrorCode(code)
	c.Data["json"] = result
	c.Ctx.Output.SetStatus(500)
	c.ServeJSON(true)
}

func (c *jsonApiController) renderData(data interface{}) {
	result := common.RestResult{}
	result.SetData(data)
	c.Data["json"] = result
	c.ServeJSON(true)
}
