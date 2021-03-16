docker run  --name mysql-test -d -p3306:3306 --restart always \
 -v `pwd`/mysql:/etc/mysql \
 -v `pwd`/mysql-files:/var/lib/mysql-files \
 -v `pwd`/logs/:/logs \
 -v `pwd`/data/:/var/lib/mysql \
 -e MYSQL_ALLOW_EMPTY_PASSWORD=true mysql:latest
