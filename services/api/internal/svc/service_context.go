package svc

import (
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"gorm-zero-example/services/api/internal/config"
	"gorm-zero-example/services/model"
	"log"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	db, err := mysql.Connect(c.Mysql)
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersModel(db, c.CacheRedis),
	}
}
