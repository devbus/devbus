package apis

import (
	"net/http"

	"github.com/devbus/devbus/common"
	"github.com/gin-gonic/gin"
	"github.com/kuun/slog"
)

var jsonApiLog = slog.GetLogger()

func renderError(context *gin.Context, err error) {
	var opErr common.OpError
	var ok bool

	if opErr, ok = err.(common.OpError); !ok {
		opErr = common.OpError{common.ErrUnknown}
		jsonApiLog.Error("handle unknown error: %+v", err)
	} else {
		jsonApiLog.Debug("handle error: %s", err)
	}
	result := &common.RestResult{}
	result.SetError(opErr)
	context.JSON(http.StatusInternalServerError, result)
}

func renderErrorCode(context *gin.Context, code common.ErrorCode) {
	result := &common.RestResult{}
	result.SetErrorCode(code)
	context.JSON(http.StatusInternalServerError, result)
}

func renderData(context *gin.Context, data interface{}) {
	result := common.RestResult{}
	result.SetData(data)
	context.JSON(http.StatusOK, result)
}
