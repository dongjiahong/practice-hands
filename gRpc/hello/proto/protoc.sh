
# --go_out 其实是让protoc去调用protoc-gen-go工具
# Protobuf的protoc编译器是通过插件机制实现对不同语言的支持。比如protoc命令出现--xxx_out格式的参数，
# 那么protoc将首先查询是否有内置的xxx插件，如果没有内置的xxx插件那么将继续查询当前系统中是否存在
# protoc-gen-xxx命名的可执行程序，最终通过查询到的插件生成代码。对于Go语言的protoc-gen-go插件来说，
# 里面又实现了一层静态插件系统。比如protoc-gen-go内置了一个gRPC插件，
# 用户可以通过--go_out=plugins=grpc参数来生成gRPC相关代码，否则只会针对message生成相关代码。
protoc -I . --go_out=plugins=grpc:. hello.proto
