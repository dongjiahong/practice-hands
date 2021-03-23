package models

import (
	"context"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"sqlent/ent"
	"sqlent/ent/hook"
)

var client *ent.Client

func Init() {
	var err error
	//client, err = ent.Open("mysql", "root:@tcp(localhost:3306)/db_ent", ent.Debug())
	client, err = ent.Open("mysql", "root:@tcp(localhost:3306)/db_ent")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v\n", err)
	}
	//defer client.Close()
	ctx := context.Background()

	// 使用钩子更新updated字段
	client.Use(hook.On(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			created := time.Now().Unix()
			if err := m.SetField("updated", created); err != nil {
				// an error is returned, if the field is not defined in
				// the schema, or if the type mismatch the field type.
				log.Println("==> update err: ", err)
			}
			return next.Mutate(ctx, m)
		})
	}, ent.OpUpdateOne|ent.OpCreate)) // 每当更新字段时，更新updated字段

	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			start := time.Now()
			defer func() {
				log.Printf("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
			}()
			return next.Mutate(ctx, m)
		})
	})

	if err := client.Schema.Create(ctx); err != nil {
		panic(err)
	}
}

func Close() {
	if client != nil {
		client.Close()
	}
}

func GetClient() *ent.Client {
	return client
}
