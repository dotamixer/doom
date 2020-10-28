package resync

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type Mutex struct {
	key    string
	client *redis.Client
}

func NewMutex(client *redis.Client, key string) *Mutex {
	return &Mutex{
		key:    key,
		client: client,
	}
}

func (m *Mutex) Lock() (err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), defaultOpts.timeOut)
	defer cancel()
	con := m.client.Conn(ctx)
	defer con.Close()
	ts := time.Now().Unix()
	for i := 1; i < defaultOpts.retry; i++ {
		boolCmd := con.SetNX(ctx, m.key, ts, defaultOpts.timeLock)
		if boolCmd.Val() {
			return nil
		}
		time.Sleep(time.Second)
	}
	return fmt.Errorf("failed to lock. key:[%s]", m.key)
}

func (m *Mutex) UnLock() error {
	ctx, cancel := context.WithTimeout(context.TODO(), defaultOpts.timeOut)
	defer cancel()
	con := m.client.Conn(ctx)
	defer con.Close()
	con.Del(ctx, m.key)
	return fmt.Errorf("failed to unlock. key:[%s]", m.key)
}
