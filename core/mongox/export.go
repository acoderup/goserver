package mongox

import (
	"errors"

	"github.com/acoderup/goserver/core/logger"

	"github.com/acoderup/goserver/core/mongox/internal"
)

var NotInitError = errors.New("mongo manager is nil, please call Init() first")

type Config = internal.Config
type DatabaseConfig = internal.DatabaseConfig
type Collection = internal.Collection
type Database = internal.Database

var _manager *internal.Manager

// GetConfig 获取配置
func GetConfig() *Config {
	if _manager == nil {
		return nil
	}
	return _manager.GetConfig()
}

// Init 初始化
func Init(conf *Config) {
	_manager = internal.NewManager(conf)
}

// Restart 重启
func Restart() {
	if _manager == nil {
		logger.Logger.Error(NotInitError)
		return
	}
	_manager.Restart(_manager.GetConfig())
}

// Close 关闭
func Close() {
	internal.Close(_manager)
}
