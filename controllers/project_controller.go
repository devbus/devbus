package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/devbus/devbus/daos"
	"github.com/devbus/devbus/models"
)

type ProjectController struct {
	jsonApiController
}

func (ctrl *ProjectController) Get() {
	var err error

	var projs []*models.Project
	dao := new(daos.ProjectDAO)
	if id, err := ctrl.GetInt("id"); err == nil {
		if proj, err := dao.GetById(int32(id)); err != nil {
			logs.Info("error: %+v", err)
			goto FAIL
		} else {
			projs = append(projs, proj)
		}
	} else {
		projs = dao.GetAll()
	}
	ctrl.renderData(projs)
	return
FAIL:
	ctrl.renderError(err)
}

func (ctrl *ProjectController) Post() {

}
