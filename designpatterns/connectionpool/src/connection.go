package connectionpool

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"

	"github.com/google/uuid"
)

type Response interface{}

// Connection is an interface used in Pool, so that we can plugin any type of connection wrapper
type Connection interface {
	Connect() error
	Release() error
	Execute(statement string) (Response, error)
	IsAvailable() bool
	Assign()
	Unassign()
	ID() string
	Details()
}

// connection is going to be the wrapper for original connection
// connection is unexported because external parties shouldn't messup with con
// connection implements Connection
type connection struct {
	id        uuid.UUID
	available bool
	//wrapping of original connection
	conn  net.Conn
	mutex *sync.Mutex
}

//Abstarct Method to call original implementation of Connect according to the connection
//Ex: Wrapper method for DB.Connect()
func (c *connection) Connect() error {
	fmt.Println("\t" + c.ID() + " is connected")
	return nil
}

//Abstarct Method to call original implementation of Release according to the connection
//Ex: Wrapper method for DB.Release()
func (c *connection) Release() error {
	fmt.Println("\t" + c.ID() + " is disconnected")
	return nil
}

//Abstarct Method to call original implementation of Execute according to the connection
//Ex: Wrapper method for DB.Execute()
func (c *connection) Execute(statement string) (Response, error) {
	fmt.Println("\t" + "statement executed : " + statement + " on connection " + c.ID())
	return map[string]interface{}{"result": statement}, nil
}

func (c *connection) IsAvailable() bool {
	return c.available
}

func (c *connection) Assign() {
	c.available = false
}

func (c *connection) ID() string {
	str, _ := json.MarshalIndent(c.id, "", "")
	return string(str)
}

//Unassign sets the connection to be available again
func (c *connection) Unassign() {
	c.available = true
}

//Details prints the details of connection
func (c *connection) Details() {
	fmt.Print("\t{\"id\":", c.ID())
	fmt.Print(", \"available\":", c.IsAvailable(), "}")
	fmt.Println()
}

//CreatConnection returns a new connection
func CreateConnection(mutex *sync.Mutex) (Connection, error) {
	id := uuid.New()
	fmt.Println("\t"+"new connection created with id ", id)
	return &connection{
		id:        id,
		available: true,
		conn:      nil,
		mutex:     mutex,
	}, nil
}
