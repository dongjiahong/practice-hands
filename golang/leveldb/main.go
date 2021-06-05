package main

import (
	"fmt"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"

	"level/ldb"
)

func Raw() {
	// 创建并打开数据库
	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 写入数据
	db.Put([]byte("key1"), []byte("value1"), nil)
	db.Put([]byte("key2"), []byte("value2"), nil)
	db.Put([]byte("key3"), []byte("value3"), nil)
	db.Put([]byte("key4"), []byte("value4"), nil)
	db.Put([]byte("key5"), []byte("value5"), nil)

	// 循环遍历
	/* -- 结果 --
	循环遍历
	key: key1, value: value1
	key: key2, value: value2
	key: key3, value: value3
	key: key4, value: value4
	key: key5, value: value5
	*/
	fmt.Println("循环遍历")
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Printf("key: %s, value: %s\n", iter.Key(), iter.Value())
	}
	iter.Release()

	// 读取某条数据
	/* -- 结果 --
	读取单挑数据key2: value2
	*/
	data, err := db.Get([]byte("key2"), nil)
	if err != nil {
		fmt.Println("value err: ", err, " key: key2")
	}
	fmt.Printf("读取单挑数据key2: %s\n", data)

	// 批量写入数据
	batch := new(leveldb.Batch)
	batch.Put([]byte("key6"), []byte(strconv.Itoa(10000)))
	batch.Put([]byte("key7"), []byte(strconv.Itoa(20000)))
	batch.Delete([]byte("key4"))
	db.Write(batch, nil)

	// 查找数据
	/* -- 结果 --
	查找数据: key1, value: value1
	查找数据: key2, value: value2
	查找数据: key3, value: value3
	查找数据: key5, value: value5
	查找数据: key6, value: 10000
	查找数据: key7, value: 20000
	*/
	key := "key"
	iter = db.NewIterator(nil, nil)
	for ok := iter.Seek([]byte(key)); ok; ok = iter.Next() {
		fmt.Printf("查找数据: %s, value: %s\n", iter.Key(), iter.Value())
	}
	iter.Release()

	// 按key范围遍历数据
	/* -- 结果 --
	按key范围遍历数据
	key: key3, value: value3
	key: key5, value: value5
	key: key6, value: 10000
	*/
	fmt.Println("按key范围遍历数据")
	iter = db.NewIterator(&util.Range{Start: []byte("key3"), Limit: []byte("key7")}, nil)
	for iter.Next() {
		fmt.Printf("key: %s, value: %s\n", iter.Key(), iter.Value())
	}
	iter.Release()

	// 使用特定前缀迭代数据库内容的子集
	/* -- 结果 --
	查询以key为前缀的value
	key: key1, value: value1
	key: key2, value: value2
	key: key3, value: value3
	key: key5, value: value5
	key: key6, value: 10000
	key: key7, value: 20000
	*/
	fmt.Println("查询以key为前缀的value")
	iter = db.NewIterator(util.BytesPrefix([]byte("key")), nil)
	for iter.Next() {
		fmt.Printf("key: %s, value: %s\n", string(iter.Key()), string(iter.Value()))
	}
	iter.Release()

}

func main() {
	// Raw()
	var conf = ldb.Config{DBPath: "./db"}
	db, err := ldb.NewDB(&conf)
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()

	db.Put("k1", "v1")
	db.Put("k2", "v3")
	if data, err := db.Get("k1"); err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", data)
	}

}
