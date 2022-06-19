package config

import (
	"github.com/jishulangcom/go-fun"
	"github.com/spf13/viper"
	"path/filepath"
	"time"
)

// new一个配置
func NewConfig(conf interface{}, env string) {
	// 配置目录
	dirName := "config"
	destPath := "./config"
	is, _ := fun.DirIsExist(destPath)
	if !is {
		srcPath := fun.GetImportPackagePath("github.com/jishulangcom/go-config", "v0.0.7")
		srcPath = filepath.Join(srcPath, dirName)
		err := fun.DirCopy(srcPath, destPath)
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 1)
	}

	// 默认配置文件名
	if env == "" {
		env = "local" // 默认本地
	}

	//
	viperObj := viper.New()
	viperObj.AddConfigPath(dirName) // 配置目录
	viperObj.SetConfigName(env) // 配置文件名
	viperObj.SetConfigType("json") // 配置文件名后缀
	if err := viperObj.ReadInConfig(); err != nil {
		panic(err)
	}

	//var config ConfigDto
	if err := viperObj.Unmarshal(&conf); err != nil {
		panic(err)
	}
}