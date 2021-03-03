# 代码结构

[golang项目结构标准说明](https://github.com/golang-standards/project-layout/blob/master/README_zh-CN.md)

```sh
├── LICENSE.md     # 证书
├── Makefile       # 执行scripts里的脚本
├── README.md
├── api
├── assets         # 项目中使用的其他资源
├── build
├── cmd            # 项目主要的应用程序。对于每个程序来说，这个目录的名字应该和可执行文件名字匹配
├── configs        # 默认文件模板或默认配置。将confd或者consul-template文件放在这里
├── deployments
├── docs
├── examples
├── githooks
├── go.mod
├── init          # 系统初始化（systemd、upstart、sysv）和进程管理(runit、supervisord)配置
├── internal      # 私有的应用程序代码库。这些是不希望被其他人导入的代码
├── pkg           # 外部应用程序可以使用的库代码。其他项目将会导入这些库来保证项目可以正常运行，项目的私有代码一般放在/internal目录
├── scripts
├── test          # 外部测试应用程序和测试数据。随时更加需要构建/test目录
├── third_party   # 外部辅助工具，fork的代码和其他第三方工具如(Swagger UI)
├── tools         # 此项目的工具。请注意这些工具可以从/pkg和/internal目录导入代码
├── vendor
├── web           # web应用程序特定的组件：静态web资源、服务器端模板和但也应用(Single-Page App、SPA)
└── website       # 如果不适用github page, 则在这里放置项目的网站数据
```
