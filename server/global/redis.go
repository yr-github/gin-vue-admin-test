package global

import (
	"context"
	"crypto/sha1"
	"time"
)

func RedisSetByValue(value string,expiration time.Duration)error{
	ctx:=context.TODO()
	h := sha1.New()
	h.Write([]byte(value))
	bs := h.Sum(nil)
	err :=GVA_REDIS.Set(ctx,string(bs),value,expiration).Err()
	return err
}

func RedisDelByValue(value string) error{
	ctx:=context.TODO()
	h := sha1.New()
	h.Write([]byte(value))
	bs := h.Sum(nil)
	err := GVA_REDIS.Del(ctx,string(bs)).Err()
	return err
}
// go协程删除
func RedisDelByValueInChannel() {
	ctx:=context.TODO()
	go func() {
		for value :=range DBTOREDIS{
			h := sha1.New()
			h.Write([]byte(value))
			bs := h.Sum(nil)
			GVA_REDIS.Del(ctx,string(bs)).Err()
		}
	}()
}


func RedisSetByKey(key string,value string,expiration time.Duration)error{
	ctx:=context.TODO()
	err :=GVA_REDIS.Set(ctx,key,value,expiration).Err()
	return err
}

func RedisDelByKey(key string) error{
	ctx:=context.TODO()
	err := GVA_REDIS.Del(ctx,key).Err()
	return err
}

func RedisGetByKey(key string) (string ,error){
	ctx:=context.TODO()
	value ,err:=GVA_REDIS.Get(ctx,key).Result()
	return value, err
}