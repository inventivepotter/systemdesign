package connectionpool

import (
	"errors"
	"fmt"
	"sync"
)

//PoolConfig is a config struct used to create a pool
type PoolConfig struct {
	MaxConnections int
	mutex          *sync.Mutex
}

//pool is the struct that has only one objecte and holds all details of the pool
//equivalent to pool class
type pool struct {
	maxConnections int
	connections    map[string]*Connection
	mutex          *sync.Mutex
}

//Pool variable for storing the pool struct object
//Setup method will configure the Pool assign it to this var
var Pool *pool

//AcquireConnection returns a connection from the pool which is available and marks it's busy
//AcquireConnection returns Connection interface so that it's abstracted from implementation of type of connection
func (p *pool) AcquireConnection() (*Connection, error) {
	p.mutex.Lock()
	for conid, conval := range p.connections {
		con := *conval
		if con.IsAvailable() {
			fmt.Println("acquired connection with Id " + conid)
			con.Assign()
			conval = &con
			defer p.mutex.Unlock()
			return conval, nil
		}
	}
	defer p.mutex.Unlock()
	return nil, errors.New("error: no connections available")
}

//ReleaseConnection marks the particular connection to be available for the pool
func (p *pool) ReleaseConnection(c Connection) error {
	p.mutex.Lock()
	c.Unassign()
	defer p.mutex.Unlock()
	return nil
}

//Details returns the Pool Details
func (p *pool) Details() {
	fmt.Println("Max Connections", p.maxConnections)
	fmt.Println("Connections")
	for _, con := range p.connections {
		c := *con
		c.Details()
	}
}

func makeConnections(maxCon int, mutex *sync.Mutex) map[string]*Connection {
	cons := map[string]*Connection{}
	for i := 0; i < maxCon; i++ {
		con, _ := CreateConnection(mutex)
		con.Connect()
		cons[con.ID()] = &con
	}
	return cons
}

//Setup is the static function that creats Pool if not already created and assign it to Pool variable
func Setup(pc PoolConfig, mutex *sync.Mutex, once *sync.Once) (*pool, error) {
	// Signleton implementation to create Pool object only once
	if Pool == nil {
		mutex.Lock()
		if Pool == nil {
			Pool = &pool{
				maxConnections: pc.MaxConnections,
				connections:    makeConnections(pc.MaxConnections, mutex),
				mutex:          mutex,
			}
			fmt.Println("\nNew Pool successfully created with following details")
			Pool.Details()
		}
		mutex.Unlock()
	}
	return Pool, nil
}
