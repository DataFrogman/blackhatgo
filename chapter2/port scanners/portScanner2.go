package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		connection, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		connection.Close()
		fmt.Printf("%d open\n", i)
	}
}
