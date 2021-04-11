package repo

import (
	"context"
	"errors"
	casbin_model "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/sfshf/sprout/model"
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
				{"v5", bsonx.Int32(1)},
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

func (a *Casbin) loadPolicyLine(ctx context.Context, line model.Casbin, m casbin_model.Model) {
	var p string
	if line.PType != nil && line.V0 != nil {
		p += *line.PType + ", " + *line.V0
	} else {
		return
	}
	if line.V1 != nil {
		p += ", " + *line.V1
	}
	if line.V2 != nil {
		p += ", " + *line.V2
	}
	if line.V3 != nil {
		p += ", " + *line.V3
	}
	if line.V4 != nil {
		p += ", " + *line.V4
	}
	if line.V5 != nil {
		p += ", " + *line.V5
	}
	persist.LoadPolicyLine(p, m)
}

// LoadPolicy loads all policy rules from the storage.
func (a *Casbin) LoadPolicy(m casbin_model.Model) error {
	ctx := context.Background()
	cursor, err := a.coll.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	var p model.Casbin
	for cursor.Next(ctx) {
		if err := cursor.Decode(&p); err != nil {
			return err
		}
		a.loadPolicyLine(ctx, p, m)
	}
	if err := cursor.Err(); err != nil {
		return err
	}
	return cursor.Close(ctx)
}

func (a *Casbin) savePolicyLine(ctx context.Context, pType string, rule []string) *model.Casbin {
	line := &model.Casbin{
		PType: model.StringPtr(pType),
	}
	if len(rule) > 0 {
		line.V0 = &rule[0]
	}
	if len(rule) > 1 {
		line.V1 = &rule[1]
	}
	if len(rule) > 2 {
		line.V0 = &rule[2]
	}
	if len(rule) > 3 {
		line.V1 = &rule[3]
	}
	if len(rule) > 4 {
		line.V0 = &rule[4]
	}
	if len(rule) > 5 {
		line.V1 = &rule[5]
	}
	return line
}

// SavePolicy saves all policy rules to the storage.
func (a *Casbin) SavePolicy(m casbin_model.Model) error {
	ctx := context.Background()
	if err := a.coll.Drop(ctx); err != nil {
		return err
	}
	var lines []interface{}
	for pType, ast := range m["p"] {
		for _, rule := range ast.Policy {
			line := a.savePolicyLine(ctx, pType, rule)
			lines = append(lines, line)
		}
	}
	for pType, ast := range m["g"] {
		for _, rule := range ast.Policy {
			line := a.savePolicyLine(ctx, pType, rule)
			lines = append(lines, line)
		}
	}
	if _, err := a.coll.InsertMany(ctx, lines); err != nil {
		return err
	}
	return nil
}

// AddPolicy adds a policy rule to the storage.
func (a *Casbin) AddPolicy(sec string, pType string, rule []string) error {
	return errors.New("not implemented")
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (a *Casbin) RemovePolicy(sec string, pType string, rule []string) error {
	return errors.New("not implemented")
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (a *Casbin) RemoveFilteredPolicy(sec string, pType string, fieldIndex int, fieldValues ...string) error {
	return errors.New("not implemented")
}
