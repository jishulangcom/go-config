package test

import (
	"fmt"
	"github.com/jishulangcom/go-config"
	// "github.com/joho/godotenv"
	"testing"
)

func Test(t *testing.T) {
	/*
	// env
	err := godotenv.Load()
	if err != nil {
		logger.Panic("加载.env文件失败")
	}
	env := os.Getenv("ENV")
	if env == "" {
		logger.Panic(".env未配置ENV")
	}
	*/

	type ConfigDto struct {
		MysqlCnf config.MysqlCnfDto
		RedisCnf config.RedisCnfDto
	}
	var conf ConfigDto
	config.NewConfig(&conf, "") // 会自动在当前目录生成config目录，里面有json配置文件

	fmt.Println("conf:", conf)
}
