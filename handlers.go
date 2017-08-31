package main

import (
	"net"
	"fmt"
)

func handleConn(conn net.Conn) {
	fmt.Println("--------------------")
	fmt.Println("handling client:", conn.RemoteAddr())
	defer conn.Close()
	fmt.Println(conn.RemoteAddr())

	words := "test string"
	data := []byte(words)
	lengthBytes := IntToBytes(len(data))
	fmt.Println("sending data:", words)
	conn.Write(lengthBytes)
	conn.Write(data)

	fmt.Println("reading data...")
	recvLengthBytes := make([]byte, 4)
	if _, err := conn.Read(recvLengthBytes); err != nil {
		fmt.Println("cannot read data length:", err)
		return
	}
	dataBytes := make([]byte, BytesToInt(recvLengthBytes))
	if _, err := conn.Read(dataBytes); err != nil {
		fmt.Println("cannot read data:", err)
		return
	}
	fmt.Println("received data:", string(dataBytes))
	fmt.Println("connection handled")
	fmt.Println("--------------------")
}
