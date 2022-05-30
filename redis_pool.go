package config

type RedisPoolCnfDto struct {
	PoolSize           int `yaml:"pool_size" toml:"pool_size" json:"pool_size" bson:"pool_size"`                                             // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
	MinIdleConns       int `yaml:"min_idle_conns" toml:"min_idle_conns" json:"min_idle_conns" bson:"min_idle_conns"`                         // 在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。
	DialTimeout        int `yaml:"dial_timeout" toml:"dial_timeout" json:"dial_timeout" bson:"dial_timeout"`                                 // 连接建立超时时间，默认5秒。
	ReadTimeout        int `yaml:"read_timeout" toml:"read_timeout" json:"read_timeout" bson:"read_timeout"`                                 // 读超时，默认3秒， -1表示取消读超时
	WriteTimeout       int `yaml:"write_timeout" toml:"write_timeout" json:"write_timeout" bson:"write_timeout"`                             // 写超时，默认等于读超时
	PoolTimeout        int `yaml:"pool_timeout" toml:"pool_timeout" json:"pool_timeout" bson:"pool_timeout"`                                 // 当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
	IdleCheckFrequency int `yaml:"idle_check_frequency" toml:"idle_check_frequency" json:"idle_check_frequency" bson:"idle_check_frequency"` // 闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
	IdleTimeout        int `yaml:"idle_timeout" toml:"idle_timeout" json:"idle_timeout" bson:"idle_timeout"`                                 // 闲置超时，默认5分钟，-1表示取消闲置超时检查
	MaxConnAge         int `yaml:"max_conn_age" toml:"max_conn_age" json:"max_conn_age" bson:"max_conn_age"`                                 // 连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
	MaxRetries         int `yaml:"max_retries" toml:"max_retries" json:"max_retries" bson:"max_retries"`                                     // 命令执行失败时，最多重试多少次，默认为0即不重试
	MinRetryBackoff    int `yaml:"min_retry_backoff" toml:"min_retry_backoff" json:"min_retry_backoff" bson:"min_retry_backoff"`             // 每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
	MaxRetryBackoff    int `yaml:"max_retry_backoff" toml:"max_retry_backoff" json:"max_retry_backoff" bson:"max_retry_backoff"`             // 每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
	Timeout            int `yaml:"timeout" toml:"timeout" json:"timeout" bson:"timeout"`                                                     //
	KeepAlive          int `yaml:"keep_alive" toml:"keep_alive" json:"keep_alive" bson:"keep_alive"`                                         //

}

var RedisPoolCnf RedisPoolCnfDto = RedisPoolCnfDto{
	PoolSize:           4,
	MinIdleConns:       10,
	DialTimeout:        5,
	ReadTimeout:        3,
	WriteTimeout:       3,
	PoolTimeout:        4,
	IdleCheckFrequency: 60,
	IdleTimeout:        5,
	MaxConnAge:         0,
	MaxRetries:         0,
	MinRetryBackoff:    8,
	MaxRetryBackoff:    512,
	Timeout:            5,
	KeepAlive:          5,

}