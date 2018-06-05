package cache

import (
	"github.com/garyburd/redigo/redis"

	"time"
	"github.com/qq976739120/zhihu-golang-web/pkg/setting"
	"encoding/json"
)

var (
	RedisPool *redis.Pool
)

func InitRedis() {

	var (
		redis_address                      string
		max_idle, max_active, idle_timeout int
	)

	redis_address = setting.Redis_.RedisAddress
	max_idle = setting.Redis_.RedisMaxIdle
	max_active = setting.Redis_.RedisMaxActive
	idle_timeout = setting.Redis_.RedisIdleTimeout

	RedisPool = &redis.Pool{
		MaxIdle:     max_idle,
		MaxActive:   max_active,
		IdleTimeout: time.Duration(idle_timeout),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redis_address)
			if err != nil {
				return nil, err
			}
			if setting.Redis_.PassWord != "" {
				if _, err := c.Do("AUTH", setting.Redis_.PassWord); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

}

func CloseRedis() {
	defer RedisPool.Close()
}

func Set(key string, data interface{}, time int) (bool, error) {
	conn := RedisPool.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	reply, err := redis.Bool(conn.Do("SET", key, value))
	conn.Do("EXPIRE", key, time)
	return reply, err
}
func Exists(key string) bool {
	conn := RedisPool.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}
func Get(key string) ([]byte, error) {
	conn := RedisPool.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}
func Delete(key string) (bool, error) {
	conn := RedisPool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisPool.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}
	return nil
}
