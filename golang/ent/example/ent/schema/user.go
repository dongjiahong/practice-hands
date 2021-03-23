package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Comment("用户自增id"). // 注释
			Immutable().       // 不可更改
			Min(10000),        // id自增，最小10000
		field.String("phone").
			Comment("用户电话"). // 注释
			Unique().        // 唯一
			NotEmpty(),      // 不为空
		field.String("password").
			NotEmpty(),
		field.Int("p_id").
			Default(0), // 默认为0
		field.String("invited_code").
			NotEmpty(),
		field.Int64("created").
			Default(time.Now().Unix()). // 默认为时间戳
			Immutable(),                // 不变的，只有在创建时设置
		field.Int64("updated").
			Default(0), // 默认为时间戳
		field.Int64("deleted").
			Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("count", UserCount.Type).
			StorageKey(edge.Column("user_id")), // storageKey指定外键为user_id
	}
}
