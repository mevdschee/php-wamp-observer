package tracking

import (
	"sync"
	"time"
)

type Tracking struct {
	mutex  sync.Mutex
	msgIds map[string]time.Time
}

func New() *Tracking {
	t := Tracking{
		msgIds: map[string]time.Time{},
	}
	return &t
}

func (t *Tracking) Add(name string, val time.Time) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.msgIds[name] = val
}

func (t *Tracking) Del(name string) time.Time {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	val, exists := t.msgIds[name]
	if !exists {
		return time.Now()
	}
	return val
}
