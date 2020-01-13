package db

import mysql_db "wdkj/account/utils/mysql-db"

type AccountDB struct {
	*mysql_db.MysqlDB
}