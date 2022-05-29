package config

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type ConfigDto struct {
		MysqlCnf MysqlCnfDto
		RedisCnf RedisCnfDto
	}

	var config ConfigDto
	NewConfig(&config, "")

	fmt.Println("cinfig:", config)
}
