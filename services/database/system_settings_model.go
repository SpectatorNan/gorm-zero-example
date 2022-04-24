package database

import (
	"fmt"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stringx"
	"gorm.io/gorm"
	"strings"
	"time"
)

var (
	systemSettingsFieldNames          = builder.RawFieldNames(&SystemSettings{}, true)
	systemSettingsRows                = strings.Join(systemSettingsFieldNames, ",")
	systemSettingsRowsExpectAutoSet   = strings.Join(stringx.Remove(systemSettingsFieldNames, "create_time", "update_time"), ",")
	systemSettingsRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(systemSettingsFieldNames, "sys_key", "create_time", "update_time"))

	cachePublicSystemSettingsSysKeyPrefix = "cache:public:systemSettings:sysKey:"
)

type (
	SystemSettingsModel interface {
		Insert(data *SystemSettings) error
		FindOne(sysKey string) (*SystemSettings, error)
		Update(data *SystemSettings) error
		Delete(sysKey string) error
	}

	defaultSystemSettingsModel struct {
		gormc.CachedConn
	}

	SystemSettings struct {
		SysKey     string // 参数key
		SysVal     string // 设置的值
		CreateTime time.Time
		UpdateTime time.Time
		IsDel      bool
	}
)

func NewSystemSettingsModel(conn *gorm.DB, c cache.CacheConf) SystemSettingsModel {
	return &defaultSystemSettingsModel{
		CachedConn: gormc.NewConn(conn, c),
	}
}

func (m *defaultSystemSettingsModel) Insert(data *SystemSettings) error {

	err := m.ExecNoCache(func(conn *gorm.DB) *gorm.DB {
		return conn.Save(data)
	})
	return err
}

func (m *defaultSystemSettingsModel) FindOne(sysKey string) (*SystemSettings, error) {
	publicSystemSettingsSysKeyKey := fmt.Sprintf("%s%v", cachePublicSystemSettingsSysKeyPrefix, sysKey)
	var resp SystemSettings
	err := m.QueryRow(&resp, publicSystemSettingsSysKeyKey, func(conn *gorm.DB) *gorm.DB {
		return conn.Model(&SystemSettings{}).Where("sys_key = ?", sysKey)
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

func (m *defaultSystemSettingsModel) Update(data *SystemSettings) error {
	publicSystemSettingsSysKeyKey := fmt.Sprintf("%s%v", cachePublicSystemSettingsSysKeyPrefix, data.SysKey)
	err := m.Exec(func(conn *gorm.DB) *gorm.DB {
		return conn.Save(data)
	}, publicSystemSettingsSysKeyKey)
	return err
}

func (m *defaultSystemSettingsModel) Delete(sysKey string) error {

	publicSystemSettingsSysKeyKey := fmt.Sprintf("%s%v", cachePublicSystemSettingsSysKeyPrefix, sysKey)
	err := m.Exec(func(conn *gorm.DB) *gorm.DB {
		return conn.Delete(&SystemSettings{}, sysKey)
	}, publicSystemSettingsSysKeyKey)
	return err
}
