package internal

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/acoderup/goserver/core/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Global    map[string]*DatabaseConfig
	Platforms map[string]map[string]*DatabaseConfig
}

type DatabaseConfig struct {
	HostName string // 主机地址
	HostPort int32  // 端口
	Database string // 数据库名
	Username string // 用户名
	Password string // 密码
	Options  string // 配置
}

type Collection struct {
	Database *Database
	*mongo.Collection
}

type Database struct {
	*DatabaseConfig
	Client     *mongo.Client
	Database   *mongo.Database
	Collection sync.Map
}

func (d *Database) Connect() error {
	if d.DatabaseConfig == nil {
		err := fmt.Errorf("mongo Connect error, DatabaseConifg not found")
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
		port = 27017
	}
	myOptions := d.DatabaseConfig.Options
	if myOptions != "" {
		myOptions = "?" + myOptions
	}

	s := fmt.Sprintf("mongodb://%s%s:%d/admin%s", login, host, port, myOptions)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(s))
	if err != nil {
		logger.Logger.Errorf("mongo Connect %v error: %v config:%+v", s, err, *d.DatabaseConfig)
		return err
	}
	logger.Logger.Tracef("mongo connect success %+v", *d.DatabaseConfig)
	d.Client = client
	d.Database = client.Database(d.DatabaseConfig.Database)
	return nil
}

func (d *Database) GetCollection(name string) (*Collection, error) {
	if d.Database == nil {
		err := fmt.Errorf("mongo GetCollection error, collection:%v, database is nil", name)
		logger.Logger.Error(err)
		return nil, err
	}
	v, ok := d.Collection.Load(name)
	if !ok {
		v = &Collection{
			Database:   d,
			Collection: d.Database.Collection(name),
		}
		d.Collection.Store(name, v)
	}
	c, _ := v.(*Collection)
	return c, nil
}

type Manager struct {
	conf      *Config
	global    *sync.Map // 内部库名称:Database
	platforms *sync.Map // 平台id:内部库名称:Database
}

func (m *Manager) GetCollection(key, database, collection string) (*Collection, error) {
	d, err := m.GetDatabase(key, database)
	if err != nil {
		return nil, err
	}
	return d.GetCollection(collection)
}

func (m *Manager) GetDatabase(key, database string) (*Database, error) {
	switch key {
	case "global":
		v, ok := m.global.Load(database)
		if !ok {
			db := &Database{
				DatabaseConfig: m.conf.Global[database],
				Collection:     sync.Map{},
			}
			if err := db.Connect(); err != nil {
				return nil, err
			}
			v = db
			m.global.Store(database, v)
		}
		d, _ := v.(*Database)
		return d, nil

	default:
		var mp *sync.Map
		v, ok := m.platforms.Load(key) // 平台id
		if !ok {
			mp = new(sync.Map)
			m.platforms.Store(key, mp)
		} else {
			mp = v.(*sync.Map)
		}
		v, ok = mp.Load(database)
		if !ok {
			db := &Database{
				DatabaseConfig: m.conf.Platforms[key][database],
				Collection:     sync.Map{},
			}
			if err := db.Connect(); err != nil {
				return nil, err
			}
			v = db
			mp.Store(database, v)
		}
		d, _ := v.(*Database)
		return d, nil
	}
}

func (m *Manager) Restart(conf *Config) {
	logger.Logger.Infof("mongo manager restart...")
	old := *m
	time.AfterFunc(time.Minute, func() {
		Close(&old)
	})

	m.conf = conf
	m.global = &sync.Map{}
	m.platforms = &sync.Map{}
}

func Close(m *Manager) {
	logger.Logger.Infof("mongo manager close")
	m.global.Range(func(key, value any) bool {
		if v, ok := value.(*Database); ok {
			v.Client.Disconnect(nil)
		}
		return true
	})

	m.platforms.Range(func(key, value any) bool {
		if v, ok := value.(*sync.Map); ok {
			v.Range(func(key, value any) bool {
				if v, ok := value.(*Database); ok {
					v.Client.Disconnect(nil)
				}
				return true
			})
		}
		return true
	})
}

func (m *Manager) GetConfig() *Config {
	return m.conf
}

func NewManager(conf *Config) *Manager {
	return &Manager{
		conf:      conf,
		global:    &sync.Map{},
		platforms: &sync.Map{},
	}
}
