package tracking

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/mevdschee/php-observability/metrics"
)

type Tracking struct {
	mutex    sync.Mutex
	msgIds   map[string]time.Time
	msgNames map[string]string
	timers   map[string]*time.Timer
	stats    *metrics.Metrics
	timeout  time.Duration
}

func New(stats *metrics.Metrics, timeout time.Duration) *Tracking {
	t := Tracking{
		msgIds:   map[string]time.Time{},
		msgNames: map[string]string{},
		timers:   map[string]*time.Timer{},
		stats:    stats,
		timeout:  timeout,
	}
	return &t
}

func (t *Tracking) add(msgId string, msgName string, val time.Time, timeout time.Duration, timeoutFunc func()) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.msgIds[msgId] = val
	t.msgNames[msgId] = msgName
	t.timers[msgId] = time.AfterFunc(timeout, timeoutFunc)
}

func (t *Tracking) del(msgId string) (time.Time, string, bool) {
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

func (t *Tracking) Track(messageType, messageString string) error {
	var message []any
	err := json.Unmarshal([]byte(messageString), &message)
	if err != nil {
		return fmt.Errorf("malformed message: %v", messageString)
	}
	msgType := int(message[0].(float64))
	msgId := message[1].(string)
	if msgType == 2 {
		msgName := message[2].(string)
		t.add(msgId, msgName, time.Now(), t.timeout, func() {
			start, msgName, ok := t.del(msgId)
			if ok {
				duration := time.Since(start).Seconds()
				t.stats.Add(messageType+"_timeouts", "message", msgName, duration)
			}
		})
	}
	if msgType == 3 {
		start, msgName, ok := t.del(msgId)
		if ok {
			duration := time.Since(start).Seconds()
			t.stats.Add(messageType+"_responses", "message", msgName, duration)
		}
	}
	if msgType == 4 {
		start, msgName, ok := t.del(msgId)
		if ok {
			duration := time.Since(start).Seconds()
			t.stats.Add(messageType+"_errors", "message", msgName, duration)
		}
	}
	return nil
}
