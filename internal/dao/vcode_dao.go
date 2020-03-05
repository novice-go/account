package dao

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
	"wdkj/account/model"
	"wdkj/account/utils"
)

type VCodeDB interface {
	SaveVCode(code *model.VCode) error
	QueryVCodeCountByPhone(phone, typ, status string) (count int, err error)
	SaveVCodeFlow(flow *model.SMSFlow) error
	QueryLastVCode(map[string]interface{}) (*model.VCode, error)
	UpdateLastVCode(fields map[string]interface{}, condition *model.VCode) error
}

type VCodeDAO struct {
	db    VCodeDB
	cache *redis.Client
}

func NewVCodeDAO(db VCodeDB, cache *redis.Client) *VCodeDAO {
	return &VCodeDAO{db: db, cache: cache}
}

func (dao *VCodeDAO) SaveVCode(data *model.VCode) error {
	return  dao.db.SaveVCode(data)
}

func (dao *VCodeDAO) UpdateVCode(fields map[string]interface{}, condition *model.VCode) error {
	return dao.db.UpdateLastVCode(fields, condition)
}

func (dao *VCodeDAO) QueryVCode(phone, typ string) (*model.VCode, error) {
	return dao.db.QueryLastVCode(map[string]interface{}{"phone": phone, "v_code_type": typ})
}

/*
	数据库返回"not found record"时, count=0
*/
func (dao *VCodeDAO) GetVCodeCountByPhone(phone, typ, status string) (count uint, err error) {
	// cache get
	key := utils.CombineStr(phone, typ)
	// cache get
	cnt, err := dao.getCacheVCodeCount(key)
	if err == nil {
		return cnt, nil
	}
	// TODO log err

	c, err := dao.db.QueryVCodeCountByPhone(phone, typ, status)
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return 0, nil
	}

	return uint(c), err
}

func (dao *VCodeDAO) SaveVCodeFlow(flow *model.SMSFlow) error {
	if err := dao.db.SaveVCodeFlow(flow); err != nil {
		return err
	}

	key := utils.CombineStr(flow.Phone, flow.SendType)
	// cache get
	cnt, err := dao.getCacheVCodeCount(key)
	if err != nil {
		return err
	}

	// cache set
	now := time.Now()
	nextDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, 1)
	if err := dao.cache.Set(key, cnt+1, nextDate.Sub(now)).Err(); err != nil {
		return err
	}

	return nil
}

func (dao *VCodeDAO) getCacheVCodeCount(key string) (uint, error) {
	result, err := dao.cache.Get(key).Result()
	switch err {
	case nil:

	case redis.Nil:
		result = "0"
	default:
		return 0, err
	}

	num, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}

	return uint(num), nil
}

func (dao *VCodeDAO) setCache(key, val string, expiration time.Duration) error {
	return dao.cache.Set(key, val, expiration).Err()
}

func (dao *VCodeDAO) getCache(key string) (string, error) {
	val, err := dao.cache.Get(key).Result()
	switch err {
	case nil:

	case redis.Nil:
		return "", nil
	default:
		return "", err
	}

	return val, nil
}
