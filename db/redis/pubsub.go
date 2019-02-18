package gredis

import (
	"context"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

//ListenSub listen subscribe.
// onStart subscribe success called, onMessage each message recv called.
func ListenSub(ctx context.Context, addr, passwd string,
	onStart func() error,
	onMessage func(topic string, data []byte) error,
	channels []string) error {

	// connect ...
	cr := NewRedisClient(addr, passwd)
	conn := cr.p.Get()
	if conn.Err() != nil {
		log.Printf("redis ListenSub failed. %v", conn.Err())
		return conn.Err()
	}

	psc := redis.PubSubConn{Conn: conn}
	if err := psc.Subscribe(redis.Args{}.AddFlat(channels)...); err != nil {
		log.Printf("redis Subscribe failed. %v", err)
		return err
	}

	// recv subscribe data from redis
	done := make(chan error, 1)
	go func() {
		for {
			switch n := psc.Receive().(type) {
			case error:
				done <- n
				return
			case redis.Message:
				if err := onMessage(n.Channel, n.Data); err != nil {
					done <- err
					return
				}
			case redis.Subscription:
				switch n.Count {
				case len(channels):
					if err := onStart(); err != nil {
						done <- err
						return
					}
				case 0:
					done <- nil
					return
				}
			}
		}
	}()

	//block & health check
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	defer psc.Unsubscribe(channels)

	for {
		select {
		case <-ticker.C:
			if err := psc.Ping(""); err != nil {
				return err
			}
		case <-ctx.Done():
			// do noting ...
		case err := <-done:
			return err
		}
	}
}

