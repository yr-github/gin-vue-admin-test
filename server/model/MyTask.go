// 自动生成模板MyTask
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type MyTask struct {
      global.GVA_MODEL
      Userid  int `json:"userid" form:"userid" gorm:"column:userid;comment:;type:int;"`
      Endid  int `json:"endid" form:"endid" gorm:"column:endid;comment:;type:int;"`
      EndDepartment  int `json:"endDepartment" form:"endDepartment" gorm:"column:end_department;comment:;type:int;"`
}



// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type MyTaskWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	MyTask   `json:"business"`
// }

// func (MyTask) TableName() string {
// 	return ""
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["mytask"] = func() model.GVA_Workflow {
//   return new(model.MyTaskWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["mytask"] = func() interface{} {
// 	return new(model.MyTask)
// }
