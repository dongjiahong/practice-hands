//+build wireinject
// 上面注释告诉编译器不要编译这个文件

package main

import (
	"github.com/google/wire"
	"learnwire/internal/config"
	"learnwire/internal/db"
)

func InitApp() (*App, error) {
	panic(wire.Build(config.Provider, db.Provider, NewApp)) // 调用wire.Build方法传入所有的依赖对象以及构建最终对象的函数得到目标对象
}
