package config

type RedisConfig struct {
	Addr     string `yaml:"addr" toml:"addr" json:"addr" bson:"addr"`
	Passwd   string `yaml:"passwd" toml:"passwd" json:"passwd" bson:"passwd"`
	PoolSize int    `yaml:"pool_size" toml:"pool_size" json:"pool_size" bson:"pool_size"`
	DB       int    `yaml:"db" toml:"db" json:"db" bson:"db"`
}