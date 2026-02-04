package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// AuthSession holds the schema definition for the AuthSession entity.
type AuthSession struct {
	ent.Schema
}

// Fields of the AuthSession.
func (AuthSession) Fields() []ent.Field {
	return []ent.Field{
		field.String("refresh_token_id").Unique().Immutable().NotEmpty(),
		field.Int("user_id"),
		field.String("refresh_token").Unique(),
		field.String("user_agent").MaxLen(255),
		field.Time("expires_at"),
	}
}

// Edges of the AuthSession.
func (AuthSession) Edges() []ent.Edge {
	return nil
}

func (AuthSession) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("refresh_token_id"),
	}
}

func (AuthSession) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
