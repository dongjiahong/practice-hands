// https://geektutu.com/post/quick-gomock.html#1-gomock-%E7%AE%80%E4%BB%8B
package main

// 用来mock DB的接口
type DB interface {
	Get(key string) (int, error)
}

// GetFromDB是要测试的函数
func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}
	return -1
}
