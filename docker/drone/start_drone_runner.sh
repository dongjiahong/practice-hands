docker run -d \
  -v `pwd`/run/docker.sock:/var/run/docker.sock \
  -e DRONE_RPC_PROTO=http \
  -e DRONE_RPC_HOST=127.0.0.1:3080 \
  -e DRONE_RPC_SECRET=5c5431e8b005e9ebe6fa5f9f517c9988 \
  -e DRONE_RUNNER_CAPACITY=2 \
  -e DRONE_RUNNER_NAME=runner-docker \
  -e TZ="Asia/Shanghai" \
  -p 3000:3000 \
  --restart always \
  --name drone-runner \
  drone/drone-runner-docker
