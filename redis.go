package config

type RedisCnfDto struct {
	Host         string `yaml:"host" toml:"host" json:"host" bson:"host"` // 地址
	Port         string `yaml:"port" toml:"port" json:"port" bson:"port"` // 端口
	User         string `yaml:"user" toml:"user" json:"user" bson:"user"` // 用户名
	Pwd          string `yaml:"pwd" toml:"pwd" json:"pwd" bson:"pwd"`     // 密码
	DB           string `yaml:"db" toml:"db" json:"db" bson:"db"`         // 库
	PoolSize int    `yaml:"pool_size" toml:"pool_size" json:"pool_size" bson:"pool_size"`
}