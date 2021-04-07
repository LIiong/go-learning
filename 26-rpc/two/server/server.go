package main

import (
	"log"
	"net"
	"net/rpc"
	"go-learning/26-rpc/two/inter"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
    *reply = "helloworld:" + request
    return nil
}

func main() {
    inter.RegisterHelloService(new(HelloService))

    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("ListenTCP error:", err)
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal("Accept error:", err)
        }

        go rpc.ServeConn(conn)
    }
}