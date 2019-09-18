package db

import (
	"github.com/go-redis/redis"
	"log"
)

var client * redis.Client

func init()  {
	client = redis.NewClient(&redis.Options{
		Addr:"127.0.0.1:6379",
		Password:"",
		DB:0,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatal("redis connect error!")
	}
}

func GetList(k string, len int) []string {
	l, _ := client.LRange(k, 0, int64(len-1)).Result()
	return l
}

func MGet(keys []string) []interface{} {
	vs, _ := client.MGet(keys...).Result()
	return vs
}

func Get(k string) string {
	v, _ := client.Get(k).Result()
	return v
}

func Set(k string, v interface{})  {
	client.Set(k, v, 0)
}

func LPush(k string, v interface{}) {
	client.LPush(k, v)
}