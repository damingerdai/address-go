package config

type DBConfig struct {
	Host, Port, User, Password, Dbname string
}

func NewDBConfig(e *Env) DBConfig {
	return DBConfig{Host: e.Host(), Port: e.Port(), User: e.User(), Password: e.Password(), Dbname: e.Db()}
}
