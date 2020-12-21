package main

import (
	"fmt"
	"testing"

	// 引入go-sqlmock包
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// 测试recordStats函数
func TestShouldUpdateStats(t *testing.T) {
	// 创建一个sqlmock的db句柄
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close() // 记得关闭

	// mock的事务，这里对应于recordStats函数，应为它里面也是使用了事务
	mock.ExpectBegin()
	// 捕获事务“update products”开头的sql语句，并返回结果，下面的sqlmock.NewResult(1,1)
	// 第一个1表示最后一个受影响的id；第二个1表示一共有多少数据受影响
	mock.ExpectExec("update products").WillReturnResult(sqlmock.NewResult(1, 1))
	// 捕获事务"insert into product_viewers"开头的sql语句并支持"."语法糖
	// WithArgs(2,3)表示，在捕获的sql语句"insert into product_viewers"中给两个参数为2和3 ，
	//   由于要测试的原sql为“insert into product_viewers(user_id, product_id) values(?,?)”,
	//   所以这里的2表示传入的user_id,3表示product_id
	// WillReturnResult的含义和上面的一样，主要是表示sql语句执行成功
	mock.ExpectExec("insert into product_viewers").
		WithArgs(2, 3).
		WillReturnResult(sqlmock.NewResult(1, 1))
	// 提交事务
	mock.ExpectCommit()

	// 调用要测试的函数，并传入我们的mock数据db
	if err := recordStats(db, 2, 3); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// 这里的ExpectationsWereMet函数是的功能是检查我们mock的sql是否匹配到原本待测是sql
	// 当测试不如预期时，有助于我们了解是否是由于捕获出问题
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// 跟上面的大体一样，不过是测试当db返回错误信息的情况
func TestShouldRollbackStatUpdatesStatUpdatesOnFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("update products").WillReturnResult(sqlmock.NewResult(1, 1))
	// 这里返回的结果不再是正确的，而是模拟db返回错误的信息
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
	// 如mock.ExpectBegin和mock.ExpectCommit；而这里使用的不是事务，而是普通的查询语句，
	// 所以这里直接使用了mock.ExpectQuery，如果前后不对应是不能mock数据的
	productRows := sqlmock.NewRows([]string{"id", "name"}).AddRow("22", "headset")
	mock.ExpectQuery("SELECT (.+) FROM products WHERE id = \\?").
		WithArgs("22").
		WillReturnRows(productRows)

	expectdProduct := &Product{
		id:   "22",
		name: "headset",
	}
	product, _ := GetSingleProductDao(db, "22")

	assert.Equal(t, expectdProduct, product)
}
