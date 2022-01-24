package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/inventivepotter/systemdesign/designpatterns/decorator"
)

func main() {
	//MainMethodforSimpleFetch()
	//MainMethodforCrashbackLoop()
	//MainMethodforCB()
	//ComboExpoCB()
	//MultiPolicies()
	MultiThreadedCB()
}

type httpReq struct {
	url    string
	method string
	body   io.Reader
}

// Mock Fetch
func (hr *httpReq) Fetch(args decorator.Args) (decorator.Response, error) {
	if args["res"] == "200" {
		return http.NewRequest(hr.method, hr.url, hr.body)
	}
	return nil, errors.New("intentional error")
}

func Init() (decorator.Args, *httpReq) {
	args := decorator.Args{
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
	sr := decorator.NewSRR(hr, 3)
	sr.Fetch(args)
}

func MainMethodforCrashbackLoop() {
	args, hr := Init()
	ebo := decorator.NewEBO(hr, 3, 10)
	fmt.Println("Exponential Back off")
	ebo.Fetch(args)
}

func MainMethodforCB() {
	mutex := sync.Mutex{}
	args, hr := Init()
	cb := decorator.NewCBO(hr, 5, 3, 2*time.Second, &mutex)
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

func MultiPolicies() (decorator.Response, error) {
	mutex := sync.Mutex{}
	args, hr := Init()
	return decorator.NewCBO(decorator.NewEBO(decorator.NewSRR(hr, 3), 3, 20), 5, 3, 2*time.Second, &mutex).Fetch(args)
}

func MultiThreadedCB() {
	args, hr := Init()
	var wg sync.WaitGroup
	mutex := sync.Mutex{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			decorator.NewCBO(hr, 5, 3, 1*time.Second, &mutex).Fetch(args)
			wg.Done()
		}()
	}
	wg.Wait()
}
