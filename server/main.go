package main

import (
	"gin-vue-admin/core"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"gin-vue-admin/mq2db"
	_ "net/http/pprof"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	//ch := make(chan int)
	//ch <- 1
	//ch <- 2
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	global.GVA_MQ = global.Rabbit()	  // rabbit 简单封装
	global.MQTODB = make(chan string) // mq消费者传递给db操作
	global.DBTOREDIS = make(chan string) //db操作传递给redis删除

	if global.GVA_DB != nil {
		initialize.MysqlTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	//查看
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()
	global.GVA_MQ.MqReceive("hello")
	global.RedisDelByValueInChannel()
	mq2db.MqToDbMiddleware()
	core.RunWindowsServer()
}
