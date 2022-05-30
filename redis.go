package config

type RedisCnfDto struct {
	Host string `yaml:"host" toml:"host" json:"host" bson:"host"` // 地址
	Port string `yaml:"port" toml:"port" json:"port" bson:"port"` // 端口
	Pwd  string `yaml:"pwd" toml:"pwd" json:"pwd" bson:"pwd"`     // 密码
	DB   int    `yaml:"db" toml:"db" json:"db" bson:"db"`         // 库
}