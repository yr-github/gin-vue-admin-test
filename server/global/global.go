package global

import (
	"context"
	"crypto/sha1"
	"gin-vue-admin/config"
	"gin-vue-admin/mq"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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
	h := sha1.New()
	h.Write([]byte(value))
	bs := h.Sum(nil)
	err :=GVA_REDIS.Set(ctx,string(bs),value,0)
	return err.Err()
}
func RedisGet(key string) (string ,error){
	return key, nil
}
func RedisDel(value string) error{
	ctx:=context.TODO()
	h := sha1.New()
	h.Write([]byte(value))
	bs := h.Sum(nil)
	err := GVA_REDIS.Del(ctx,string(bs))
	return err.Err()
}
