package main

import (
	"fmt"
	"net"
	"strconv"
)

func GetServer() net.Listener {
	addr := conf.HostIp + ":" + strconv.Itoa(conf.ServerPort)
	fmt.Println("starting server on", addr)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error building server: ", err)
		return nil
	}
	fmt.Println("server started")
	return l
}
