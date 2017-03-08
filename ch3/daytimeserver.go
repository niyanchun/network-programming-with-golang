package main

import (
	"net"
	"time"
)

func main()  {
	service := ":2000"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkErr(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
