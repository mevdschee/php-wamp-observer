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
}

func New(stats *metrics.Metrics) *Tracking {
	t := Tracking{
		msgIds:   map[string]time.Time{},
		msgNames: map[string]string{},
		timers:   map[string]*time.Timer{},
		stats:    stats,
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

func (t *Tracking) Track(protocol, direction, messageString string) error {
	var message []any
	err := json.Unmarshal([]byte(messageString), &message)
	if err != nil {
		return fmt.Errorf("malformed message: %v", messageString)
	}
	msgType := int(message[0].(float64))
	msgId := message[1].(string)
	if msgType == 2 {
		msgName := message[2].(string)
		t.Add(msgId, msgName, time.Now(), 300*time.Millisecond, func() {
			start, msgName, ok := t.Del(msgId)
			if ok {
				duration := time.Since(start).Seconds()
				t.stats.Add(protocol+"_"+direction+"_timeouts", "message", msgName, duration)
			}
		})
	}
	if msgType == 3 {
		start, msgName, ok := t.Del(msgId)
		if ok {
			duration := time.Since(start).Seconds()
			t.stats.Add(protocol+"_"+direction+"_responses", "message", msgName, duration)
		}
	}
	if msgType == 4 {
		start, msgName, ok := t.Del(msgId)
		if ok {
			duration := time.Since(start).Seconds()
			t.stats.Add(protocol+"_"+direction+"_errors", "message", msgName, duration)
		}
	}
	return nil
}
