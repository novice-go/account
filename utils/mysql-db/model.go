package mysql_db

import "github.com/jinzhu/gorm"

type MysqlDB struct {
	DB     *gorm.DB
	Config *DBConfig
}

type DBConfig struct {
	Host     string // 地址
	Port     string // 端口号
	Name     string // 数据库用户名
	Password string // 数据库密码
	DBName   string // 数据库名称
}


func (d *MysqlDB) GetDB() *gorm.DB {
	return d.DB
}