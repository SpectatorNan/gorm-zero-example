package svc

import (
	"github.com/SpectatorNan/gorm-zero/config/mysql"
	"gorm-zero-example/services/api/internal/config"
	"gorm-zero-example/services/model"
	"gorm-zero-example/services/model_noCache"
	"log"
)

type ServiceContext struct {
	Config           config.Config
	UserCacheModel   model.UsersModel
	UserNoCacheModel model_noCache.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	db, err := mysql.Connect(c.Mysql)
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:           c,
		UserCacheModel:   model.NewUsersModel(db, c.CacheRedis),
		UserNoCacheModel: model_noCache.NewUsersModel(db),
	}
}
