package server

import (
	"github.com/go-redis/redis"
	"wdkj/account/internal/dao"
	"wdkj/account/internal/dao/db"
	"wdkj/account/internal/service"
	account_service "wdkj/account/internal/service/account-service"
	vcode_service "wdkj/account/internal/service/vcode-service"
	"wdkj/account/utils"
	mysql_db "wdkj/account/utils/mysql-db"
)

func Builder() *Server {
	redisCache := getRedis()
	mysqlDB := mysql_db.NewMysqlGormConn()

	vcodeDB := db.NewVCodeDB(mysqlDB)

	accDAO := dao.NewAccountDAO()
	vcodDAO := dao.NewVCodeDAO(vcodeDB, redisCache)
	mockSender := vcode_service.NewMockSenderImpl() // 暂用mock

	vcodeService := vcode_service.NewVCodeService(vcodDAO, vcodDAO, mockSender)
	accService := account_service.NewAccountService(accDAO)

	loginManager := service.NewLoginManager(accService, vcodeService)

	return &Server{loginManger:loginManager}
}


func getRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     utils.GetRedisUrl(),
		Password: "", // no password set
		DB:       10,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}