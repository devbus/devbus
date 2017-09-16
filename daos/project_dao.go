package daos

import (
	"github.com/astaxie/beego/logs"
	"github.com/devbus/devbus/common"
	"github.com/astaxie/beego/orm"
	"github.com/devbus/devbus/models"
	"github.com/pkg/errors"
)

type ProjectDAO struct{}

func (ProjectDAO) Create(proj *models.Project) error {
	o := orm.NewOrm()
	if _, err := o.Insert(proj); err != nil {
		logs.Error("%+v", errors.Wrapf(err, "failed to create project, %v", proj))
		return common.OpError{common.ErrProjectCreate}
	}
	return nil
}

func (ProjectDAO) GetAll() []*models.Project {
	o := orm.NewOrm()
	projs := make([]*models.Project, 0, 0)
	_, err := o.QueryTable("project").All(projs)
	if err != nil {
		logs.Info("%+v", errors.Wrap(err, "failed to query project"))
	}
	return projs
}

func (ProjectDAO) GetById(id int32) (*models.Project, error) {
	o := orm.NewOrm()
	proj := &models.Project{Id: id}
	err := o.Read(proj)
	if err != nil {
		logs.Info("%+v", errors.Wrapf(err, "failed to find project by id: %s", proj.Id))
		err = common.OpError{common.ErrProjectNotFound}
		return nil, err
	}
	return proj, nil
}
