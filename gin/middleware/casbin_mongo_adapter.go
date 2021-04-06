package middleware

import (
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCasbinMongoAdapter(coll *mongo.Collection) *CasbinMongoAdapter {
	return &CasbinMongoAdapter{
		coll: coll,
	}
}

var _ persist.Adapter = (*CasbinMongoAdapter)(nil)

type CasbinMongoAdapter struct {
	coll *mongo.Collection
}

// LoadPolicy loads all policy rules from the storage.
func (a *CasbinMongoAdapter) LoadPolicy(model model.Model) error {
	return nil
}

// SavePolicy saves all policy rules to the storage.
func (a *CasbinMongoAdapter) SavePolicy(model model.Model) error {
	return nil
}

// AddPolicy adds a policy rule to the storage.
func (a *CasbinMongoAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (a *CasbinMongoAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (a *CasbinMongoAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}
