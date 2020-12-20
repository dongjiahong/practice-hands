package main

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestShouldUpdateStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("update products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("insert into product_viewers").
		WithArgs(2, 3).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err := recordStats(db, 2, 3); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldRollbackStatUpdatesStatUpdatesOnFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectBegin()
	mock.ExpectExec("update products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("insert into product_viewers").
		WithArgs(2, 3).
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	if err := recordStats(db, 2, 3); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestShouldGetSingleProductDao(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// 这里碰到一个问题，上面的函数中sql使用的是“事务”，所以测试里也使用“事务”来模拟捕获
	// 如mock.ExpectQuery和mock.ExpectCommit；而这里使用的不是事务，而是普通的查询语句，
	// 所以这里直接使用了mock.ExpectQuery，如果前后不对应是不能mock数据的
	productRows := sqlmock.NewRows([]string{"id", "name"}).AddRow("22", "headset")
	mock.ExpectQuery("SELECT (.+) FROM products WHERE id = \\?").
		WithArgs("22").
		WillReturnRows(productRows)

	t.Log("====> productRows: ", productRows)

	expectdProduct := &Product{
		id:   "22",
		name: "headset",
	}
	t.Log("====> expectdProduct: ", expectdProduct)

	product, _ := GetSingleProductDao(db, "22")

	t.Log("====> product: ", product)

	assert.Equal(t, expectdProduct, product)
}
