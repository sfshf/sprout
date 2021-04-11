package repo

import (
	"context"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func CasbinRepo() *Casbin {
	return _casbin
}

func InitCasbinRepo(ctx context.Context, db *mongo.Database) {
	_casbin = &Casbin{
		coll: db.Collection(casbinCollName),
	}
	_casbin.coll.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{
				{"pType", bsonx.Int32(1)},
				{"v0", bsonx.Int32(1)},
				{"v1", bsonx.Int32(1)},
				{"v2", bsonx.Int32(1)},
				{"v3", bsonx.Int32(1)},
				{"v4", bsonx.Int32(1)},
			},
			Options: options.Index().SetUnique(true),
		},
	})
}

var (
	_casbin *Casbin

	casbinCollName = "casbin"
)

// A implementation of Adapter, BatchAdapter, FilteredAdapter interfaces of github.com/casbin/casbin/v2/persist package.
type Casbin struct {
	coll *mongo.Collection
}

func (a *Casbin) Collection() *mongo.Collection {
	return a.coll
}

var _ persist.Adapter = (*Casbin)(nil)

// LoadPolicy loads all policy rules from the storage.
func (a *Casbin) LoadPolicy(model model.Model) error {
	return nil
}

// SavePolicy saves all policy rules to the storage.
func (a *Casbin) SavePolicy(model model.Model) error {
	return nil
}

// AddPolicy adds a policy rule to the storage.
func (a *Casbin) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (a *Casbin) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (a *Casbin) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}
