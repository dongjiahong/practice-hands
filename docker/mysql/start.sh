docker run -d -p 3306:3306 --name mysql-test \
 -v /home/lele/develop/docker/mysql/conf/:/etc/mysql \
 -v /home/lele/develop/docker/mysql/mysql-files:/var/lib/mysql-files \
 -v /home/lele/develop/docker/mysql/logs/:/logs \
 -v /home/lele/develop/docker/mysql/data/:/var/lib/mysql \
 -e MYSQL_ALLOW_EMPTY_PASSWORD=true mysql
