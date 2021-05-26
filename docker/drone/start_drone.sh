docker run \
  -v `pwd`/data:/data \
  --env=DRONE_GITLAB_SERVER=127.0.0.1:10080 \
  --env=DRONE_GITLAB_CLIENT_ID=a7686bbb2410ae24200c2febe2bba6c722c8ae4deb54bddb97cf6b44d9716177 \
  --env=DRONE_GITLAB_CLIENT_SECRET=4d9543501973898ea6e7bb299e4e3b85c85410dca244dc3429761b96d12ad991 \
  --env=DRONE_RPC_SECRET=5c5431e8b005e9ebe6fa5f9f517c9988 \
  --env=DRONE_SERVER_HOST=127.0.0.1:3080 \
  --env=DRONE_SERVER_PROTO=http \
  --env=DRONE_LOGS_TEXT=true \
  --env=DRONE_LOGS_PRETTY=true \
  --env=DRONE_LOGS_COLOR=true \
  --env=DRONE_LOGS_DEBUG=true \
  --publish=3080:80 \
  --restart=always \
  --detach=true \
  --name=drone-test \
  drone/drone

# 查看server的日志 
# docker logs -f drone-test
