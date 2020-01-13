package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"wdkj/account/model"
	mysql_db "wdkj/account/utils/mysql-db"
)

func main() {
	fmt.Println("db action...")
	app := cli.NewApp()
	app.Name = "db action"
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "action", Usage: "migrate, create, drop"},
	}
	app.Commands = []cli.Command{
		{Name: "create", Usage: "create db", Action: create},
		{Name: "drop", Usage: "drop db", Action: drop},
		{Name: "migrate", Usage: "migrate db", Action: migrate},
	}
	app.Action = cli.ShowAppHelp
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func create(c *cli.Context) error {
	fmt.Println("create")
	db := mysql_db.NewMysqlGormConn()
	return db.DB.CreateTable(
		&model.Account{},
		&model.SMSFlow{},
		&model.VCode{},
	).Error
}

func drop(c *cli.Context) error {
	fmt.Println("drop")
	db := mysql_db.NewMysqlGormConn()
	return db.DB.DropTableIfExists().Error
}

func migrate(c *cli.Context) error {
	fmt.Println("migrate")
	db := mysql_db.NewMysqlGormConn()
	db.DB.LogMode(true)
	return db.DB.AutoMigrate(
		&model.Account{},
		&model.SMSFlow{},
		&model.VCode{},
		).Error
}
