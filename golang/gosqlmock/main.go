package main

import (
	"fmt"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	id   string
	name string
}

func recordStats(db *sql.DB, userID, productID int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
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
