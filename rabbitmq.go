package config

type RabbitMQCnfDto struct {
	Host  string `yaml:"host" toml:"host" json:"host" bson:"host"`     // 服务器地址
	Port  int    `yaml:"port" toml:"port" json:"port" bson:"port"`     // 端口
	User  string `yaml:"user" toml:"user" json:"user" bson:"user"`     // 用户名
	Pwd   string `yaml:"pwd" toml:"pwd" json:"pwd" bson:"pwd"`         // 密码
	Vhost   string `yaml:"vhost" toml:"vhost" json:"vhost" bson:"vhost"`         // 密码
	Debug bool   `yaml:"debug" toml:"debug" json:"debug" bson:"debug"` // 调试
}

var RabbitMQCnf = RabbitMQCnfDto{
	Host:  "127.0.0.1",
	Port:  5672,
	User: "",
	Pwd:   "",
	Vhost: "",
	Debug: true,
}
