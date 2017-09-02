package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func handleClientConn(conn net.Conn, path string) {
	fmt.Println("start to sending data...", path)
	header := Header{Uuid: GetUUID(), Filename: path}

	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("stat file failed:", err)
		return
	}
	header.TotalSize = fileInfo.Size()

	headerBytes, _ := json.Marshal(header)
	headerLength := len(headerBytes)

	fmt.Println("sending file header:", headerLength, header)
	if _, err := conn.Write(IntToBytes(headerLength)); err != nil {
		fmt.Println("send header length failed:", err)
		return
	}
	if _, err := conn.Write(headerBytes); err != nil {
		fmt.Println("send header content failed:", err)
		return
	}

	fmt.Println("file header sent successfully")

	fp, err := os.Open(path)
	if err != nil {
		fmt.Println("opening file failed:", err)
		return
	}

	start := time.Now()
	var speed int64

	offset := int64(0)
	for {
		dataBytes := make([]byte, conf.BatchSize)

		n, err := fp.Read(dataBytes)
		if err != nil {
			fmt.Println("reading file failed:", err)
			break
		}

		fmt.Println("sending batch data...")
		//fmt.Println("batch data length:", n)
		//if _, err := conn.Write(IntToBytes(n)); err != nil {
		//	fmt.Println("send batch length failed:", err)
		//	break
		//}

		wn, err := conn.Write(dataBytes[:n])
		if err != nil {
			fmt.Println("sending data failed:", err)
			return
		}
		fmt.Println("sent data:", wn)

		offset += int64(n)
		end := time.Now()
		delta := end.Sub(start)
		elapsed := delta.Hours() * 3600 + delta.Minutes() * 60 + delta.Seconds()
		if elapsed > 0 {
			speed = offset / int64(elapsed)
		}
		fmt.Printf("sending process: %d / %d, speed: %d Bytes/Second\n",
		offset, header.TotalSize, speed)

		if offset >= header.TotalSize {
			fmt.Printf("file read: %d, total: %d\n", offset, header.TotalSize)
			break
		}
	}
	fmt.Printf("%s sent successfully\n", path)
}

func StartClient() {
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

	for _, path := range conf.Paths {
		fmt.Println("handling file:", path)
		handleClientConn(conn, path)
	}
}
