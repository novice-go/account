package utils

const (
	SMSMaxLimitNum = 5 // 每天每种类型最多发送短信次数

	DateFormat           = "2006-01-02"          // 日期格式化
	DateTimeFormat       = "2006-01-02 15:04:05" // 时间格式化
	DateLayoutFormat     = "20060102"            // 日期格式化, 不带横杆
	DateTimeLayoutFormat = "20060102150405"      // 时间格式化, 不带横杆

	// redis
	RedisHost = "localhost" // redis 地址
	RedisPort = "6379"      // redis 端口

	// vcode type
	VCodeTypeLogin    = "login"    // 登录
	VcodeTypeRegister = "register" // 注册
)

// 状态
const (
	StatusSuccess = "success"
	StatusFail    = "fail"
)
