package session

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type RdsSessionMgr struct {
	SessionMap map[string]Session
	rwLock     sync.RWMutex
	addr       string
	pwd        string
	pool       *redis.Client
}

func NewRdsSessionMgr(addr, pwd string) *RdsSessionMgr {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       0,
	})

	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()
	_, _ = rdb.Ping(ctx).Result()

	return &RdsSessionMgr{
		SessionMap: make(map[string]Session, 1024),
		pool:       rdb,
	}
}

func (r *RdsSessionMgr) Init(addr string, options ...string) (err error) {
	return
}

func (r *RdsSessionMgr) CreateSession() (session Session, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	id := uuid.NewV4()
	session = NewMemorySession(id.String())
	r.SessionMap[id.String()] = session

	return
}

func (r *RdsSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()

	session, ok := r.SessionMap[sessionId]
	if !ok {
		err = errors.New("不存在的 session id")
	}

	return
}
