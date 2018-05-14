package Nano

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

const (
	loopForever = -1
)

var (
	// default timer backlog
	timerBacklog = 1 << 8

	// timerManager manager for all timers
	timerManager = &struct {
		incrementID    int64            // auto increment id
		timers         map[int64]*Timer // all timers
		chClosingTimer chan int64       // timer for closing
		chCreatedTimer chan *Timer
	}{}

	// timerPrecision indicates the precision of timer, default is time.Second
	timerPrecision = time.Second

	// globalTicker represents global ticker that all cron job will be executed
	// in globalTicker.
	globalTicker *time.Ticker
)

type (
	// TimerFunc represents a function which will be called periodically in main
	// logic gorontine.
	TimerFunc func()

	// TimerCondition represents a checker that returns true when cron job needs
	// to execute
	TimerCondition interface {
		Check(now time.Time) bool
	}

	// Timer represents a cron job
	Timer struct {
		id        int64          // timer id
		fn        TimerFunc      // function that execute
		createAt  int64          // timer create time
		interval  time.Duration  // execution interval
		condition TimerCondition // condition to cron job execution
		elapse    int64          // total elapse time
		closed    int32          // is timer closed
		counter   int            // counter
	}
)

// NewCountTimer returns a new Timer containing a function that will be called
// with a period specified by the duration argument. After count times, timer
// will be stopped automatically, It adjusts the intervals for slow receivers.
// The duration d must be greater than zero; if not, NewCountTimer will panic.
// Stop the timer to release associated resources.
func NewCountTimer(interval time.Duration, count int, fn TimerFunc) *Timer {
	if fn == nil {
		panic("nano/timer: nil timer function")
	}
	if interval <= 0 {
		panic("non-positive interval for NewTimer")
	}

	id := atomic.AddInt64(&timerManager.incrementID, 1)
	t := &Timer{
		id:       id,
		fn:       fn,
		createAt: time.Now().UnixNano(),
		interval: interval,
		elapse:   int64(interval), // first execution will be after interval
		counter:  count,
	}

	// add to manager
	timerManager.chCreatedTimer <- t
	return t
}

// Stop turns off a timer. After Stop, fn will not be called forever
func (t *Timer) Stop() {
	if atomic.LoadInt32(&t.closed) > 0 {
		return
	}

	// guarantee that logic is not blocked
	if len(timerManager.chClosingTimer) < timerBacklog {
		timerManager.chClosingTimer <- t.id
		atomic.StoreInt32(&t.closed, 1)
	} else {
		t.counter = 0 // automatically closed in next Cron
	}
}

// TODO: if closing timers'count in single cron call more than timerBacklog will case problem.
func cron() {
	if len(timerManager.timers) < 1 {
		return
	}

	now := time.Now()
	unn := now.UnixNano()
	for id, t := range timerManager.timers {
		// prevent chClosingTimer exceed
		if t.counter == 0 {
			if len(timerManager.chClosingTimer) < timerBacklog {
				t.Stop()
			}
			continue
		}

		// condition timer
		if t.condition != nil {
			if t.condition.Check(now) {
				pexec(id, t.fn)
			}
			continue
		}

		// execute job
		if t.createAt+t.elapse <= unn {
			pexec(id, t.fn)
			t.elapse += int64(t.interval)

			// update timer counter
			if t.counter != loopForever && t.counter > 0 {
				t.counter--
			}
		}
	}
}

// execute job function with protection
func pexec(id int64, fn TimerFunc) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(fmt.Sprintf("Call timer function error, TimerID=%d, Error=%v", id, err))
			println(stack())
		}
	}()
	fn()
}

// NewTimer returns a new Timer containing a function that will be called
// with a period specified by the duration argument. It adjusts the intervals
// for slow receivers.
// The duration d must be greater than zero; if not, NewTimer will panic.
// Stop the timer to release associated resources.
func NewTimer(interval time.Duration, fn TimerFunc) *Timer {
	return NewCountTimer(interval, loopForever, fn)
}
