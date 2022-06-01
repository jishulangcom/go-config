package config

type PostgreSqlPoolCnfDto struct {
	MaxConn int `yaml:"max_conn" toml:"max_conn" json:"max_conn" bson:"max_conn"` // 最大连接
}

var PostgreSqlPoolCnf = PostgreSqlPoolCnfDto{
	MaxConn: 10,
}
