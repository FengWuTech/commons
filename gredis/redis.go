package gredis

import (
	"encoding/json"
	"github.com/FengWuTech/commons/logger"
	"github.com/FengWuTech/commons/setting"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

var RedisConn *redis.Pool

// Setup Initialize the Redis instance
func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
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

	return nil
}

// Set a key/value
func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	var value []byte
	switch data.(type) {
	case int:
		value = []byte(strconv.Itoa(data.(int)))
	case string:
		value = []byte(data.(string))
	default:
		value, _ = json.Marshal(data)
	}

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func LikeDeletes(key string) error {
	conn := RedisConn.Get()
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

func Publish(channel string, message string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	return conn.Send("PUBLISH", channel, message)
}

func Subscribe(channel string, fun func(message string)) error {
	conn := RedisConn.Get()
	psc := redis.PubSubConn{Conn: conn}
	defer conn.Close()
	defer psc.Close()
	psc.Subscribe(channel)

	for {
		v := psc.Receive()
		switch v.(type) {
		case redis.Message:
			fun(string(v.(redis.Message).Data))
		case redis.Subscription:
		case error:
			//断网重连
			logger.Warnf("Subscribe异常 %v", v)
			time.Sleep(time.Second * 5)
			conn := RedisConn.Get()
			psc = redis.PubSubConn{Conn: conn}
			break
		}
	}
}
