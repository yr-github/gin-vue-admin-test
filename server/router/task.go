package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTaskRouter(Router *gin.RouterGroup) {
	TaskRouter := Router.Group("task").Use(middleware.OperationRecord())
	{
		TaskRouter.POST("createTask", v1.CreateTask)   // 新建Task
		TaskRouter.DELETE("deleteTask", v1.DeleteTask) // 删除Task
		TaskRouter.DELETE("deleteTaskByIds", v1.DeleteTaskByIds) // 批量删除Task
		TaskRouter.PUT("updateTask", v1.UpdateTask)    // 更新Task
		TaskRouter.GET("findTask", v1.FindTask)        // 根据ID获取Task
		TaskRouter.GET("getTaskList", v1.GetTaskList)  // 获取Task列表
	}
}
