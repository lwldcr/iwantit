package main

import (
	"net"
	"fmt"
	"strconv"
)

func StartClient() {
	fmt.Println("--------------")
	fmt.Println("starting client...")
	addr := conf.HostIp + ":" + strconv.Itoa(conf.ServerPort)
	fmt.Println("resolving server address:", addr)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		fmt.Println("failed:", err)
		return
	}

	fmt.Println("dialing server", tcpAddr)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Cannot connect server: ", err)
		return
	}
	defer conn.Close()

	fmt.Println("connected to server: ", conn.RemoteAddr(), "local port", conn.LocalAddr())

	fmt.Println("---------")
	fmt.Println("trying to read data length")
	d := make([]byte, 4)
	conn.Read(d)

	length := BytesToInt(d)
	fmt.Println("got length info:", length)
	fmt.Println("reading data...")
	data := make([]byte, length)
	conn.Read(data)
	fmt.Println("got:", string(data))

	fmt.Println("test sending data:...")
	clientStr := "Hello server, please serve for me!"
	lengthBytes := IntToBytes(len(clientStr))
	fmt.Println("first sending data length...")
	if _, err := conn.Write(lengthBytes); err != nil {
		fmt.Println("send data length failed:", err)
		return
	}
	fmt.Println("sent data length:", len(clientStr))
	fmt.Println("sending data...")
	if _, err := conn.Write([]byte(clientStr)); err != nil {
		fmt.Println("failed:", err)
		return
	}
	fmt.Println("done")
}