package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	// Connect to server
	conn, err := net.Dial("tcp", "127.0.0.1:65432")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send an integer and a string to server
	a := int32(10)
	s := "Badla"
	data := make([]byte, 4+len(s))
	binary.LittleEndian.PutUint32(data[:4], uint32(a))
	copy(data[4:], s)
	_, err = conn.Write(data)
	if err != nil {
		panic(err)
	}

	// Receive an array of strings from server
	var n int32
	err = binary.Read(conn, binary.LittleEndian, &n)
	if err != nil {
		panic(err)
	}
	strings := make([]string, n)
	for i := int32(0); i < n; i++ {
		var strLen int32
		err = binary.Read(conn, binary.LittleEndian, &strLen)
		if err != nil {
			panic(err)
		}
		strData := make([]byte, strLen)
		_, err = conn.Read(strData)
		if err != nil {
			panic(err)
		}
		strings[i] = string(strData)
	}

	// Print result
	fmt.Printf("Received %d strings from server: %v\n", n, strings)
}
