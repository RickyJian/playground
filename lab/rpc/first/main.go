package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(req string, resp *string) error {
	*resp = fmt.Sprintf("hello: %s", req)
	return nil
}

func main() {

	go func() {
		server()
	}()

	done := make(chan struct{})
	go func() {
		client()
		done <- struct{}{}
	}()
	<-done
}

func server() {
	if err := rpc.RegisterName("HelloService", new(HelloService)); err != nil {
		log.Fatalf("[SERVER] register name failed: %v", err)
	}
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("[SERVER] listen error: %v", err)
	}

	conn, err := listen.Accept()
	if err != nil {
		log.Fatalf("[SERVER] accept error: %v", err)
	}
	rpc.ServeConn(conn)
}

func client() {
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("[CLIENT] dial server failed: %v", err)
	}
	defer client.Close()

	var reply string
	if err := client.Call("HelloService.Hello", "world!!", &reply); err != nil {
		log.Fatalf("[CLIENT] call server failed: %v", err)
	}
	fmt.Println(reply)
}
