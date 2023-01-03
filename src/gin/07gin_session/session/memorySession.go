package session

import (
	"errors"
	"sync"
)

type MemorySession struct {
	SessionId string
	Data      map[string]interface{}
	rwLock    sync.RWMutex
}

func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		SessionId: id,
		Data:      make(map[string]interface{}, 16),
	}
	return s
}

func (m *MemorySession) Set(key string, val interface{}) (err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()

	m.Data[key] = val
	return
}

//Get(key string) (interface{}, error)
//
//Del(key string) error
//
//Save() error
func (m *MemorySession) Get(key string) (val interface{}, err error) {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()

	val, ok := m.Data[key]
	if !ok {
		err = errors.New("值为空")
		return
	}

	return
}

func (m *MemorySession) Del(key string) (err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()

	delete(m.Data, key)

	return
}

func (m *MemorySession) Save() error {
	return nil
}
