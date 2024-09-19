package tracking

import (
	"sync"
	"time"
)

type Tracking struct {
	mutex    sync.Mutex
	msgIds   map[string]time.Time
	msgNames map[string]string
	timers   map[string]*time.Timer
}

func New() *Tracking {
	t := Tracking{
		msgIds:   map[string]time.Time{},
		msgNames: map[string]string{},
		timers:   map[string]*time.Timer{},
	}
	return &t
}

func (t *Tracking) Add(msgId string, msgName string, val time.Time, timeout time.Duration, timeoutFunc func()) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.msgIds[msgId] = val
	t.msgNames[msgId] = msgName
	t.timers[msgId] = time.AfterFunc(timeout, timeoutFunc)
}

func (t *Tracking) Del(msgId string) (time.Time, string, bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	val := t.msgIds[msgId]
	msgName := t.msgNames[msgId]
	timer := t.timers[msgId]
	delete(t.msgIds, msgId)
	delete(t.msgNames, msgId)
	delete(t.timers, msgId)
	if timer == nil {
		return time.Now(), msgName, false
	}
	timer.Stop()
	return val, msgName, true
}

func (t *Tracking) Len() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return len(t.msgIds)
}
