package main

import (
	"fmt"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// 定义我们的产品数据结构，有两个字段分别是id、和name
type Product struct {
	id   string
	name string
}

// 要测试的函数，主要记录的是产品被浏览的次数，并插入一条产品浏览日志
func recordStats(db *sql.DB, userID, productID int64) (err error) {
	// 这里我们选择了事务
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit() // 提交事务
		default:
			tx.Rollback() // 回滚事务
		}
	}()
	if _, err = tx.Exec("update products set views = views+1"); err != nil {
		return
	}
	if _, err = tx.Exec("insert into product_viewers(user_id, product_id) values(?,?)", userID, productID); err != nil {
		return
	}
	return
}

// 这里要测试的是产品的查询功能
func GetSingleProductDao(db *sql.DB, id string) (*Product, error) {
	p := Product{}
	err := db.QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(&p.id, &p.name)
	if err != nil {
		return nil, fmt.Errorf("query err")
	}
	return &p, nil
}

func main() {
	db, err := sql.Open("mysql", "root@/blog")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = recordStats(db, 1, 5); err != nil {
		panic(err)
	}

}
