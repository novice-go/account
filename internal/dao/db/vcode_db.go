package db

import (
	"github.com/jinzhu/now"
	"wdkj/account/model"
	mysql_db "wdkj/account/utils/mysql-db"
)

type VCodeDB struct {
	*mysql_db.MysqlDB
}

func NewVCodeDB(mysqlDB *mysql_db.MysqlDB) *VCodeDB {
	return &VCodeDB{MysqlDB: mysqlDB}
}

// 查询最新一条记录
func (db *VCodeDB) QueryLastVCode(condition map[string]interface{}) (*model.VCode, error) {
	var resp *model.VCode
	err := db.GetDB().Where(condition).Last(&resp).Error
	return resp, err
}

// 更新最新一条记录
func (db *VCodeDB) UpdateLastVCode(fields map[string]interface{}, condition *model.VCode) error {
	return db.GetDB().Model(&model.VCode{}).Where(condition).Update(fields).Error
}

// 创建
func (db *VCodeDB) SaveVCode(code *model.VCode) error {
	return db.GetDB().Create(code).Error
}

// 根据条件统计
func (db *VCodeDB) QueryVCodeCountByPhone(phone, typ, status string) (count int, err error) {
	err = db.GetDB().Model(&model.SMSFlow{}).Where("phone=?", phone).
		Where("send_type=?", typ).
		Where("send_status=?", status).
		Where("created_at BETWEEN ? AND ?", now.BeginningOfDay(), now.EndOfDay()).
		Count(&count).Error
	return
}

func (db *VCodeDB) SaveVCodeFlow(flow *model.SMSFlow) error {
	return db.GetDB().Create(flow).Error
}




