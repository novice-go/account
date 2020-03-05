package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"wdkj/account/model"
	"wdkj/account/utils/config"
	mysqlDb "wdkj/account/utils/mysql-db"
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
	db := mysqlDb.NewMysqlGormConn(getDBConfig())
	return db.DB.CreateTable(
		&model.Account{},
		&model.SMSFlow{},
		&model.VCode{},
	).Error
}

func drop(c *cli.Context) error {
	fmt.Println("drop")
	db := mysqlDb.NewMysqlGormConn(getDBConfig())
	return db.DB.DropTableIfExists(
		&model.Account{},
		&model.SMSFlow{},
		&model.VCode{},
	).Error
}

func migrate(c *cli.Context) error {
	fmt.Println("migrate")
	db := mysqlDb.NewMysqlGormConn(getDBConfig())
	db.DB.LogMode(true)
	return db.DB.AutoMigrate(
		&model.Account{},
		&model.SMSFlow{},
		&model.VCode{},
	).Error
}

func getDBConfig() *mysqlDb.DBConfig {
	conf := &model.Config{}
	if err := config.InitConfig(os.Getenv("GOPATH")+"/src/wdkj/account/config.yaml", conf); err != nil {
		panic(err)
	}

	return &mysqlDb.DBConfig{
		Host:     conf.DBConfig.Host,
		Port:     conf.DBConfig.Port,
		Name:     conf.DBConfig.User,
		Password: conf.DBConfig.Pw,
		DBName:   conf.DBConfig.DbName,
	}
}
