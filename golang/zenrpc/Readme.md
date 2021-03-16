# zenrpc框架
##  启动服务端
```sh
cd server
go run main.go
```
## 调用客户端
```sh
cd client
go run client.go
```

## zenrpc的说明
zenrpc起的是http服务，所以client也是通过http来请求，不同的是使用的是json2.0,
所以在调用是传参要参考json2.0的模式

在服务端代码中要记得加`//go:genreate zenrpc`，这样我们可以使用`zenrpc`指令来生成对应代码

### 调用示例
```sh
# 请求
{"jsonrpc": "2.0", "method": "subtract", "params": [42, 23], "id": 1}
# 答复
{"jsonrpc": "2.0", "result": 19, "id": 1}

# 错误示例
# 请求
{"jsonrpc": "2.0", "method": 1, "params": "bar"}
# 答复
{"jsonrpc": "2.0", "error": {"code": -32600, "message": "Invalid Request"}, "id": null}
```
