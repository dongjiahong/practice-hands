docker run \
    -p 10022:22 \
    -p 10800:3000 \
    --name=gogs-test \
    -e TZ="Asia/Shanghai" \
    -v `pwd`/data:/data  \
    -d gogs/gogs
