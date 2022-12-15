package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UnameModel = (*customUnameModel)(nil)

type (
	// UnameModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUnameModel.
	UnameModel interface {
		unameModel
		customUnameLogicModel
	}

	customUnameModel struct {
		*defaultUnameModel
	}

	customUnameLogicModel interface {
	}
)

// NewUnameModel returns a model for the database table.
func NewUnameModel(conn *gorm.DB, c cache.CacheConf) UnameModel {
	return &customUnameModel{
		defaultUnameModel: newUnameModel(conn, c),
	}
}

func (m *defaultUnameModel) customCacheKeys(data *Uname) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
