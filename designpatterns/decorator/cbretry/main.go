package decorator

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type httpReq struct {
	url    string
	method string
	body   io.Reader
}

// Mock Fetch
func (hr *httpReq) Fetch(args Args) (Response, error) {
	if args["res"] == "200" {
		return http.NewRequest(hr.method, hr.url, hr.body)
	}
	return nil, errors.New("intentional error")
}

func Init() (Args, *httpReq) {
	args := Args{
		"id": "google",
	}
	hr := httpReq{
		url:    "https://google.com",
		method: "GET",
		body:   nil,
	}
	return args, &hr
}

func MainMethodforSimpleFetch() {
	args, hr := Init()
	sr := NewSRR(hr, 3)
	sr.Fetch(args)
}

func MainMethodforCrashbackLoop() {
	args, hr := Init()
	ebo := NewEBO(hr, 3, 10)
	fmt.Println("Exponential Back off")
	ebo.Fetch(args)
}

func MainMethodforCB() {
	mutex := sync.Mutex{}
	args, hr := Init()
	cb := NewCBO(hr, 5, 3, 2*time.Second, &mutex)
	fmt.Println("Circuit Breaker")
	cb.Fetch(args) //400(1)
	cb.Fetch(args) //400(2)
	cb.Fetch(args) //400(3)
	cb.Fetch(args) //400(4)
	cb.Fetch(args) //400(5) open circuit
	cb.Fetch(args) //400(1) open circuit fail fast error
	args["res"] = "200"
	cb.Fetch(args) //200(1) open circuit fail fast error
	time.Sleep(3 * time.Second)
	cb.Fetch(args) //200(1) half open circuit
	args["res"] = "400"
	cb.Fetch(args) //400(1) open circuit
	args["res"] = "200"
	time.Sleep(3 * time.Second)
	cb.Fetch(args) //200(1) half open circuit
	cb.Fetch(args) //200(2) half open circuit
	cb.Fetch(args) //200(3) half open circuit
	cb.Fetch(args) //200(1) closed circuit
}

func MultiPolicies() (Response, error) {
	mutex := sync.Mutex{}
	args, hr := Init()
	return NewCBO(NewEBO(NewSRR(hr, 3), 3, 20), 5, 3, 2*time.Second, &mutex).Fetch(args)
}

func MultiThreadedCB() {
	args, hr := Init()
	var wg sync.WaitGroup
	mutex := sync.Mutex{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			NewCBO(hr, 5, 3, 1*time.Second, &mutex).Fetch(args)
			wg.Done()
		}()
	}
	wg.Wait()
}
