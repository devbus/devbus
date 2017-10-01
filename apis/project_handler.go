package apis

import (
	"strconv"

	"github.com/devbus/devbus/daos"
	"github.com/devbus/devbus/models"
	"github.com/gin-gonic/gin"
	"github.com/kuun/slog"
)

var projectLog = slog.GetLogger()

func GetProject(context *gin.Context) {
	var err error

	var projs []*models.Project
	dao := new(daos.ProjectDAO)
	if id := context.Query("id"); id != "" {
		if i, err := strconv.Atoi(id); err != nil {
			if proj, err := dao.GetById(int32(i)); err != nil {
				projectLog.Info("error: %+v", err)
				goto FAIL
			} else {
				renderData(context, proj)
				return
			}
		}
	} else {
		projs = dao.GetAll()
	}
	renderData(context, projs)
	return
FAIL:
	renderError(context, err)
}

func CreateProject(context *gin.Context) {

}

func ModifyProject(context *gin.Context) {

}

func DeleteProject(context *gin.Context)  {

}

