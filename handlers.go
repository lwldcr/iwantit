package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	//	"time"
)

func handleConn(conn net.Conn) {
	fmt.Println("--------------------")
	fmt.Println("handling client:", conn.RemoteAddr())
	defer conn.Close()
	fmt.Println(conn.RemoteAddr())

	fmt.Println("reading file header...")
	recvLengthBytes := make([]byte, 4)
	if _, err := conn.Read(recvLengthBytes); err != nil {
		fmt.Println("cannot read header length:", err)
		return
	}

	headerBytes := make([]byte, BytesToInt(recvLengthBytes))
	if _, err := conn.Read(headerBytes); err != nil {
		fmt.Println("cannot read header content:", err)
		return
	}

	var header Header
	if err := json.Unmarshal(headerBytes, &header); err != nil {
		fmt.Println("cannot decode header content:", err)
		return
	}
	fmt.Println("received file header:", header)

	saveFile := "new_" + header.Filename

	fp, err := os.OpenFile(saveFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed:", err)
		return
	}
	defer fp.Close()

	offset := int64(0)
	//	var batchLength int
	for {
		//		batchLengthBytes := make([]byte, 4)
		//		n, err := conn.Read(batchLengthBytes)
		//		if err != nil {
		//			fmt.Println("reading batch data length failed:", err)
		//			break
		//		}
		//		if n != 4 {
		//			fmt.Println("reading batch data length error:", "expect: 4", "got:", n)
		//			break
		//			continue
		//		}
		//		batchLength = BytesToInt(batchLengthBytes)
		//		if batchLength <= 0 || batchLength > int(conf.BatchSize) {
		//			fmt.Println("Got wrong batch length:", batchLength)
		//			time.Sleep(1 * time.Microsecond)
		//			break
		//		}
		//		fmt.Println("Got batch length:", batchLength)
		dataBytes := make([]byte, conf.BatchSize)
		n, err := conn.Read(dataBytes)
		if err != nil {
			fmt.Println("reading batch data failed:", err)
			break
		}

		fmt.Println("writing batch data into new file:", n)
		wn, err := fp.Write(dataBytes[:n])
		if err != nil {
			fmt.Println("write failed:", err)
			return
		}

		offset += int64(wn)
		if offset >= header.TotalSize {
			fmt.Printf("detected file length: %d, should be: %d\n", offset, header.TotalSize)
			break
		}
	}

	fmt.Println("new file saved")
	fmt.Println("connection handled")
	fmt.Println("--------------------")
}
