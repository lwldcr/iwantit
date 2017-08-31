package main

import (
	"flag"
	"fmt"
	"os"
)

var conf Conf
func init() {
	conf = Conf{
		ServerPort:8098, ClientPort:8099,
		HostIp:"127.0.0.1",}
}

func Usage() {
	fmt.Println("welcome to IWantYou.")
	fmt.Println("Usage: ")
	fmt.Println("\t ./IwantYou ${ROLE}")
	fmt.Println("ROLE: server|client")
	os.Exit(1)
}

func main() {
	flag.Parse()

	if len(os.Args) != 2 {
		Usage()
	}
	role := os.Args[1]
	switch role {
	case "server":
		conf.Role = Host
	case "client":
		conf.Role = Client
	default:
		Usage()
	}

	fmt.Println("------------------------")
	switch conf.Role {
	case Host:
		StartServer()
	case Client:
		StartClient()
	}

}

func StartServer() {
	s := GetServer()
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println("accept client failed: ", err)
			continue
		}
		fmt.Println("handing new connect...")
		go handleConn(c)
	}
}
