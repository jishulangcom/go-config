package config

type RedisPoolCnfDto struct {
	PoolSize           int `yaml:"pool_size" toml:"pool_size" json:"pool_size" bson:"pool_size"`                                             //
	DialTimeout        int `yaml:"dial_timeout" toml:"dial_timeout" json:"dial_timeout" bson:"dial_timeout"`                                 //
	ReadTimeout        int `yaml:"read_timeout" toml:"read_timeout" json:"read_timeout" bson:"read_timeout"`                                 //
	WriteTimeout       int `yaml:"write_timeout" toml:"write_timeout" json:"write_timeout" bson:"write_timeout"`                             //
	PoolTimeout        int `yaml:"pool_timeout" toml:"pool_timeout" json:"pool_timeout" bson:"pool_timeout"`                                 //
	IdleCheckFrequency int `yaml:"idle_check_frequency" toml:"idle_check_frequency" json:"idle_check_frequency" bson:"idle_check_frequency"` //
	IdleTimeout        int `yaml:"idle_timeout" toml:"idle_timeout" json:"idle_timeout" bson:"idle_timeout"`                                 //
	MaxConnAge         int `yaml:"max_conn_age" toml:"max_conn_age" json:"max_conn_age" bson:"max_conn_age"`                                 //
	MaxRetries         int `yaml:"max_retries" toml:"max_retries" json:"max_retries" bson:"max_retries"`                                     //
	MinRetryBackoff    int `yaml:"min_retry_backoff" toml:"min_retry_backoff" json:"min_retry_backoff" bson:"min_retry_backoff"`             //
	MaxRetryBackoff    int `yaml:"max_retry_backoff" toml:"max_retry_backoff" json:"max_retry_backoff" bson:"max_retry_backoff"`             //
	Timeout            int `yaml:"timeout" toml:"timeout" json:"timeout" bson:"timeout"`                                                     //
	KeepAlive          int `yaml:"keep_alive" toml:"keep_alive" json:"keep_alive" bson:"keep_alive"`                                         //
	MinIdleConns       int `yaml:"min_idle_conns" toml:"min_idle_conns" json:"min_idle_conns" bson:"min_idle_conns"`                         //
}
