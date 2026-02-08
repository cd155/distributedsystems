package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"
)

// 1. Define Arguments and Reply structs
type PutArgs struct {
	Key, Value string
}
type PutReply struct{} // Empty

type GetArgs struct {
	Key string
}
type GetReply struct {
	Value string
	Found bool
}

// 2. Define the Server Object
type KVServer struct {
	mu   sync.Mutex
	data map[string]string
}

// 3. Implement the RPC Methods
func (s *KVServer) Put(args *PutArgs, reply *PutReply) error {
	// TODO: Lock, set value, unlock
	return nil
}

func (s *KVServer) Get(args *GetArgs, reply *GetReply) error {
	// TODO: Lock, get value, unlock
	return nil
}

// 4. Main logic to start server and run client
func main() {
	// --- Server Start ---
	kv := new(KVServer)
	kv.data = make(map[string]string)
	rpc.Register(kv)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil) // Start server in background

	// --- Client Logic ---
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// TODO: Make a synchronous call to Put("name", "Sentinel")
	// TODO: Make a synchronous call to Get("name") and print result
	fmt.Println("Client finished!")
}
