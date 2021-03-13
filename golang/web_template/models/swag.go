package models

type ExampleSwag struct {
	UserName   string `json:"user_name" binding:"required"`                             // 用户名字
	Phone      string `json:"phone" binding:"required,len=11"`                          // 用户电话
	Gender     string `json:"gender" binding:"required,oneof=男 女"`                      // 性别 男、女
	EntryTime  string `json:"entry_time" binding:"required,datetime=2006.01.02"`        // 入职日期
	UserStatus string `json:"user_status" binding:"required,oneof=小组组长 正式员工 试用员工 试岗员工"` // 用户的状态: 正式、试用、试岗、组长
}

type HelloSwag struct {
	Say string `json:"say" binding:"required"`
}
