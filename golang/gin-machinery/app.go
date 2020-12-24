package main

import (
	"fmt"
	"os"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli" // 用来构建命令行应用

	"gin-machinery/config"
	"gin-machinery/gredis"
	handlerouter "gin-machinery/handle"
	tasks "gin-machinery/tasks"
)

var (
	server   *machinery.Server
	cnf      *config.Config
	app      *cli.App
	tasksMap map[string]interface{}
)

func init() {
	// 初始化配置
	conf.InitConfig()

	// 命令行标签主要是命令中用不同的标签运行worker或者sender
	var err error
	app = cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "c",
			Value: "",
			Usage: "Path to a configuration file.",
		},
	}

	tasksMap = map[string]interface{}{
		"add":               tasks.Add,
		"long_running_task": tasks.LongRunningTask,
	}
	if cnf, err = loadConfig(conf.Cfg.ConfigPath); err != nil {
		panic(err)
	}

	if server, err = machinery.NewServer(cnf); err != nil {
		panic(err)
	}
	gredis.InitRedisClient()
}

func loadConfig(configPath string) (*config.Config, error) {
	if configPath != "" {
		return config.NewFromYaml(configPath, true)
	}
	return nil, fmt.Errorf("no find config file")
}

func runWorker() (err error) {
	server.RegisterTasks(tasksMap)
	if err != nil {
		panic(err)
	}
	workers := server.NewWorker("worker_test", 10)
	err = workers.Launch()
	if err != nil {
		panic(err)
	}
	return
}

func runSender() (err error) {
	r := gin.Default()

	// ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/add", func(c *gin.Context) {
		handlerouter.Add(c, server)
	})
	r.POST("/longRunningTask", func(c *gin.Context) {
		handlerouter.LongRunningTask(c, server)
	})
	err = r.Run(fmt.Sprintf(":%d", conf.Cfg.AppPort))
	return
}

// meachinery 实例初始化
func startServer() (err error) {
	// Create server instance
	server, err = machinery.NewServer(cnf)
	if err != nil {
		return
	}

	// 注册任务
	err = server.RegisterTasks(tasksMap)
	return
}

func main() {
	// 运行worker: go run app.go worker
	// 运行sender：go run app.go send
	app.Commands = []cli.Command{
		{
			Name:  "worker",
			Usage: "launch machinery worker",
			Action: func(c *cli.Context) error {
				if err := runWorker(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
		{
			Name:  "send",
			Usage: "send async tasks",
			Action: func(c *cli.Context) error {
				if err := runSender(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
	}

	// run the cli app
	app.Run(os.Args)
}
