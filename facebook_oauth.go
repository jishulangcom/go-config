package config

type FacebookOAuthCnfDto struct {
	ClientID     string `yaml:"client_id" toml:"client_id" json:"client_id" bson:"client_id"`                 // 客户端ID
	ClientSecret string `yaml:"client_secret" toml:"client_secret" json:"client_secret" bson:"client_secret"` // 客户端密钥
	RedirectURL  string `yaml:"redirect_url" toml:"redirect_url" json:"redirect_url" bson:"redirect_url"`     // 回调地址
	Debug        bool   `yaml:"debug" toml:"debug" json:"debug" bson:"debug"`                                 // 调试
}
