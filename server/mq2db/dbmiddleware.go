package mq2db

import (
	"gin-vue-admin/global"
	"gin-vue-admin/service"
	"reflect"
)

func resetstr(str string) (string, string) {
	return str, "CreateMyTaskFromMq"
}

// MqToDbMiddleware
//从mq读取数据，根据数据获取执行函数名，利用反射执行该函数
//并且可以执行多个协程来处理 MQTODB
func MqToDbMiddleware() {
	go func() {
		for data := range global.MQTODB {
			value, method := resetstr(data)
			t := reflect.ValueOf(service.MytaskReflect{})
			m := t.MethodByName(method)
			args := []reflect.Value{reflect.ValueOf(value)}
			go m.Call(args)
			//if err[0].IsNil() {
			//	//不应该在这里塞入，应该在db操作塞入
			//	//global.DBTOREDIS<-value
			//	global.GVA_LOG.Info(method+"  执行成功")
			//}else {
			//	//TODO error 此处当传入错误参数到数据库后，此处只能触发一次
			//	//猜测是因为使用range此channel，那么在这里就不能重新给channel传值
			//	//而且好像global.MQTODB <- data导致了channel异常错误，此处只能在数据库那里做插入异常处理了
			//	//println(len(global.MQTODB))
			//	global.GVA_LOG.Error("")
			//	//以下错误代码
			//	//global.MQTODB <- data
		}
	}()
}
