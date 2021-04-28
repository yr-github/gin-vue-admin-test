package global

import (
	"context"
	"gin-vue-admin/mq"
	"go.uber.org/zap"
	"hash/fnv"

	"gin-vue-admin/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
	GVA_MQ		*mq.RabbitMQ
)

func RedisSet(value string)error{
	ctx:=context.TODO()
	h := fnv.New32a()
	h.Write([]byte(value))
	GVA_REDIS.Set(ctx,string(h.Sum32()),value,0)
	return nil
}
func RedisGet(key string) (string ,error){
	return key, nil
}
func RedisDel(value string) error{
	ctx:=context.TODO()
	h := fnv.New32a()
	h.Write([]byte(value))
	GVA_REDIS.Del(ctx,string(h.Sum32()))
	return nil
}
