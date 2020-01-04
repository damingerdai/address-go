package config

type DBConfig struct {
	Host, Port, User, Password, Dbname string
}

var conf *DBConfig

func initDBConfig() {
	env := New()
	a := createDBConfig(env)
	conf = &a
}

func createDBConfig(e *Env) DBConfig {
	return DBConfig{Host: e.Host(), Port: e.Port(), User: e.User(), Password: e.Password(), Dbname: e.Db()}
}

func GetDBConfig() *DBConfig {
	if conf == nil {
		initDBConfig()
	}
	return conf
}
