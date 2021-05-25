 docker run --name openresty-test -d -p2000:80 --restart always \
    -v `pwd`/conf.d/:/etc/nginx/conf.d/ \
    -v `pwd`/webroot/shared1/:/tmp/webroot/shared1/ \
    -v `pwd`/webroot/myshare/:/tmp/webroot/myshare/ \
    -v `pwd`/logs/:/var/log/nginx/logs \
    openresty/openresty:centos
