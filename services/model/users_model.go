package model

import (
	"context"
	"fmt"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"time"
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
		FindOneWithExpire(ctx context.Context, id int64, expire time.Duration) (*Users, error)
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn *gorm.DB, c cache.CacheConf) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c),
	}
}

var (
	cacheGormzeroUsersIdExpirePrefix = "cache:gormzero:users:id:expire:"
)

func (m *defaultUsersModel) customCacheKeys(data *Users) []string {
	if data == nil {
		return []string{}
	}
	return []string{
		fmt.Sprintf("%s%v", cacheGormzeroUsersIdExpirePrefix, data.Id),
	}
}

func (m *customUsersModel) FindOneWithExpire(ctx context.Context, id int64, expire time.Duration) (*Users, error) {
	gormzeroUsersIdKey := fmt.Sprintf("%s%v", cacheGormzeroUsersIdExpirePrefix, id)
	var resp Users
	err := m.QueryWithExpireCtx(ctx, &resp, gormzeroUsersIdKey, expire, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Users{}).Where("`id` = ?", id).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
