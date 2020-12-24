package conf

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type BaseConfig struct {
	ConfigPath        string `toml:"config_path"`
	AppPort           int64  `toml:"app_port"`
	TaskRedisHost     string `toml:"task_redis_host"`
	TaskRedisPort     int    `toml:"task_redis_port"`
	TaskRedisPassword string `toml:"task_redis_password"`
	TaskRedisDB       int    `toml:"task_redis_db"`
	TaskRedisPoolSize int    `toml:"task_redis_pool_size"`
}

// Cfg 全局配置
var Cfg = &BaseConfig{}

// InitConfig 初始化配置
func InitConfig() {
	path, pErr := os.Getwd()
	if pErr != nil {
		log.Panic(pErr)
	}
	path += "/config/settings.toml"
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		log.Panic(err)
	}
}
