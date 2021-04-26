// 自动生成模板Task
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Task struct {
      global.GVA_MODEL
      StartUserid  int `json:"startUserid" form:"startUserid" gorm:"column:start_userid;comment:;type:int;size:10;"`
      EndDepartment  string `json:"endDepartment" form:"endDepartment" gorm:"column:end_department;comment:;type:varchar(191);size:191;"`
      EndUserid  int `json:"endUserid" form:"endUserid" gorm:"column:end_userid;comment:;type:int;size:10;"`
      TaskTittle  string `json:"taskTittle" form:"taskTittle" gorm:"column:task_tittle;comment:;type:varchar(191);size:191;"`
}


func (Task) TableName() string {
  return "task"
}


// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type TaskWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Task   `json:"business"`
// }

// func (Task) TableName() string {
// 	return "task"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["task"] = func() model.GVA_Workflow {
//   return new(model.TaskWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["task"] = func() interface{} {
// 	return new(model.Task)
// }
