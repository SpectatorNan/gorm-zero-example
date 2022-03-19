package svc

import (
	"gorm-zero-example/services/api/internal/config"
	"gorm-zero-example/services/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	db, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersModel(db, c.CacheRedis),
	}
}
