package main

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"time"
	"wdkj/account/internal/server"
	"wdkj/account/utils"
	"wdkj/account/utils/log"
)

func main () {
	logger := log.InitLevelLogger(zapcore.InfoLevel, fmt.Sprintf("account_%s.log", time.Now().Format(utils.DateTimeLayoutFormat)))
	defer logger.Sync()

	if err := server.Builder().NewRouter().Run(":1011"); err != nil {
		panic(err)
	}
}