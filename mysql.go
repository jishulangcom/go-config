package config

type MysqlConfig struct {
	Host         string `yaml:"host" toml:"host" json:"host" bson:"host"` //
	Port         string `yaml:"port" toml:"port" json:"port" bson:"port"` // 端口
	User         string `yaml:"user" toml:"user" json:"user" bson:"user"` // 用户名
	Pwd          string `yaml:"pwd" toml:"pwd" json:"pwd" bson:"pwd"`     // 密码
	DB           uint8  `yaml:"db" toml:"db" json:"db" bson:"db"`         // 库
	Charset      string `yaml:"charset" toml:"charset" json:"charset" bson:"charset"`
	MaxIdleCount int    `yaml:"max_idle_count" toml:"max_idle_count" json:"max_idle_count" bson:"max_idle_count"`
	MaxOpenCount int    `yaml:"max_open_count" toml:"max_open_count" json:"max_open_count" bson:"max_open_count"`
	Options      string `yaml:"options" toml:"options" json:"options" bson:"options"`
	TimeoutSec   int    `yaml:"timeout_sec" toml:"timeout_sec" json:"timeout_sec" bson:"timeout_sec"`
	Tracing      bool   `yaml:"tracing" toml:"tracing" json:"tracing" bson:"tracing"`
}