package internal

import (
	"fmt"
	"sync"
	"time"

	"github.com/acoderup/goserver/core/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Platforms       map[string]*DatabaseConfig
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
	ConnMaxIdletime int
}

type DatabaseConfig struct {
	HostName string
	HostPort int32
	Database string
	Username string
	Password string
	Options  string
}

type Database struct {
	*Manager
	*Config
	*DatabaseConfig
	*gorm.DB
}

func (d *Database) Connect() error {
	if d.DatabaseConfig == nil {
		err := fmt.Errorf("mysql Connect error, DatabaseConifg not found")
		logger.Logger.Error(err)
		return err
	}

	login := ""
	if d.DatabaseConfig.Username != "" {
		login = d.DatabaseConfig.Username + ":" + d.DatabaseConfig.Password + "@"
	}
	host := d.DatabaseConfig.HostName
	if d.DatabaseConfig.HostName == "" {
		host = "127.0.0.1"
	}
	port := d.DatabaseConfig.HostPort
	if d.DatabaseConfig.HostPort == 0 {
		port = 3306
	}
	database := d.DatabaseConfig.Database
	if database == "" {
		database = "mysql"
	}
	myOptions := d.DatabaseConfig.Options
	if myOptions != "" {
		myOptions = "?" + myOptions
	}

	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	s := fmt.Sprintf("%stcp(%s:%d)/%s%s", login, host, port, "mysql", myOptions)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})
	if err != nil {
		logger.Logger.Errorf("mysql Connect %v error: %v config:%+v", s, err, *d.DatabaseConfig)
		return err
	}
	logger.Logger.Tracef("mysql connect success %+v", *d.DatabaseConfig)

	err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci", d.DatabaseConfig.Database)).Error
	if err != nil {
		logger.Logger.Errorf("mysql create database %s error: %v", d.DatabaseConfig.Database, err)
		return err
	}

	s = fmt.Sprintf("%stcp(%s:%d)/%s%s", login, host, port, d.DatabaseConfig.Database, myOptions)
	db, err = gorm.Open(mysql.Open(s), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		logger.Logger.Errorf("mysql Connect %v error: %v config:%+v", s, err, *d.DatabaseConfig)
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Logger.Errorf("mysql get DB error: %v", err)
		return err
	}

	if len(d.tables) > 0 {
		if err := db.AutoMigrate(d.tables...); err != nil {
			logger.Logger.Warnf("mysql migrate error: %v", err)
		}
	}

	sqlDB.SetMaxIdleConns(d.MaxIdleConns)
	sqlDB.SetMaxOpenConns(d.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(d.ConnMaxLifetime))
	sqlDB.SetConnMaxIdleTime(time.Duration(d.ConnMaxIdletime))

	d.DB = db.Session(&gorm.Session{SkipDefaultTransaction: true})

	return nil
}

type Manager struct {
	conf      *Config
	platforms sync.Map // 平台id:Database
	tables    []interface{}
}

func (m *Manager) GetConfig() *Config {
	return m.conf
}

func (m *Manager) SetAutoMigrateTables(tables []interface{}) {
	m.tables = tables
}

func (m *Manager) GetDatabase(key string) (*Database, error) {
	v, ok := m.platforms.Load(key) // 平台id
	if !ok {
		db := &Database{
			Manager:        m,
			Config:         m.conf,
			DatabaseConfig: m.conf.Platforms[key],
		}
		if err := db.Connect(); err != nil {
			return nil, err
		}
		v = db
		m.platforms.Store(key, v)
	}
	d, _ := v.(*Database)
	return d, nil
}

func (m *Manager) Close() {

}

func NewManager(conf *Config) *Manager {
	return &Manager{
		conf:      conf,
		platforms: sync.Map{},
	}
}
