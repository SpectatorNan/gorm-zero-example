package config

import (
	"github.com/SpectatorNan/gorm-zero/config/mysql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql      mysql.Mysql
	CacheRedis cache.CacheConf
}
