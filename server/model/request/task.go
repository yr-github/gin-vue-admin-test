package request

import "gin-vue-admin/model"

type TaskSearch struct{
    model.Task
    PageInfo
}