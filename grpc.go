package config

type GrpcConfig struct {
	Addr string `yaml:"addr" toml:"addr" json:"addr" bson:"addr"`
}