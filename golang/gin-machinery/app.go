package main

import (
	"fmt"
	"os"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"

	config "gin-machinery/config"
	"gin-machinery/gredis"
	handlerouter "gin-machinery/handle"
	tasks "gin-machinery/tasks"
)

var (
	server *machinery.Server
	cnf    *config.Config
	app    *cli.App
	tasks  map[string]interface{}
)

func init() {
	// 初始化配置
	config.InitConfig()

	// 命令行标签主要是命令中用不同的标签运行worker或者sender
	var err error
	app = cli.NewApp()
}
