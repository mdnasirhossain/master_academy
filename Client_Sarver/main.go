package main

import (
	"fmt"
	"net"
)

func main() {

	var conn net.Conn
	var err error
	conn, err = net.Dial("tcp", "91.205.173.170:8888")
	if err != nil {
		fmt.Println(err.Error())
	}

	//way-1
	// bs := []byte("Hello")
	// conn.Write(bs)

	//way-2
	conn.Write([]byte("Sir Amader golang Book diban na"))

	bs := make([]byte, 1024)
	n, _ := conn.Read(bs)
	fmt.Println(string(bs[:n]))

	conn.Close()

}
