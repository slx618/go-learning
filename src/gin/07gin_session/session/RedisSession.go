package session

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

const (
	SessionFlagNone = iota
	SessionFlagModify
)

type RdsSession struct {
	SessionId string
	Data      map[string]interface{}
	rwLock    sync.RWMutex
	pool      *redis.Client
	flag      int
}

func NewRdsSession(id string) *RdsSession {
	s := &RdsSession{
		SessionId: id,
		Data:      make(map[string]interface{}, 16),
	}
	return s
}

func (r *RdsSession) Set(key string, val interface{}) (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()

	r.Data[key] = val
	r.flag = SessionFlagModify
	return
}

//Get(key string) (interface{}, error)
//
//Del(key string) error
//
//Save() error
func (r *RdsSession) Get(key string) (val interface{}, err error) {
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()

	val, ok := r.Data[key]
	if !ok {
		r.loadFromRedis()
		err = errors.New("值为空")
		return
	}

	return
}

func (r *RdsSession) Del(key string) (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()

	delete(r.Data, key)

	return
}

func (r *RdsSession) Save() (err error) {

	r.rwLock.RLock()
	defer r.rwLock.RUnlock()

	if r.flag == SessionFlagNone {
		return
	}

	data, err := json.Marshal(r.Data)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_ = r.pool.Set(ctx, r.SessionId, data, redis.KeepTTL).Err()
	r.flag = SessionFlagNone
	return nil
}

func (r RdsSession) loadFromRedis() (err error) {
	ctx, canncel := context.WithTimeout(context.Background(), time.Second*5)
	defer canncel()
	data, _ := r.pool.Get(ctx, r.SessionId).Result()
	_ = json.Unmarshal([]byte(data), &r.Data)

	return
}
