package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTask
//@description: 创建Task记录
//@param: task model.Task
//@return: err error

func CreateTask(task model.Task) (err error) {
	err = global.GVA_DB.Create(&task).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTask
//@description: 删除Task记录
//@param: task model.Task
//@return: err error

func DeleteTask(task model.Task) (err error) {
	err = global.GVA_DB.Delete(&task).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTaskByIds
//@description: 批量删除Task记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTaskByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Task{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTask
//@description: 更新Task记录
//@param: task *model.Task
//@return: err error

func UpdateTask(task model.Task) (err error) {
	err = global.GVA_DB.Save(&task).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTask
//@description: 根据id获取Task记录
//@param: id uint
//@return: err error, task model.Task

func GetTask(id uint) (err error, task model.Task) {
	err = global.GVA_DB.Where("id = ?", id).First(&task).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTaskInfoList
//@description: 分页获取Task记录
//@param: info request.TaskSearch
//@return: err error, list interface{}, total int64

func GetTaskInfoList(info request.TaskSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Task{})
    var tasks []model.Task
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&tasks).Error
	return err, tasks, total
}