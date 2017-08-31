package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// convert int to []byte
func IntToBytes(n int) []byte {
	bufReader := bytes.NewBuffer([]byte{})
	if err := binary.Write(bufReader, binary.BigEndian, int32(n)); err != nil {
		fmt.Println("IntToBytes error:", err)
		return []byte{}
	}
	return bufReader.Bytes()
}

// convert []byte to int
func BytesToInt(data []byte) int {
	bufReader := bytes.NewBuffer(data)
	var i int32
	if err := binary.Read(bufReader, binary.BigEndian, &i); err != nil {
		fmt.Println("BytesToInt error:", err)
		return -1
	}
	return int(i)
}