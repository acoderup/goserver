package template

const CounterExternalTemplate = `
package ${VarDaoPackageName}

import (
	"go.mongodb.org/mongo-driver/mongo"
	"${VarDaoPackagePath}/internal"
)

type ${VarDaoClassName} struct {
	*internal.${VarDaoClassName}
}

func New${VarDaoClassName}(db *mongo.Database) *${VarDaoClassName} {
	return &${VarDaoClassName}{${VarDaoClassName}: internal.New${VarDaoClassName}(db)}
}
`

const CounterInternalTemplate = `
// --------------------------------------------------------------------------------------------------
// The following code is automatically generated by the mongo-dao-generator tool.
// Please do not modify this code manually to avoid being overwritten in the next generation. 
// For more tool details, please click the link to view https://github.com/dobyte/mongo-dao-generator
// --------------------------------------------------------------------------------------------------

package internal

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ${VarDaoClassName} struct {
	Columns    *${VarDaoPrefixName}Columns
	Database   *mongo.Database
	Collection *mongo.Collection
}

type ${VarDaoPrefixName}Model struct {
    ID    string ${SymbolBacktick}bson:"_id"${SymbolBacktick}
    Value int64  ${SymbolBacktick}bson:"value"${SymbolBacktick}
}

type ${VarDaoPrefixName}Columns struct {
	ID    string
	Value string
}

var ${VarDaoVariableName}Columns = &${VarDaoPrefixName}Columns{
	ID:    "_id",
	Value: "value",
}

func New${VarDaoClassName}(db *mongo.Database) *${VarDaoClassName} {
	return &${VarDaoClassName}{
		Columns:    ${VarDaoVariableName}Columns,
		Database:   db,
		Collection: db.Collection("${VarCollectionName}"),
	}
}

// Incr 自增值
func (dao *${VarDaoClassName}) Incr(ctx context.Context, key string, incr ...int) (int64, error) {
	var (
		upsert         = true
		returnDocument = options.After
		counter        = &${VarDaoPrefixName}Model{}
		value          = 1
	)

	if len(incr) > 0 {
		if incr[0] == 0 {
			return 0, errors.New("invalid increment value")
		}
		value = incr[0]
	}

	rst := dao.Collection.FindOneAndUpdate(ctx, bson.M{
		dao.Columns.ID: key,
	}, bson.M{"$inc": bson.M{
		dao.Columns.Value: value,
	}}, &options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &returnDocument,
	})

	if err := rst.Decode(counter); err != nil {
		return 0, err
	}

	return counter.Value, nil
}
`
