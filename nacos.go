package config

type NacosConfig struct {
	NameSpace string `yaml:"name_space" toml:"name_space" json:"name_space" bson:"name_space"`
	Host      string `yaml:"host" toml:"host" json:"host" bson:"host"`
	Port      uint64 `yaml:"port" toml:"port" json:"port" bson:"port"`
	DataId    string `yaml:"data_id" toml:"data_id" json:"data_id" bson:"data_id"`
	Group     string `yaml:"group" toml:"group" json:"group" bson:"group"`
}