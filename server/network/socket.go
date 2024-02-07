package network

import (
	"fmt"
	"net"
	"server/modules/handler"
	"server/network/pool"
	"sync"
)

func StartServer(network net.Listener) {
	var mutex sync.Mutex
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
		hdl := handler.New(conn, &mutex)
		go hdl.MainConnHandler()
	}
}
