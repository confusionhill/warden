package main

import (
	"fmt"
	"net"
	"server/modules/handler"
	"server/pool"
	"sync"
)

func main() {
	const (
		NET_TYPE = "tcp"
		PORT     = ":8081"
	)
	network, err := net.Listen(NET_TYPE, PORT)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}
	defer network.Close()
	var mutex sync.Mutex
	fmt.Printf("listening on %s \n", PORT)
	for {
		conn, err := network.Accept()
		if err != nil {
			fmt.Errorf("%s", err.Error())
			continue
		}
		mutex.Lock()
		if pool.IsPoolMaxed() {
			mutex.Unlock()
			conn.Close()
			continue
		}
		mutex.Unlock()
		go handler.MainConnHandler(conn, &mutex)
	}
}
