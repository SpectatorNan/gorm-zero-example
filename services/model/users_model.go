package model

import (
	"fmt"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"time"
)

var (
	cachePublicUsersIdPrefix = "gormc:cache:public:users:id:"
)

type (
	UsersModel interface {
		Insert(data *Users) error
		FindOne(id int64) (*Users, error)
		Update(data *Users) error
		Delete(id int64) error
	}

	Users struct {
		Account  string
		NickName string
		Password string
		gorm.Model
	}

	defaultUsersModel struct {
		gormc.CachedConn
	}
)

func (u *Users) TableName() string {
	return "users"
}

func NewUsersModel(conn *gorm.DB, c cache.CacheConf) UsersModel {
	return &defaultUsersModel{
		CachedConn: gormc.NewConn(conn, c, func(o *cache.Options) {
			o.NotFoundExpiry = time.Hour
		}),
	}
}

func (m *defaultUsersModel) Insert(data *Users) error {
	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, data.ID)

	err := m.Exec(func(conn *gorm.DB) *gorm.DB {
		return conn.Save(data)
	}, publicUsersIdKey)

	return err
}

func (m *defaultUsersModel) InsertWithoutCache(data *Users) error {

	err := m.ExecNoCache(func(conn *gorm.DB) *gorm.DB {
		return conn.Save(data)
	})

	return err
}

func (m *defaultUsersModel) FindOne(id int64) (*Users, error) {
	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, id)
	var resp Users
	err := m.QueryRow(&resp, publicUsersIdKey, func(conn *gorm.DB) *gorm.DB {
		return conn.Model(&Users{}).Where("id = ?", id)
	})
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneWithoutCache(id int64) (*Users, error) {
	var resp Users
	err := m.QueryRowNoCache(&resp, func(conn *gorm.DB) *gorm.DB {
		return conn.Where("id = ?", id)
	})
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//func (m *defaultUsersModel) FindOneByAccount(account string) (*Users, error) {
//	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, id)
//	var resp Users
//	err := m.QueryRow(&resp, publicUsersIdKey, func(conn *gorm.DB) *gorm.DB {
//		return conn.Model(&Users{}).Where("id = ?", id)
//	})
//	switch err {
//	case nil:
//		return &resp, nil
//	case gorm.ErrRecordNotFound:
//		return nil, ErrNotFound
//	default:
//		return nil, err
//	}
//}
//
//func (m *defaultUsersModel) FindOne(id int64) (*Users, error) {
//	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, id)
//	var resp Users
//	err := m.QueryRow(&resp, publicUsersIdKey, func(conn *gorm.DB) *gorm.DB {
//		return conn.Model(&Users{}).Where("id = ?", id)
//	})
//	switch err {
//	case nil:
//		return &resp, nil
//	case gorm.ErrRecordNotFound:
//		return nil, ErrNotFound
//	default:
//		return nil, err
//	}
//}

func (m *defaultUsersModel) Update(data *Users) error {
	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, data.ID)
	err := m.Exec(func(conn *gorm.DB) *gorm.DB {
		return conn.Save(data)
	}, publicUsersIdKey)
	return err
}

func (m *defaultUsersModel) UpdateWithoutCache(data *Users) error {

	err := m.ExecNoCache(func(conn *gorm.DB) *gorm.DB {
		return conn.Save(data)
	})
	return err
}

func (m *defaultUsersModel) Delete(id int64) error {
	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, id)
	err := m.Exec(func(conn *gorm.DB) *gorm.DB {
		return conn.Delete(&Users{}, id)
	}, publicUsersIdKey)
	return err
}

func (m *defaultUsersModel) DeleteWithoutCache(id int64) error {

	err := m.CachedConn.ExecNoCache(func(conn *gorm.DB) *gorm.DB {
		return conn.Delete(&Users{}, id)
	})
	return err
}
