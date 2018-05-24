package cache

import (
	"github.com/garyburd/redigo/redis"

	"fmt"
	"time"
	"github.com/qq976739120/zhihu-golang-web/pkg/setting"
)
var (
	RedisPool *redis.Pool
)
func init() {

	var (
		redis_address                      string
		max_idle, max_active, idle_timeout int
	)

	redis_address  = setting.Redis_.RedisAddress
	max_idle = setting.Redis_.RedisMaxIdle
	max_active = setting.Redis_.RedisMaxActive
	idle_timeout = setting.Redis_.RedisIdleTimeout

	RedisPool = &redis.Pool{
		MaxIdle:     max_idle,
		MaxActive:   max_active,
		IdleTimeout: time.Duration(idle_timeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redis_address)
		},
	}
	conn :=  RedisPool.Get()
	defer conn.Close()

	_, err := conn.Do("ping")
	if err != nil {
		fmt.Println("ping 不通")
		return
	}
}

func CloseRedis() {
	defer RedisPool.Close()
}