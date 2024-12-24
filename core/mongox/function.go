package mongox

import (
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/acoderup/goserver/core/logger"
)

type DatabaseType string

const (
	KeyGlobal = "global"

	DatabaseUser    DatabaseType = "user"
	DatabaseLog     DatabaseType = "log"
	DatabaseMonitor DatabaseType = "monitor"
)

// GetClient 获取数据库连接
// 默认获取的是 Global, log 的数据库连接
func GetClient() (*mongo.Client, error) {
	if _manager == nil {
		return nil, NotInitError
	}

	c, err := _manager.GetCollection(KeyGlobal, string(DatabaseLog), "empty")
	if err != nil {
		return nil, err
	}

	return c.Database.Client, nil
}

// GetDatabase 获取数据库
// platform: 平台id
// database: 数据库名称
func GetDatabase(platform string, database DatabaseType) (*Database, error) {
	if _manager == nil {
		return nil, NotInitError
	}

	return _manager.GetDatabase(platform, string(database))
}

func GetUserDatabase(platform string) (*Database, error) {
	return GetDatabase(platform, DatabaseUser)
}

func GetLogDatabase(platform string) (*Database, error) {
	return GetDatabase(platform, DatabaseLog)
}

// GetGlobalDatabase 获取全局库
// database: 数据库名称
func GetGlobalDatabase(database DatabaseType) (*Database, error) {
	if _manager == nil {
		return nil, NotInitError
	}

	return _manager.GetDatabase(KeyGlobal, string(database))
}

func GetGlobalUserDatabase() (*Database, error) {
	return GetGlobalDatabase(DatabaseUser)
}

func GetGlobalLogDatabase() (*Database, error) {
	return GetGlobalDatabase(DatabaseLog)
}

func GetGlobalMonitorDatabase() (*Database, error) {
	return GetGlobalDatabase(DatabaseMonitor)
}

// GetGlobalCollection 获取全局库
// database: 数据库名称
// collection: 集合名称
func GetGlobalCollection(database DatabaseType, collection string) (*Collection, error) {
	if _manager == nil {
		return nil, NotInitError
	}

	return _manager.GetCollection(KeyGlobal, string(database), collection)
}

func GetGlobalUserCollection(collection string) (*Collection, error) {
	return GetGlobalCollection(DatabaseUser, collection)
}

func GetGlobalLogCollection(collection string) (*Collection, error) {
	return GetGlobalCollection(DatabaseLog, collection)
}

func GetGlobalMonitorCollection(collection string) (*Collection, error) {
	return GetGlobalCollection(DatabaseMonitor, collection)
}

// GetCollection 获取平台库
// platform: 平台id
// database: 数据库名称
// collection: 集合名称
func GetCollection(platform string, database DatabaseType, collection string) (*Collection, error) {
	if _manager == nil {
		return nil, NotInitError
	}

	return _manager.GetCollection(platform, string(database), collection)
}

func GetUserCollection(platform string, collection string) (*Collection, error) {
	return GetCollection(platform, DatabaseUser, collection)
}

func GetLogCollection(platform string, collection string) (*Collection, error) {
	return GetCollection(platform, DatabaseLog, collection)
}

// ICollectionName 文档名称接口
type ICollectionName interface {
	CollectionName() string
}

// GetTableName 获取文档名
func GetTableName(model any) string {
	if m, ok := model.(ICollectionName); ok {
		return m.CollectionName()
	}

	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic("model must be a struct or a pointer to a struct")
	}

	return strings.ToLower(t.Name())
}

// GetCollectionDao 获取文档操作接口
// key: 平台id 或 KeyGlobal
// database: 数据库类型 DatabaseType
// f: 文档接口创建函数; 结合 tools/mongoctl 生成
func GetCollectionDao[T any](key string, database DatabaseType, model any, f func(database *mongo.Database, c *mongo.Collection) T) (T, error) {
	collectionName := GetTableName(model)
	c, err := GetCollection(key, database, collectionName)
	if err != nil {
		var z T
		logger.Logger.Errorf("GetCollectionDao key:%v database:%v model:%v error: %v", key, database, collectionName, err)
		return z, err
	}
	return f(c.Database.Database, c.Collection), nil
}
