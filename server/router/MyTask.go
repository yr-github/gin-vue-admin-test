package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitMyTaskRouter(Router *gin.RouterGroup) {
	MyTaskRouter := Router.Group("mytask").Use(middleware.OperationRecord())
	{
		MyTaskRouter.POST("createMyTask", v1.CreateMyTask)  // 新建MyTask
		MyTaskRouter.DELETE("deleteMyTask", v1.DeleteMyTask) // 删除MyTask
		MyTaskRouter.DELETE("deleteMyTaskByIds", v1.DeleteMyTaskByIds) // 批量删除MyTask
		MyTaskRouter.PUT("updateMyTask", v1.UpdateMyTask)    // 更新MyTask
		MyTaskRouter.GET("findMyTask", v1.FindMyTask)        // 根据ID获取MyTask
		MyTaskRouter.GET("getMyTaskList", v1.GetMyTaskList)  // 获取MyTask列表
	}
}
