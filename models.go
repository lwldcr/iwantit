package main

type Conf struct {
	ServerPort int // server port
	ClientPort int // client port
	HostIp	string // host ip address
	Role	int // role definition
}

const (
	Host = iota
	Client
)
