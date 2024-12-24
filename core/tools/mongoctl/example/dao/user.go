package dao

import (
	"github.com/acoderup/goserver/mongoctl/example/dao/internal"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserColumns = internal.UserColumns

type User struct {
	*internal.User
}

func NewUser(db *mongo.Database, c *mongo.Collection) *User {
	v := internal.NewUser(nil)
	v.Database = db
	v.Collection = c
	panic("创建索引")
	//c.Indexes().CreateOne()
	//c.Indexes().CreateMany()
	return &User{User: v}
}
