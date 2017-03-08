package main

import (
	"net"
	"io"
)

func main()  {
	unixAddr, err := net.ResolveUnixAddr("unix", "/tmp/echo.sock")
	checkErr(err)

	l, err := net.ListenUnix("unix", unixAddr)
	checkErr(err)

	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}

		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
