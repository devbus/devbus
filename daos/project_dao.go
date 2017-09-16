package daos

import (
	"github.com/devbus/devbus/common"
	"github.com/devbus/devbus/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)



type ProjectDAO struct{
	*gorm.DB
}

func NewProjectDAO(db *gorm.DB) *ProjectDAO {
	return &ProjectDAO{db}
}

func (dao *ProjectDAO) Create(proj *models.Project) error {
	if dao.NewRecord(proj) {
		log.Errorf("%+v", errors.Wrapf(dao.Error, "failed to create project, %v", proj))
		return common.OpError{common.ErrProjectCreate}
	}
	return nil
}

func (dao *ProjectDAO) GetAll() []*models.Project {
	projs := make([]*models.Project, 0, 0)
	dao.Find(projs)
	if dao.Error != nil {
		log.Infof("%+v", errors.Wrap(dao.Error, "failed to query project"))
	}
	return projs
}

func (dao *ProjectDAO) GetById(id int32) (*models.Project, error) {

	proj := &models.Project{ID: id}

	if dao.Find(proj); dao.Error != nil {
		log.Infof("%+v", errors.Wrapf(dao.Error, "failed to find project by id: %s", proj.ID))
		err := common.OpError{common.ErrProjectNotFound}
		return nil, err
	}
	return proj, nil
}
