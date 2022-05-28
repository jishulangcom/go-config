package config

import (
	"github.com/spf13/viper"
	"log"
)


type ConfigDto struct {
	MysqlCnf MysqlConfig
	RedisCnf RedisConfig
}


var Config ConfigDto


func ConfigInit(config interface{}) {
	viperObj := viper.New()
	viperObj.AddConfigPath("config/")
	viperObj.SetConfigName("settings")
	viperObj.SetConfigType("json")
	if err := viperObj.ReadInConfig(); err != nil {
		panic(err)
	}

	//var config ConfigDto
	if err := viperObj.Unmarshal(&config); err != nil {
		log.Println("配置转换结构异常")
		panic(err)
	}
}
