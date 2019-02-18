package gredis

import (
	"errors"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

//参考 beego cache .  https://github.com/astaxie/beego.git

// Cache redis object
type Cache struct {
	p      *redis.Pool
	addr   string
	passwd string
	dbNum  int
}

// NewRedisClient  client
func NewRedisClient(addr, passwd string) *Cache {
	return &Cache{
		addr:   addr,
		passwd: passwd,
		dbNum:  0,
	}
}

// Start init redis connection pool
func (c *Cache) Start() {
	// connect
	c.p = &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 180 * time.Second,
		Dial: func() (conn redis.Conn, err error) {
			log.Printf("connect redis %s %s %d ...", c.addr, c.passwd, c.dbNum)

			conn, err = redis.Dial("tcp", c.addr)
			if err != nil {
				return nil, err
			}

			if c.passwd != "" {
				if _, err = conn.Do("AUTH", c.passwd); err != nil {
					conn.Close()
					return nil, err
				}
			}

			if c.dbNum != 0 {
				if _, err = conn.Do("SELECT", c.dbNum); err != nil {
					conn.Close()
					return nil, err
				}
			}

			return
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}

			_, err := conn.Do("PING")
			return err
		},
	}

	if c.p.Get().Err() == nil {
		log.Printf("redis %s established.", c.addr)
	}
}

func (c *Cache) do(cmdName string, args ...interface{}) (reply interface{}, err error) {
	if len(args) <= 0 {
		return nil, errors.New("missing required arguments")
	}

	conn := c.p.Get()
	defer conn.Close()

	return conn.Do(cmdName, args...)
}

//Get cache from redis
func (c *Cache) Get(key string) interface{} {
	reply, err := c.do("GET", key)
	if err != nil {
		log.Printf("redis Get failed. %v", err)
		return nil
	}
	return reply
}

//GetMulti from redis
func (c *Cache) GetMulti(keys []string) []interface{} {
	args := make([]interface{}, 0)
	for _, v := range keys {
		args = append(args, v)
	}

	reply, err := redis.Values(c.do("MGET", args...))
	if err != nil {
		log.Printf("redis GetMulti failed. %v", err)
		return nil
	}
	return reply
}

//Put key to redis and set timeout.
func (c *Cache) Put(key string, value interface{}, timeout time.Duration) error {
	_, err := c.do("SETEX", key, timeout/time.Second, value)
	return err
}

//Incr key to redis
func (c *Cache) Incr(key string) error {
	_, err := c.do("INCR", key)
	return err
}

//Decr key to redis
func (c *Cache) Decr(key string) error {
	_, err := c.do("DECR", key)
	return err
}

//Delete key
func (c *Cache) Delete(key string) error {
	_, err := c.do("del", key)
	return err
}

//IsExist key
func (c *Cache) IsExist(key string) bool {
	v, err := redis.Bool(c.do("EXISTS", key))
	if err != nil {
		log.Printf("redis exists failed. %v", err)
	}
	return v
}

//Hgetall key
func (c *Cache) Hgetall(key string) []interface{} {
	v, err := redis.Values(c.do("HGETALL", key))
	if err != nil {
		log.Printf("redis hgetall failed. %v", err)
		return nil
	}
	return v
}

//Hmset key
func (c *Cache) Hmset(key string, values []string) error {
	args := make([]interface{}, 0)
	for _, v := range values {
		args = append(args, v)
	}

	_, err := c.do("HMSET", args...)
	return err
}
