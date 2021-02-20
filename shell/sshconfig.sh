# ssh 的配置文件位于~/.ssh/config
# 保活，每隔20秒
ServerAliveInterval 20
# 只发10次,如果10次都失败也断开
ServerAliveCountMax 10

Host test
    Hostname 118.31.185.134
    Port 22
    User root
Host online
    Hostname 112.125.27.11
    Port 22
    User root

