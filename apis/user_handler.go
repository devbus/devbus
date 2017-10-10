package apis

import (
	"github.com/devbus/devbus/services"
	"github.com/gin-gonic/gin"
)

type PassModInfo struct {
	Email string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func ModifyPassword(context *gin.Context) {
	var service services.UserService
	mod := &PassModInfo{}
	err := context.BindJSON(mod)
	if err != nil {
		log.Debug("failed to parse request, error: %v", err)
		goto FAIL
	}
	service = services.GetUserService()
	if err = service.ModifyPassword(mod.Email, mod.OldPassword, mod.NewPassword); err != nil {
		goto FAIL
	}
	renderData(context, nil)
	return
FAIL:
	renderError(context, err)
}
