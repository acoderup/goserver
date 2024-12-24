package mysqlx

import (
	"errors"

	"github.com/acoderup/goserver/core/mysqlx/internal"
)

var NotInitError = errors.New("mysql manager is nil, please call Init() first")

type Config = internal.Config
type DatabaseConfig = internal.DatabaseConfig
type Database = internal.Database

var manager *internal.Manager

func Init(conf *Config) error {
	manager = internal.NewManager(conf)
	return nil
}

func SetAutoMigrateTables(tables []interface{}) {
	if manager == nil {
		return
	}
	manager.SetAutoMigrateTables(tables)
}

func GetConfig() *Config {
	if manager == nil {
		return nil
	}
	return manager.GetConfig()
}

func Close() {
	if manager == nil {
		return
	}
	manager.Close()
}

func GetDatabase(platform string) (*Database, error) {
	if manager == nil {
		return nil, NotInitError
	}
	return manager.GetDatabase(platform)
}
