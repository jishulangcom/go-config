package config

type KafkaCnfDto struct {
	Host  string `yaml:"host" toml:"host" json:"host" bson:"host"`     // 服务器地址
	Port  int    `yaml:"port" toml:"port" json:"port" bson:"port"`     // 端口
	Debug bool   `yaml:"debug" toml:"debug" json:"debug" bson:"debug"` // 调试
}

var KafkaCnf = KafkaCnfDto{
	Host:  "127.0.0.1",
	Port:  9092,
	Debug: true,
}
