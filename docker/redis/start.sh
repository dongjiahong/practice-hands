# 拉取镜像
# docker pull redis:alpine
docker run  -d -p 6379:6379 --name redis-test --restart always redis:alpine 
# 登录redis
# docker exec -it 12uu89u1i23 redis-cli --raw
