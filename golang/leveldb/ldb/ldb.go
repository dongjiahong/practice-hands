// Package ldb leveldb包的封装
package ldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type Config struct {
	DBPath string // db路径
}

type LDB struct {
	DB   *leveldb.DB
	conf *Config
}

func NewDB(conf *Config) (*LDB, error) {
	db, err := leveldb.OpenFile(conf.DBPath, nil)
	if err != nil {
		return nil, err
	}
	return &LDB{
		DB:   db,
		conf: conf,
	}, nil
}

func Close(db *leveldb.DB) {
	if db == nil {
		return
	}
	db.Close()
}

// Put 记录状态
func (d *LDB) Put(key, value string) error {
	return d.DB.Put([]byte(key), []byte(value), nil)
}

// Get 获取指定key的status
func (d *LDB) Get(key string) (string, error) {
	res, err := d.DB.Get([]byte(key), nil)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func (d *LDB) Delete(key string) error {
	return d.DB.Delete([]byte(key), nil)
}

// Search 搜索key
func (d *LDB) Search(key string) []string {
	iter := d.DB.NewIterator(nil, nil)
	defer iter.Release()
	res := make([]string, 0, 1)
	for ok := iter.Seek([]byte(key)); ok; ok = iter.Next() {
		res = append(res, string(iter.Value()))
	}
	return res
}

// SearchRange 查询指定范围的结果集
func (d *LDB) SearchRange(key, limit string) []string {
	res := make([]string, 0, 1)
	iter := d.DB.NewIterator(&util.Range{Start: []byte(key), Limit: []byte(limit)}, nil)
	defer iter.Release()
	for iter.Next() {
		res = append(res, string(iter.Value()))
	}
	return res
}
