package main

import (
	"flag"
	"fmt"
	"os"
)

var conf Conf

func init() {
	conf = Conf{
		ServerPort: 8098, ClientPort: 8099,
		HostIp:    "127.0.0.1",
		BatchSize: 1024 * 1024}
}

func Usage() {
	fmt.Println("welcome to IWantYou.")
	fmt.Println("Usage: ")
	fmt.Println("\t ./IwantYou ${ACTION} [filepath]")
	fmt.Println("ACTION: startserver|send")
	fmt.Println("filepath: path to file you want to send, required as client")
	os.Exit(1)
}

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		Usage()
	}
	role := os.Args[1]
	switch role {
	case "startserver":
		conf.Role = Host
	case "send":
		if len(os.Args) != 3 {
			Usage()
		}
		conf.Role = Client
		paths := make([]string, 0)
		path := os.Args[2]
		paths = append(paths, path)
		conf.Paths = paths
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
