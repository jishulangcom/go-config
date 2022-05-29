package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

// new一个配置
func NewConfig(config interface{}, env string) {
	dir := "config"
	if env == "" {
		env = "dev"
	}

	fmt.Println(os.Getwd())

	is := FileIsExist(dir)
	fmt.Println(is)
	return

	viperObj := viper.New()
	viperObj.AddConfigPath(dir)
	viperObj.SetConfigName(env)
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


// 文件是否存在
func FileIsExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}