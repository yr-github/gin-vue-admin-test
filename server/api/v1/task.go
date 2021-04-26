package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags Task
// @Summary 创建Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Task true "创建Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task/createTask [post]
func CreateTask(c *gin.Context) {
	var task model.Task
	_ = c.ShouldBindJSON(&task)
	if err := service.CreateTask(task); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Task
// @Summary 删除Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Task true "删除Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /task/deleteTask [delete]
func DeleteTask(c *gin.Context) {
	var task model.Task
	_ = c.ShouldBindJSON(&task)
	if err := service.DeleteTask(task); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Task
// @Summary 批量删除Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /task/deleteTaskByIds [delete]
func DeleteTaskByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteTaskByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Task
// @Summary 更新Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Task true "更新Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /task/updateTask [put]
func UpdateTask(c *gin.Context) {
	var task model.Task
	_ = c.ShouldBindJSON(&task)
	if err := service.UpdateTask(task); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Task
// @Summary 用id查询Task
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Task true "用id查询Task"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /task/findTask [get]
func FindTask(c *gin.Context) {
	var task model.Task
	_ = c.ShouldBindQuery(&task)
	if err, retask := service.GetTask(task.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retask": retask}, c)
	}
}

// @Tags Task
// @Summary 分页获取Task列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TaskSearch true "分页获取Task列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /task/getTaskList [get]
func GetTaskList(c *gin.Context) {
	var pageInfo request.TaskSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetTaskInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
