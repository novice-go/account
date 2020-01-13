package model

import "time"

type SMSFlow struct {
	Account
	Content    string     `gorm:"column:content" json:"content" name:"内容"`
	FlowNo     string     `gorm:"column:flow_no" json:"f_low_no" name:"流水号"`
	SendStatus string     `gorm:"column:send_status;index" json:"send_status" name:"发送状态"` // 索引
	SendType   string     `gorm:"column:send_type;index" json:"send_type" name:"验证码类型"`    // 索引
	SendAt     *time.Time `gorm:"column:send_at" json:"send_at" name:"发送时间"`
}
