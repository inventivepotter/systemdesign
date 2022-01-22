package decorator

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type cbStatus int

const (
	Closed cbStatus = iota
	HalfOpen
	Open
)

type ReqTracker struct {
	Id           string
	RetryCount   int
	SuccessCount int
	Status       cbStatus
	UpdatedAt    time.Time
}

func (rt *ReqTracker) tickErr() {
	rt.RetryCount++
	rt.UpdatedAt = time.Now()
}

func (rt *ReqTracker) tickSuc() {
	rt.SuccessCount++
	rt.UpdatedAt = time.Now()
}

func (rt *ReqTracker) changeStatus(st cbStatus) {
	rt.RetryCount = 0
	rt.SuccessCount = 0
	rt.Status = st
	rt.UpdatedAt = time.Now()
}

func (rt *ReqTracker) updateNow() {
	rt.UpdatedAt = time.Now()
}

func (rt *ReqTracker) resetRetryCount() {
	rt.RetryCount = 0
}

func (rt *ReqTracker) getRetryCount() int {
	return rt.RetryCount
}

func (rt *ReqTracker) getSuccessCount() int {
	return rt.SuccessCount
}

func (rt *ReqTracker) getUpdatedAt() time.Time {
	return rt.UpdatedAt
}

func newReqTracker(Id string) *ReqTracker {
	return &ReqTracker{
		Id:           Id,
		RetryCount:   0,
		SuccessCount: 0,
		Status:       Closed,
		UpdatedAt:    time.Now(),
	}
}

var tracker = make(map[string]*ReqTracker)

type CircuitBreakerRequest struct {
	Fetcher
	tripThreshold         int
	successCountThreshold int
	delay                 time.Duration
	mutex                 *sync.Mutex
}

func (cbr CircuitBreakerRequest) Fetch(args Args) (Response, error) {
	cbr.mutex.Lock()
	if _, ok := tracker[args["id"].(string)]; !ok {
		temp := newReqTracker(args["id"].(string))
		tracker[args["id"].(string)] = temp
	}
	cb := tracker[args["id"].(string)]
	//bs, _ := json.Marshal(tracker)
	//fmt.Println(string(bs))
	switch cb.Status {
	case Closed:
		res, err := cbr.Fetcher.Fetch(args)
		if err != nil {
			cb.tickErr()
			fmt.Println("error in closed circuit with retry count " + strconv.Itoa(cb.getRetryCount()))
			if cb.getRetryCount() == cbr.tripThreshold {
				fmt.Println("trip hit! circuit is open")
				cb.changeStatus(Open)
			}
		} else {
			cb.resetRetryCount()
			fmt.Println("success in closed circuit, so retry count is reset")
			defer cbr.mutex.Unlock()
			return res, nil
		}
	case HalfOpen:
		res, err := cbr.Fetcher.Fetch(args)
		if err != nil {
			cb.changeStatus(Open)
			fmt.Println("error in half open circuit, so circuit is open")
		} else {
			cb.tickSuc()
			fmt.Println("success in half open circuit with sccuess count " + strconv.Itoa(cb.getSuccessCount()))
			if cb.getSuccessCount() == cbr.successCountThreshold {
				cb.changeStatus(Closed)
				fmt.Println("success count reached, hence circuit is closed")
			}
			defer cbr.mutex.Unlock()
			return res, nil
		}
	case Open:
		if time.Since(cb.getUpdatedAt()) > cbr.delay {
			res, err := cbr.Fetcher.Fetch(args)
			if err != nil {
				cb.updateNow()
				fmt.Println("error in open circuit, resetting the time limit")
			} else {
				cb.changeStatus(HalfOpen)
				fmt.Println("success in open circuit, hence changing to half open circuit")
				defer cbr.mutex.Unlock()
				return res, nil
			}
		} else {
			fmt.Println("open circuit, fail fast error")
			defer cbr.mutex.Unlock()
			return nil, errors.New("circuit breaker resulted in open circuit")
		}
	}
	defer cbr.mutex.Unlock()
	return nil, errors.New("unhandled error in CB")
}

func NewCBO(f Fetcher, tripThreshold int, successCountThreshold int, delay time.Duration, mutex *sync.Mutex) CircuitBreakerRequest {
	return CircuitBreakerRequest{
		Fetcher:               f,
		tripThreshold:         tripThreshold,
		successCountThreshold: successCountThreshold,
		delay:                 delay,
		mutex:                 mutex,
	}
}
