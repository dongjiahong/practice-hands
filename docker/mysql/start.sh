docker run  --name mysql-test -d -p3306:3306 --restart always \
 -v `pwd`/mysql/my.cnf:/etc/mysql/my.cnf \
 -v `pwd`/mysql/conf.d:/etc/mysql/conf.d \
 -v `pwd`/data:/var/lib/mysql \
 -v MYSQL_ALLOW_EMPTY_PASSWORD=true mysql:latest
