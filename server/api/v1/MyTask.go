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

// CreateMyTask @Tags MyTask
// @Summary 创建MyTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MyTask true "创建MyTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mytask/createMyTask [post]
func CreateMyTask(c *gin.Context) {
	var mytask model.MyTask
	_ = c.ShouldBindJSON(&mytask)

	global.GVA_MQ.Send()
	if err := service.CreateMyTask(mytask); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMyTask @Tags MyTask
// @Summary 删除MyTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MyTask true "删除MyTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mytask/deleteMyTask [delete]
func DeleteMyTask(c *gin.Context) {
	var mytask model.MyTask
	_ = c.ShouldBindJSON(&mytask)
	if err := service.DeleteMyTask(mytask); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags MyTask
// @Summary 批量删除MyTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MyTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /mytask/deleteMyTaskByIds [delete]
func DeleteMyTaskByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteMyTaskByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags MyTask
// @Summary 更新MyTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MyTask true "更新MyTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mytask/updateMyTask [put]
func UpdateMyTask(c *gin.Context) {
	var mytask model.MyTask
	_ = c.ShouldBindJSON(&mytask)
	if err := service.UpdateMyTask(mytask); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags MyTask
// @Summary 用id查询MyTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MyTask true "用id查询MyTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mytask/findMyTask [get]
func FindMyTask(c *gin.Context) {
	var mytask model.MyTask
	_ = c.ShouldBindQuery(&mytask)
	if err, remytask := service.GetMyTask(mytask.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remytask": remytask}, c)
	}
}

// @Tags MyTask
// @Summary 分页获取MyTask列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MyTaskSearch true "分页获取MyTask列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mytask/getMyTaskList [get]
func GetMyTaskList(c *gin.Context) {
	var pageInfo request.MyTaskSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetMyTaskInfoList(pageInfo); err != nil {
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
