package server

import (
	"github.com/go-redis/redis"
	"os"
	"wdkj/account/internal/dao"
	"wdkj/account/internal/dao/db"
	"wdkj/account/internal/service"
	account_service "wdkj/account/internal/service/account-service"
	vcode_service "wdkj/account/internal/service/vcode-service"
	"wdkj/account/model"
	"wdkj/account/utils/config"
	mysql_db "wdkj/account/utils/mysql-db"
)

func Builder() *Server {
	conf := getConfig()

	redisCache := getRedis(conf.RedisConf)
	mysqlDB := mysql_db.NewMysqlGormConn(&mysql_db.DBConfig{
		Host:     conf.DBConfig.Host,
		Port:     conf.DBConfig.Port,
		Name:     conf.DBConfig.User,
		Password: conf.DBConfig.Pw,
		DBName:   conf.DBConfig.DbName,
	})

	vcodeDB := db.NewVCodeDB(mysqlDB)

	accDAO := dao.NewAccountDAO()
	vcodDAO := dao.NewVCodeDAO(vcodeDB, redisCache)
	mockSender := vcode_service.NewMockSenderImpl() // 暂用mock

	vcodeService := vcode_service.NewVCodeService(vcodDAO, vcodDAO, mockSender)
	accService := account_service.NewAccountService(accDAO)

	loginManager := service.NewLoginManager(accService, vcodeService)

	return &Server{loginManger: loginManager}
}

func getConfig() *model.Config {
	resp := &model.Config{}
	if err := config.InitConfig(os.Getenv("GOPATH")+"/src/wdkj/account/config.yaml", resp); err != nil {
		panic(err)
	}

	return resp
}

func getRedis(conf model.RedisConf) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.Pw, // no password set
		DB:       conf.Db, // use default DB
	})

	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
	return client
}
