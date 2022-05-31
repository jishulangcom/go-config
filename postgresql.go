package config

type PostgreSqlCnfDto struct {
	Host    string `yaml:"host" toml:"host" json:"host" bson:"host"`                 // 服务器地址
	Port    int    `yaml:"port" toml:"port" json:"port" bson:"port"`                 // 端口
	User    string `yaml:"user" toml:"user" json:"user" bson:"user"`                 // 用户名
	Pwd     string `yaml:"pwd" toml:"pwd" json:"pwd" bson:"pwd"`                     // 密码
	DbName  string `yaml:"dbname" toml:"dbname" json:"dbname" bson:"dbname"`         // 库名
	Debug   bool   `yaml:"debug" toml:"debug" json:"debug" bson:"debug"`             // 调试
	MaxConn int    `yaml:"max_conn" toml:"max_conn" json:"max_conn" bson:"max_conn"` // 最大连接
}

var PostgreSqlCnf = PostgreSqlCnfDto{
	Host:    "127.0.0.1",
	Port:    5432,
	User:    "postgres",
	Pwd:     "",
	DbName:  "postgres",
	Debug:   true,
	MaxConn: 10,
}
