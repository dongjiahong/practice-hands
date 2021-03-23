package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserBuyRecord holds the schema definition for the UserBuyRecord entity.
type UserBuyRecord struct {
	ent.Schema
}

// Fields of the UserBuyRecord.
func (UserBuyRecord) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Comment("订单id").
			Default(uuid.New),
		field.Int("power").
			Comment("每份算力(T)").
			Default(0),
		field.Int("power_num").
			Comment("购买的分数").
			Positive().
			Default(0),
		field.Float("total_power").
			Comment("购买的总算力").
			SchemaType(map[string]string{ // 映射到mysql为decimal类型
				dialect.MySQL: "decimal(18.8)",
			}),
		field.Int("total_day").
			Comment("合约的总天数").
			Min(0), // 最小值
		field.Int("remain_day").
			Comment("合约剩余天数").
			Min(0),
		field.String("node").
			Comment("节点号").
			NotEmpty(),
		field.Float("used_usdt").
			Comment("消费的usdt").
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(18.8)",
			}),
		field.String("buy_date").
			Comment("购买日期").
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

// Edges of the UserBuyRecord.
func (UserBuyRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("buy_record").
			Unique(), // 记录属于一个用户
	}
}
