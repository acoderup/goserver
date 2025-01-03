// --------------------------------------------------------------------------------------------
// The following code is automatically generated by the mongo-dao-generator tool.
// Please do not modify this code manually to avoid being overwritten in the next generation.
// For more tool details, please click the link to view https://github.com/dobyte/mongo-dao-generator
// --------------------------------------------------------------------------------------------

package internal

import (
	"context"
	"errors"
	modelpkg "github.com/acoderup/goserver/mongoctl/example/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type UserFilterFunc func(cols *UserColumns) interface{}
type UserUpdateFunc func(cols *UserColumns) interface{}
type UserPipelineFunc func(cols *UserColumns) interface{}
type UserCountOptionsFunc func(cols *UserColumns) *options.CountOptions
type UserAggregateOptionsFunc func(cols *UserColumns) *options.AggregateOptions
type UserFindOneOptionsFunc func(cols *UserColumns) *options.FindOneOptions
type UserFindManyOptionsFunc func(cols *UserColumns) *options.FindOptions
type UserUpdateOptionsFunc func(cols *UserColumns) *options.UpdateOptions
type UserDeleteOptionsFunc func(cols *UserColumns) *options.DeleteOptions
type UserInsertOneOptionsFunc func(cols *UserColumns) *options.InsertOneOptions
type UserInsertManyOptionsFunc func(cols *UserColumns) *options.InsertManyOptions

type User struct {
	Columns    *UserColumns
	Database   *mongo.Database
	Collection *mongo.Collection
}

type UserColumns struct {
	ID             string
	UID            string // 用户ID
	Account        string // 用户账号
	Password       string // 用户密码
	Salt           string // 密码
	Mobile         string // 用户手机
	Email          string // 用户邮箱
	Nickname       string // 用户昵称
	Signature      string // 用户签名
	Gender         string // 用户性别
	Level          string // 用户等级
	Experience     string // 用户经验
	Coin           string // 用户金币
	Type           string // 用户类型
	Status         string // 用户状态
	DeviceID       string // 设备ID
	ThirdPlatforms string // 第三方平台
	RegisterIP     string // 注册IP
	RegisterTime   string // 注册时间
	LastLoginIP    string // 最近登录IP
	LastLoginTime  string // 最近登录时间
}

var userColumns = &UserColumns{
	ID:             "_id",
	UID:            "uid",             // 用户ID
	Account:        "account",         // 用户账号
	Password:       "password",        // 用户密码
	Salt:           "salt",            // 密码
	Mobile:         "mobile",          // 用户手机
	Email:          "email",           // 用户邮箱
	Nickname:       "nickname",        // 用户昵称
	Signature:      "signature",       // 用户签名
	Gender:         "gender",          // 用户性别
	Level:          "level",           // 用户等级
	Experience:     "experience",      // 用户经验
	Coin:           "coin",            // 用户金币
	Type:           "type",            // 用户类型
	Status:         "status",          // 用户状态
	DeviceID:       "device_id",       // 设备ID
	ThirdPlatforms: "third_platforms", // 第三方平台
	RegisterIP:     "register_ip",     // 注册IP
	RegisterTime:   "register_time",   // 注册时间
	LastLoginIP:    "last_login_ip",   // 最近登录IP
	LastLoginTime:  "last_login_time", // 最近登录时间
}

func NewUser(db *mongo.Database) *User {
	return &User{
		Columns:    userColumns,
		Database:   db,
		Collection: db.Collection("user"),
	}
}

// Count returns the number of documents in the collection.
func (dao *User) Count(ctx context.Context, filterFunc UserFilterFunc, optionsFunc ...UserCountOptionsFunc) (int64, error) {
	var (
		opts   *options.CountOptions
		filter = filterFunc(dao.Columns)
	)

	if len(optionsFunc) > 0 {
		opts = optionsFunc[0](dao.Columns)
	}

	return dao.Collection.CountDocuments(ctx, filter, opts)
}

// Aggregate executes an aggregate command against the collection and returns a cursor over the resulting documents.
func (dao *User) Aggregate(ctx context.Context, pipelineFunc UserPipelineFunc, optionsFunc ...UserAggregateOptionsFunc) (*mongo.Cursor, error) {
	var (
		opts     *options.AggregateOptions
		pipeline = pipelineFunc(dao.Columns)
	)

	if len(optionsFunc) > 0 {
		opts = optionsFunc[0](dao.Columns)
	}

	return dao.Collection.Aggregate(ctx, pipeline, opts)
}

// InsertOne executes an insert command to insert a single document into the collection.
func (dao *User) InsertOne(ctx context.Context, model *modelpkg.User, optionsFunc ...UserInsertOneOptionsFunc) (*mongo.InsertOneResult, error) {
	if model == nil {
		return nil, errors.New("model is nil")
	}

	if err := dao.autofill(ctx, model); err != nil {
		return nil, err
	}

	var opts *options.InsertOneOptions

	if len(optionsFunc) > 0 {
		opts = optionsFunc[0](dao.Columns)
	}

	return dao.Collection.InsertOne(ctx, model, opts)
}

// InsertMany executes an insert command to insert multiple documents into the collection.
func (dao *User) InsertMany(ctx context.Context, models []*modelpkg.User, optionsFunc ...UserInsertManyOptionsFunc) (*mongo.InsertManyResult, error) {
	if len(models) == 0 {
		return nil, errors.New("models is empty")
	}

	documents := make([]interface{}, 0, len(models))
	for i := range models {
		model := models[i]
		if err := dao.autofill(ctx, model); err != nil {
			return nil, err
		}
		documents = append(documents, model)
	}

	var opts *options.InsertManyOptions

	if len(optionsFunc) > 0 {
		opts = optionsFunc[0](dao.Columns)
	}

	return dao.Collection.InsertMany(ctx, documents, opts)
}

// UpdateOne executes an update command to update at most one document in the collection.
func (dao *User) UpdateOne(ctx context.Context, filterFunc UserFilterFunc, updateFunc UserUpdateFunc, optionsFunc ...UserUpdateOptionsFunc) (*mongo.UpdateResult, error) {
	var (
		opts   *options.UpdateOptions
		filter = filterFunc(dao.Columns)
		update = updateFunc(dao.Columns)
	)

	if len(optionsFunc) > 0 {
		opts = optionsFunc[0](dao.Columns)
	}

	return dao.Collection.UpdateOne(ctx, filter, update, opts)
}

// UpdateOneByID executes an update command to update at most one document in the collection.
func (dao *User) UpdateOneByID(ctx context.Context, id string, updateFunc UserUpdateFunc, optionsFunc ...UserUpdateOptionsFunc) (*mongo.UpdateResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return dao.UpdateOne(ctx, func(cols *UserColumns) interface{} {
		return bson.M{"_id": objectID}
	}, updateFunc, optionsFunc...)
}

// UpdateMany executes an update command to update documents in the collection.
func (dao *User) UpdateMany(ctx context.Context, filterFunc UserFilterFunc, updateFunc UserUpdateFunc, optionsFunc ...UserUpdateOptionsFunc) (*mongo.UpdateResult, error) {
	var (
		opts   *options.UpdateOptions
		filter = filterFunc(dao.Columns)
		update = updateFunc(dao.Columns)
	)

	if len(optionsFunc) > 0 {
		opts = optionsFunc[0](dao.Columns)
	}

	return dao.Collection.UpdateMany(ctx, filter, update, opts)
}

// FindOne executes a find command and returns a model for one document in the collection.
func (dao *User) FindOne(ctx context.Context, filterFunc UserFilterFunc, optionsFunc ...UserFindOneOptionsFunc) (*modelpkg.User, error) {
	var (
		opts   *options.FindOneOptions
		model  = &modelpkg.User{}
		filter = filterFunc(dao.Columns)
	)

	if len(optionsFunc) > 0 {
		opts = optionsFunc[0](dao.Columns)
	}

	err := dao.Collection.FindOne(ctx, filter, opts).Decode(model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return model, nil
}

// FindOneByID executes a find command and returns a model for one document in the collection.
func (dao *User) FindOneByID(ctx context.Context, id string, optionsFunc ...UserFindOneOptionsFunc) (*modelpkg.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return dao.FindOne(ctx, func(cols *UserColumns) interface{} {
		return bson.M{"_id": objectID}
	}, optionsFunc...)
}

// FindMany executes a find command and returns many models the matching documents in the collection.
func (dao *User) FindMany(ctx context.Context, filterFunc UserFilterFunc, optionsFunc ...UserFindManyOptionsFunc) ([]*modelpkg.User, error) {
	var (
		opts   *options.FindOptions
		filter = filterFunc(dao.Columns)
	)

	if len(optionsFunc) > 0 {
		opts = optionsFunc[0](dao.Columns)
	}

	cur, err := dao.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	models := make([]*modelpkg.User, 0)

	if err = cur.All(ctx, &models); err != nil {
		return nil, err
	}

	return models, nil
}

// DeleteOne executes a delete command to delete at most one document from the collection.
func (dao *User) DeleteOne(ctx context.Context, filterFunc UserFilterFunc, optionsFunc ...UserDeleteOptionsFunc) (*mongo.DeleteResult, error) {
	var (
		opts   *options.DeleteOptions
		filter = filterFunc(dao.Columns)
	)

	if len(optionsFunc) > 0 {
		opts = optionsFunc[0](dao.Columns)
	}

	return dao.Collection.DeleteOne(ctx, filter, opts)
}

// DeleteOneByID executes a delete command to delete at most one document from the collection.
func (dao *User) DeleteOneByID(ctx context.Context, id string, optionsFunc ...UserDeleteOptionsFunc) (*mongo.DeleteResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return dao.DeleteOne(ctx, func(cols *UserColumns) interface{} {
		return bson.M{"_id": objectID}
	}, optionsFunc...)
}

// DeleteMany executes a delete command to delete documents from the collection.
func (dao *User) DeleteMany(ctx context.Context, filterFunc UserFilterFunc, optionsFunc ...UserDeleteOptionsFunc) (*mongo.DeleteResult, error) {
	var (
		opts   *options.DeleteOptions
		filter = filterFunc(dao.Columns)
	)

	if len(optionsFunc) > 0 {
		opts = optionsFunc[0](dao.Columns)
	}

	return dao.Collection.DeleteMany(ctx, filter, opts)
}

// autofill when inserting data
func (dao *User) autofill(ctx context.Context, model *modelpkg.User) error {
	if model.ID.IsZero() {
		model.ID = primitive.NewObjectID()
	}

	if model.UID == 0 {
		if id, err := NewCounter(dao.Database).Incr(ctx, "uid"); err != nil {
			return err
		} else {
			model.UID = int32(id)
		}
	}

	if model.RegisterTime == 0 {
		model.RegisterTime = primitive.NewDateTimeFromTime(time.Now())
	}

	if model.LastLoginTime == 0 {
		model.LastLoginTime = primitive.NewDateTimeFromTime(time.Now())
	}

	return nil
}
