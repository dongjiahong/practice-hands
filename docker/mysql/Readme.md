1. Mysql外部数据和配置文件路径
msyql配置文件路径：/etc/mysql
mysql数据卷路径：/var/lib/mysql
1.1 拉去mysql:5.7镜像
docker pull mysql:5.7

2. 创建mysql:5.7容器
宿主机mysql配置文件路径：/root/docker/mysqletc/mysql
宿主机mysql数据卷路径：/root/docker/mysqletc/data
注：路径可以自己定义

1. 创建一个临时的msyql:5.7容器，目的是cp 容器里mysql的配置文件
docker run -d -p 3306:3306 --name myMysql -e MYSQL_ROOT_PASSWORD=root mysql:5.7

2. 复制容器中的mysql配置文件到宿主机指定目录    
从容器中将文件拷贝出来的命令：dokcer cp 容器名称:容器目录 需要拷贝的文件或目录
docker cp myMysql:/etc/mysql /root/docker/mysqletc/mysql

3. 停止并删除刚才创建的临时容器
docker stop myMysql
dokcer rm myMysql

4. 创建并启动mysql:5.7容器
dokcer run -d --name mysql5.7 -p 33306:3306 --restart always --privileged=true -v /root/docker/mysqletc/mysql:/etc/mysql -v /root/docker/mysqletc/data:/var/lib/mysql -e MYSQL_USER="summit" -e MYSQL_PASSWORD="summit" -e MYSQL_ROOT_PASSWORD="root" mysql:5.7

5. 大功告成，查看容器日志
docker logs mysql5.7
2.1 涉及到的命令行参数
--restart always                                -> 开机启动
--privileged=true                               -> 提升容器内权限
-v /root/docker/mysqletc/mysql:/etc/mysql       -> 映射配置文件
-v /root/docker/mysqletc/data:/var/lib/mysql    -> 映射数据目录
-e MYSQL_USER="summit"                          -> 添加用户summit
-e MYSQL_PASSWORD="summit"                      -> 设置summit用户的密码为summit
-e MYSQL_ROOT_PASSWORD="root"                   -> 设置root的密码为root

