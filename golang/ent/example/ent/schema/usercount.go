package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserCount holds the schema definition for the UserCount entity.
type UserCount struct {
	ent.Schema
}

// Fields of the UserCount.
func (UserCount) Fields() []ent.Field {
	return []ent.Field{
		field.Float("self_buy").
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(18,8)", // 类型映射
			}).
			Default(0),
		field.Float("invite_buy").
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(18,8)", // 类型映射
			}).
			Default(0),
		field.Int("level").
			Default(0),
		field.Int64("created").
			Default(time.Now().Unix()). // 默认为时间戳
			Immutable(),                // 不变的，只有在创建时设置
		field.Int64("updated").
			Default(0), // 默认为时间戳
		field.Int64("deleted").
			Default(0),
	}
}

// Edges of the UserCount.
func (UserCount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("count").
			Unique(), // 表示每个count只属于一个User
	}
}
