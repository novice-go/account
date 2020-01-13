package mysql_db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

/*
	mysql 使用gorm连接
*/

func NewMysqlGormConn(conf ...*DBConfig) *MysqlDB {
	var mysqlDB MysqlDB
	if len(conf) == 0 {
		mysqlDB = MysqlDB{
			Config: &DBConfig{
				Host:     "127.0.0.1",
				Port:     "3306",
				Name:     "root",
				Password: "1234",
				DBName:   "account",
			},
		}
	} else {
		mysqlDB = MysqlDB{
			Config: conf[0],
		}
	}

	mysqlDB.DB = mysqlDB.openGormConn()
	return &mysqlDB
}

func (d *MysqlDB) openGormConn() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		d.Config.Name, d.Config.Password, d.Config.Host, d.Config.Port, d.Config.DBName))
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour * 3)

	//defer db.Close()
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}
	return db
}
