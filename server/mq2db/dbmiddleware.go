package mq2db

import (
	"gin-vue-admin/global"
	"gin-vue-admin/service"
	"reflect"
)

func resetstr(str string) (string,string) {
	return str,"CreateMyTaskFromMq"
}
//从mq读取数据，根据数据获取执行函数名，利用反射执行该函数
func MqToDbMiddleware(){
	go func() {
		for data := range global.MQTODB{
			value ,method :=resetstr(data)
			t := reflect.ValueOf(service.MytaskReflect{})
			m := t.MethodByName(method)
			args := []reflect.Value{reflect.ValueOf(value)}
			err :=m.Call(args)
			println()
			if err[0].IsNil() {
				global.DBTOREDIS<-value
			}else {
				//TODO error 此处当传入错误参数到数据库后，此处只能触发一次
				//猜测是因为使用range此channel，那么在这里就不能重新给了
				//而且好像导致了channel异常错误，此处只能在数据库那里做插入异常处理了
				println(len(global.MQTODB))
				//以下错误代码
				//args := []reflect.Value{reflect.ValueOf(method)}
				//err := global.MQTODB <- data
				//if err != nil {
				//
				//}
			}
		}
	}()
}
