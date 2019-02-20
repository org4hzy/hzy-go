# hzy
用于游戏功能的部分接口go函数

收集有用的go开源项目

```text
格式化日志系统 https://github.com/sirupsen/logrus.git
```
like this :
```go
log.WithFields(log.Fields{
    "animal": "walrus",
    "size":   10,
  }).Info("A group of walrus emerges from the ocean")
```
```text
output:
time="2019-02-20T09:39:33+08:00" level=info msg="A group of walrus emerges from the ocean" animal=walrus size=10
```

```text
go-redis接口 https://github.com/gomodule/redigo.git
```
like this :
```go
// get redis pool
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
```

```text
go-mysql驱动 https://github.com/go-sql-driver/mysql.git
go-protobuf https://github.com/golang/protobuf.git
go-colortext https://github.com/daviddengcn/go-colortext.git
```
