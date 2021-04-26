package request

import "gin-vue-admin/model"

type MyTaskSearch struct{
    model.MyTask
    PageInfo
    UserId string `form:"userid" json:"user_id"`
}



