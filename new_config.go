package config

import (
	"errors"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// new一个配置
func NewConfig(config interface{}, env string) {
	// 配置目录
	dirName := "config"
	destPath := "./config2"
	is, _ := DirIsExist(destPath)
	if !is {
		srcPath := GetPackagePath("github.com/jishulangcom/go-config", "v0.0.1")
		srcPath = filepath.Join(srcPath, dirName)
		err := DirCopy(srcPath, destPath)
		if err != nil {
			panic(err)
		}
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
	if err := viperObj.Unmarshal(&config); err != nil {
		log.Println("配置转换结构异常")
		panic(err)
	}
}


func GetPackagePath(packagePath string, version string) string {
	path := filepath.Join(os.Getenv("GOPATH"), "pkg", "mod")
	arr := Explode("/", packagePath)

	for _, v := range arr {
		path = filepath.Join(path, v)
	}
	path += "@"+ version

	return path
}

func Explode(separator, str string) []string {
	return strings.Split(str, separator)
}


/*
	【名称:】目录复制
	【参数:】拷贝路径(string)，目标路径
	【返回:】布尔
	【备注:】
*/
func DirCopy(srcPath string, destPath string) error {
	var err error
	_, err = DirIsExistAndChk(srcPath)
	if err != nil {
		return err
	}

	is, _ := DirIsExistAndChk(destPath)
	if !is { // 不存在，就创建目录
		err = DirCreate(destPath)
		if err != nil {
			return err // 目录创建失败
		}
	}

	//
	err = filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if !f.IsDir() {
			path := strings.Replace(path, "\\", "/", -1)
			destNewPath := DirEndsWithSlash(destPath)
			destNewPath += f.Name()
			err = FileCopy(path, destNewPath)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}


/*
	【名称:】目录路径以斜杠结尾
	【参数:】路径(string)
	【返回:】路径(string)
	【备注:】
*/
func DirEndsWithSlash(dirPath string) string {
	dirPath = strings.Replace(dirPath, "\\", "/", -1)
	if !strings.HasSuffix(dirPath, "/") {
		dirPath += "/"
	}

	return dirPath
}


/*
	【名称:】文件复制
	【参数:】文件路径(string)
	【返回:】是否成功(bool)，错误信息(error)
	【备注:】
*/
func FileCopy(srcPath, destPath string) (error) {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	//
	destSplitPathDirs := strings.Split(destPath, "/") //分割path目录
	destSplitPath := "" //检测时候存在目录
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs)-1 {
			destSplitPath = destSplitPath + dir + "/"
			is, _ := DirIsExist(destSplitPath)
			if is == false {
				err := os.Mkdir(destSplitPath, os.ModePerm)
				if err != nil {
					return err
				}
			}
		}
	}

	//
	dstFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

/*
	【名称:】目录是否存在，若存在检验是否是目录
	【参数:】文件路径(string)
	【返回:】布尔
	【备注:】
*/
func DirIsExistAndChk(dirPath string) (bool, error) {
	info, err := os.Stat(dirPath)
	if err != nil {
		return false, err
	}

	if !info.IsDir() {
		return false, errors.New(dirPath + "已存在，但不是目录")
	}

	return true, nil
}

/*
	【名称:】目录是否存在
	【参数:】文件路径(string)
	【返回:】布尔
	【备注:】
*/
func DirIsExist(dirPath string) (bool, error) {
	return FileIsExist(dirPath)
}

/*
	【名称:】文件是否存在
	【参数:】文件路径(string)
	【返回:】布尔
	【备注:】
*/
func FileIsExist(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

/*
	【名称:】创建目录
	【参数:】目录路径(string)，权限(os.FileMode)
	【返回:】error
	【备注:】
*/
func DirCreate(dirPath string) (error) {
	return os.MkdirAll(dirPath, 0766)
}