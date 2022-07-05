package config

type TarantoolCnfDto struct {
	Host  string `yaml:"host" toml:"host" json:"host" bson:"host"`     // 服务器地址
	Port  int    `yaml:"port" toml:"port" json:"port" bson:"port"`     // 端口
	Pwd   string `yaml:"pwd" toml:"pwd" json:"pwd" bson:"pwd"`         // 密码
	User  string `yaml:"user" toml:"user" json:"user" bson:"user"`     // 用户名
	Debug bool   `yaml:"debug" toml:"debug" json:"debug" bson:"debug"` // 调试
}

var TarantoolCnf = TarantoolCnfDto{
	Host:  "127.0.0.1",
	Port:  3301,
	Pwd:   "",
	User:  "",
	Debug: true,
}