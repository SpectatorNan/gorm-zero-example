package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		customUsersLogicModel
	}

	customUsersModel struct {
		*defaultUsersModel
	}

	customUsersLogicModel interface {
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn *gorm.DB, c cache.CacheConf) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c),
	}
}

func (m *defaultUsersModel) customCacheKeys(data *Users) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
