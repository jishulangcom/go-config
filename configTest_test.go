package config

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type ConfigDto struct {
		MysqlCnf MysqlConfig
		RedisCnf RedisConfig
	}

	var config ConfigDto
	ConfigInit(&config)

	fmt.Println("cinfig:", config)
}
