# 这里选择alpine而不是latest的原因
# 1.alpine用的镜像跟latest的一样
# 2.alpine的镜像用的是Alpine Linux内核，比ubuntu小的多
FROM golang:alpine
# 配置go的环境
ENV GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

# 设置linux上的工作目录为/home/www/hello
WORKDIR /home/www/hello
# 将代码复制到容器中
COPY ./ /home/www/hello
# go构建可执行文件
RUN go build -o hello .
# 执行命令
ENTRYPOINT ["./hello"]
