package config

type HttpCnfDto struct {
	ListenHTTP string `yaml:"listen_http" toml:"listen_http" json:"listen_http" bson:"listen_http"`
	Profile    bool   `yaml:"profile" toml:"profile" json:"profile" bson:"profile"`
	Prometheus bool   `yaml:"prometheus" toml:"prometheus" json:"prometheus" bson:"prometheus"`
	Tracing    bool   `yaml:"tracing" toml:"tracing" json:"tracing" bson:"tracing"`
	Verbose    bool   `yaml:"verbose" toml:"verbose" json:"verbose" bson:"verbose"`
}