package service

import (
	"encoding/json"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateMyTask
//@description: 创建MyTask记录
//@param: mytask model.MyTask
//@return: err error
type MytaskReflect struct {
}

func (MytaskReflect) CreateMyTaskFromMq(sql string) error {
	//目前仅仅是handle错误 出了错误如何处理是个问题
	var mytask model.MyTask
	err := json.Unmarshal([]byte(sql), &mytask)
	if err != nil {
		global.GVA_LOG.Error("插入数据格式错误！")
		return err
	}
	err = global.GVA_DB.Create(&mytask).Error
	if err != nil {
		global.GVA_LOG.Error("数据库创建错误！")
		//此处数据库错误其实可以使用global.MQTODB <- handle(sql)
		//handle中将当前函数名与sql按之前的规则组合即可
		//但最好是传递到mq而不是直接调用 MQTODB
		return err
	}
	global.DBTOREDIS <- sql
	return err
}
func (MytaskReflect) CreateMyTask(mytask model.MyTask) (err error) {
	err = global.GVA_DB.Create(&mytask).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMyTask
//@description: 删除MyTask记录
//@param: mytask model.MyTask
//@return: err error

func DeleteMyTask(mytask model.MyTask) (err error) {
	err = global.GVA_DB.Delete(&mytask).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMyTaskByIds
//@description: 批量删除MyTask记录
//@param: ids request.IdsReq
//@return: err error

func DeleteMyTaskByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.MyTask{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateMyTask
//@description: 更新MyTask记录
//@param: mytask *model.MyTask
//@return: err error

func UpdateMyTask(mytask model.MyTask) (err error) {
	err = global.GVA_DB.Save(&mytask).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMyTask
//@description: 根据id获取MyTask记录
//@param: id uint
//@return: err error, mytask model.MyTask

func GetMyTask(id uint) (err error, mytask model.MyTask) {
	err = global.GVA_DB.Where("id = ?", id).First(&mytask).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMyTaskInfoList
//@description: 分页获取MyTask记录
//@param: info request.MyTaskSearch
//@return: err error, list interface{}, total int64

func GetMyTaskInfoList(info request.MyTaskSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.MyTask{})
	var mytasks []model.MyTask
	if info.UserId != "" {
		db = db.Where("userid = ?", info.UserId)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&mytasks).Error
	return err, mytasks, total
}
