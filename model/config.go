package model

type Config struct {
	DBConfig  DBConfig  `json:"db" yaml:"db"`
	Conf      Conf      `json:"conf" yaml:"conf"`
	RedisConf RedisConf `json:"redis" yaml:"redis"`
}

type RedisConf struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
	Pw   string `json:"pw" yaml:"pw"`
	Db   int    `json:"db" yaml:"db"` // redis 存储的db
}

type DBConfig struct {
	Host   string `json:"host" yaml:"host"`
	Port   string `json:"port" yaml:"port"`
	User   string `json:"user" yaml:"user"`
	Pw     string `json:"pw" yaml:"pw"`
	DbName string `json:"db_name" yaml:"db_name"`
}

// TODO 根据config文件添加哦
type Conf struct {

}
