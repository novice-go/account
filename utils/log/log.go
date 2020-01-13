package log

import(
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	logDirPath = "/tmp/log/"
)

var Logger *zap.Logger

func InitLogger () {
	logger, err := zap.NewProductionConfig().Build()
	if err != nil {
		panic(err)
	}

	Logger = logger
}

func InitLevelLogger (level zapcore.Level, fileName string) *zap.Logger {
	if err := genLogFile(fileName); err != nil {
		panic(err)
	}

	conf := zap.NewProductionConfig()
	conf.Level.SetLevel(level)

	conf.OutputPaths = []string{logDirPath+fileName}
	conf.ErrorOutputPaths = []string{logDirPath+"err_"+fileName}

	logger, err := conf.Build()
	if err != nil {
		panic(err)
	}

	Logger = logger

	return logger
}


func genLogFile (fileName string) error {
	if err := os.MkdirAll(logDirPath, os.ModePerm); err != nil {
		return err
	}

	 _, err := os.Create(logDirPath+fileName)
	 return err
}