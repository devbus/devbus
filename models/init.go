package models

import (
	"fmt"

	"github.com/devbus/devbus/common"
	"github.com/jinzhu/gorm"
	"github.com/kuun/slog"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var log = slog.GetLogger()

func DBOpen() *gorm.DB {

	dbSource := fmt.Sprintf("user=%s password=%s host=%s port=5432 dbname=%s sslmode=disable",
		common.Config.DBUser,
		common.Config.DBPassword,
		common.Config.DBHost,
		common.Config.DB)
	log.Debug("db dbSource: %s", dbSource)
	if db, err := gorm.Open("postgres", dbSource); err != nil {
		log.Errorf("%+v", errors.Wrapf(err, "failed to open database"))
		return nil
	} else {
		return db
	}
}
