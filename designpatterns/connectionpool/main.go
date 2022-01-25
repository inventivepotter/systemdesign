package main

import (
	"fmt"
	"sync"
	"time"

	cp "github.com/inventivepotter/systemdesign/designpatterns/connectionpool"
)

func main() {
	//waitGroup for waiting until all threads/goroutines are finished
	var wg sync.WaitGroup
	//Mutex is to perform Lock/Unlock
	var m sync.Mutex
	//Once is to make sure the program runs only once
	var once sync.Once
	for i := 0; i < 6; i++ {
		wg.Add(1)
		//Spinning off 6 threads/goroutines which will try to
		//1.Create New Pool
		//2.Acquire a connection
		//3.Execute command
		//4.Sleeps to mock execution time
		//5.Release connection
		go func(i int, mutex *sync.Mutex, once *sync.Once, wg *sync.WaitGroup) {

			//1.Pool Creation, given it's singleton it will create only once and re-use the same object for others
			//Pool with 5 connections
			pool, err := cp.Setup(cp.PoolConfig{
				MaxConnections: 5,
			}, mutex, once)
			if err != nil {
				fmt.Println(err)
			}
			//2.Acquire Connection
			//There are only 5 connections available and there 6 threads asking for connection, hence one thread will see the error
			//We can improve this by waiting instead of erroring or retrying on connection error when there is no connection available.
			c, err := pool.AcquireConnection()
			if err != nil {
				fmt.Println(err)
			} else {
				con := *c
				//3.Executing cmd from a successful connection
				con.Execute("SELECT * FROM emp;")
				//4.sleeping to hold the connection, otherwise release will make the connection available back to the pool instantly
				//only for demonstration
				time.Sleep(1 * time.Second)
				//5.Releasing connection
				pool.ReleaseConnection(con)
			}
			//marking the routine's task to be done
			wg.Done()
		}(i, &m, &once, &wg)
	}
	//waiting for all routines to be finished
	wg.Wait()
	fmt.Println("\nafter all the threads are done and released the connections")
	//Creating another Pool object (will basically get the same obj) to print the connection details after execution
	pool, err := cp.Setup(cp.PoolConfig{
		MaxConnections: 5,
	}, &m, &once)
	if err != nil {
		fmt.Println(err)
	}
	pool.Details()
}
