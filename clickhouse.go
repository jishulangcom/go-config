package config

type ClickHouseCnfDto struct {
	Host   string `yaml:"host" toml:"host" json:"host" bson:"host"`         // 服务器地址
	Port   int    `yaml:"port" toml:"port" json:"port" bson:"port"`         // 端口
	User   string `yaml:"user" toml:"user" json:"user" bson:"user"`         // 用户名
	Pwd    string `yaml:"pwd" toml:"pwd" json:"pwd" bson:"pwd"`             // 密码
	DbName string `yaml:"dbname" toml:"dbname" json:"dbname" bson:"dbname"` // 库名
	Debug  bool   `yaml:"debug" toml:"debug" json:"debug" bson:"debug"`     // 调试
}

var ClickHouseCnf = ClickHouseCnfDto{
	Host:   "127.0.0.1",
	Port:   9000,
	Pwd:    "",
	DbName: "",
	Debug:  true,
}