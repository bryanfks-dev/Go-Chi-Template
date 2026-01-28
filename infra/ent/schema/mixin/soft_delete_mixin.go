package schemamixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SoftDeleteMixin holds the schema definition for the SoftDeleteMixin entity.
type SoftDeleteMixin struct {
	ent.Schema
}

// Fields of the SoftDeleteMixin.
func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("delete_time").Optional().Nillable().Default(nil),
	}
}

// Edges of the SoftDeleteMixin.
func (SoftDeleteMixin) Edges() []ent.Edge {
	return nil
}
