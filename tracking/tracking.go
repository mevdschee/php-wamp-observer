package tracking

import (
	"sync"
	"time"
)

type Tracking struct {
	mutex  sync.Mutex
	msgIds map[string]time.Time
	timers map[string]*time.Timer
}

func New() *Tracking {
	t := Tracking{
		msgIds: map[string]time.Time{},
		timers: map[string]*time.Timer{},
	}
	return &t
}

func (t *Tracking) Add(name string, val time.Time, timeout time.Duration, timeoutFunc func()) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.msgIds[name] = val
	t.timers[name] = time.AfterFunc(timeout, timeoutFunc)
}

func (t *Tracking) Del(name string) (time.Time, bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	val := t.msgIds[name]
	timer := t.timers[name]
	delete(t.msgIds, name)
	delete(t.timers, name)
	if timer == nil {
		return time.Now(), false
	}
	timer.Stop()
	return val, true
}

func (t *Tracking) Len() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return len(t.msgIds)
}
