package config

type GrpcCnfDto struct {
	Host         string `yaml:"host" toml:"host" json:"host" bson:"host"` // 地址
	Port         string `yaml:"port" toml:"port" json:"port" bson:"port"` // 端口
}