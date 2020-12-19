package main

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言DB.Get()方法释放被调用

	m := NewMockDB(ctrl)
	// 这里给接口里的Get函数mock数据，即：当传入“Tom”时，返回100和一个error
	// Eq(value) 表示与value等价的值
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))
	// Any()可以用来表示任意参数
	// m.EXPECT().Get(gomock.Any()).Return(630, nil)
	// Not(value) 用来表示非value以外的值
	// m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
	// Nil() 表示None值
	// m.EXPECT().Get(gomock.Nil()).Return(0, errors.New("nil"))

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got: ", v)
	}
}
