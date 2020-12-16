docker run -d -p 3306:3306 --name mysql-test \
 -v `pwd`/conf/:/etc/mysql \
 -v `pwd`/mysql-files:/var/lib/mysql-files \
 -v `pwd`/logs/:/logs \
 -v `pwd`/data/:/var/lib/mysql \
 -e MYSQL_ALLOW_EMPTY_PASSWORD=true mysql:latest
