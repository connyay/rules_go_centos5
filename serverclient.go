package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

func startLocalServer() net.Addr {
	l, err := net.Listen("tcp", "localhost:0")
	checkErr(err)

	go func() {

		for {
			c, err := l.Accept()
			checkErr(err)
			netData, err := bufio.NewReader(c).ReadString('\n')
			if errors.Is(err, io.EOF) {
				return
			}
			checkErr(err)

			fmt.Print("-> ", string(netData))
			t := time.Now()
			myTime := t.Format(time.RFC3339) + "\n"
			c.Write([]byte(myTime))
		}
	}()

	return l.Addr()
}

func readWriteClient(addr net.Addr) {
	c, err := net.Dial("tcp", addr.String())
	checkErr(err)

	_, err = c.Write([]byte("testing\n"))
	checkErr(err)

	rdr := bufio.NewReader(c)
	l, _, err := rdr.ReadLine()
	checkErr(err)
	fmt.Printf("read %q\n", l)
	c.Close()
}
