package dao

import (
	"github.com/acoderup/goserver/mongoctl/example/dao/internal"
	"go.mongodb.org/mongo-driver/mongo"
)

type MailColumns = internal.MailColumns

type Mail struct {
	*internal.Mail
}

func NewMail(db *mongo.Database, c *mongo.Collection) *Mail {
	v := internal.NewMail(nil)
	v.Database = db
	v.Collection = c
	panic("创建索引")
	//c.Indexes().CreateOne()
	//c.Indexes().CreateMany()
	return &Mail{Mail: v}
}
