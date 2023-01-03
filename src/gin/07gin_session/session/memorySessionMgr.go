package session

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type MemorySessionMgr struct {
	SessionMap map[string]Session
	rwLock     sync.RWMutex
}

func NewMemorySessionMgr() *MemorySessionMgr {
	return &MemorySessionMgr{
		SessionMap: make(map[string]Session, 1024),
	}
}

func (m *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

func (m *MemorySessionMgr) CreateSession() (session Session, err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	id := uuid.NewV4()
	session = NewMemorySession(id.String())
	m.SessionMap[id.String()] = session

	return
}

func (m *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	session, ok := m.SessionMap[sessionId]
	if !ok {
		err = errors.New("不存在的 session id")
	}

	return
}
