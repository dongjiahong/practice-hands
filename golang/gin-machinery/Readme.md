# gin-machinery

## 目录结构
.
├── app.go              #主程序入口
├── config              #app配置文件
│   ├── conf.go
│   └── settings.toml
├── config.yml          #machinery配置文件
├── gredis              #redis存储任务结果进度
│   └── gredis.go
├── handle              #路由处理
│   ├── handle_router.go
├── tasks               #任务处理
│   └── tasks.go
└── util 
    └── tool.go
