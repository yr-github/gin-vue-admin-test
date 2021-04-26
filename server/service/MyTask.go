package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateMyTask
//@description: 创建MyTask记录
//@param: mytask model.MyTask
//@return: err error

func CreateMyTask(mytask model.MyTask) (err error) {
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
	err = global.GVA_DB.Delete(&[]model.MyTask{},"id in ?",ids.Ids).Error
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
	if	info.UserId != "" {
		db = db.Where("userid = ?", info.UserId)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&mytasks).Error
	return err, mytasks, total
}