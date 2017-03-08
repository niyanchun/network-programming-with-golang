package main

import (
	"net"
	"time"
)

func main() {
	service := ":2000"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkErr(err)

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte

	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	daytime := time.Now().String()

	conn.WriteToUDP([]byte(daytime), addr)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
