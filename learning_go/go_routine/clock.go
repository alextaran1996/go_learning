package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	var i int
	fmt.Printf("Enter port number:")
	fmt.Scan(&i)
	address := fmt.Sprintf("localhost:%d", i)
	fmt.Println("Connect to address", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print(err)
			continue
		}
		go handleConn(conn)
	}

}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
