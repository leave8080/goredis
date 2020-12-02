package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var Address = []string{"127.0.0.1:16379"}
var Network = "tcp"

// redis redis config.
type RedisConfig struct {
	Addrs    []string
	Password string
	DB       int
}
type Redis struct {
	Redis *redis.Client
}

func InitRedis(c *RedisConfig) (r *Redis) {
	rd := redis.NewClient(&redis.Options{
		Addr:     c.Addrs[0],
		Password: c.Password,
		DB:       c.DB,
	})
	_, err := rd.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("failed to connect redis error:%+v", err))
	}
	r = &Redis{Redis: rd}
	return r
}
func main() {
	var key = "key2"
	ctx := context.Background()
	c := &RedisConfig{
		Addrs: Address,
		Password: "123456",
		DB: 10,
	}
	r := InitRedis(c)
//ExpireAt
	{
		if err := r.Redis.Set(ctx,key,100,0).Err();err !=nil{
			log.Println("err",err)
			return
		}
		ss,err := r.Redis.Get(ctx,key).Int()
		if err != nil{
			log.Println("get err",err)
			return

		}
		log.Println("key:",ss)
		//timekey := time.Now().
		//r.Redis.ExpireAt(ctx,key,timekey)
	}

/*
// SetRange
	{
		if err :=  r.Redis.Set(ctx,key,100,time.Hour).Err();err != nil{
			log.Println("err",err)
			return
		}
		ss,err := r.Redis.Get(ctx,key).Int()
		if err != nil{
			log.Println("get err",err)
			return

		}
		log.Println("key:",ss)
		r.Redis.SetRange(ctx,key,0,"200")

		sss,err := r.Redis.Get(ctx,key).Int()
		if err != nil{
			log.Println("get err",err)
			return

		}
		log.Println("key:",sss)
	}


 */

}
