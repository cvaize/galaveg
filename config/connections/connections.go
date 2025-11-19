package connections

type Config struct {
	//Mariadb *MariadbConfig
	//Mongodb *MongodbConfig
	Mysql *MysqlConfig
	//Pgsql   *PgsqlConfig
	Redis *RedisConfig
	//Sqlite  *SqliteConfig
	//Valkey  *ValkeyConfig
}

func NewConfig() (*Config, error) {
	return &Config{
		//Mariadb: NewMariadbConfig(),
		//Mongodb: NewMongodbConfig(),
		Mysql: NewMysqlConfig(),
		//Pgsql:   NewPgsqlConfig(),
		Redis: NewRedisConfig(),
		//Sqlite:  NewSqliteConfig(),
		//Valkey:  NewValkeyConfig(),
	}, nil
}

func MustConfig() *Config {
	c, e := NewConfig()
	if e != nil {
		panic(e)
	}
	return c
}
