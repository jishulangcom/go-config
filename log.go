package config

type LogCnfDto struct {
	Debug    bool   `yaml:"debug" toml:"debug" json:"debug" bson:"debug"`
	Output   string `yaml:"output" toml:"output" json:"output" bson:"output"`
	Level    string `yaml:"level" toml:"level" json:"level" bson:"level"`
	Encoding string `yaml:"encoding" toml:"encoding" json:"encoding" bson:"encoding"`
}
