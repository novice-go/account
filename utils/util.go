package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成6位随机字符
func GenNumber() string {
	// 目前生成六位数
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

// 获取msg id，当前纳秒后加6位随机数
func GetMsgId() string {
	curNano := time.Now().UnixNano()
	r := rand.New(rand.NewSource(curNano))
	return fmt.Sprintf("%d%06v", curNano, r.Int31n(1000000))
}

func CombineStr(data ...string) string {
	var str string
	for _, v := range data {
		str = str + v
	}

	return str
}

func GetRedisUrl() string {
	return RedisHost + ":" + RedisPort
}
