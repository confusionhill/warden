package broadcaster

import (
	"fmt"
	"net"
	"server/network/pool"
	"sync"
)

func BroadcastMessage(sender net.Conn, mutex *sync.Mutex, message string) {
	mutex.Lock()
	clients := pool.GetClients()
	mutex.Unlock()
	for conn := range clients {
		if conn != sender {
			_, err := conn.Write([]byte(message))
			if err != nil {
				fmt.Println("Error broadcasting message to client:", err.Error())
			}
		}
	}
}
