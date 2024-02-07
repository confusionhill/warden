package handler

import (
	"bufio"
	"fmt"
	"net"
	"server/pool"
	"sync"
)

func MainConnHandler(conn net.Conn, mutex *sync.Mutex) {
	defer func() {
		conn.Close()
		fmt.Println("conn closed")
	}()

	mutex.Lock()
	pool.AddClient(conn)
	mutex.Unlock()
	fmt.Println("new client added to client pool")
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err.Error())
			break
		}
		broadcastMessage(conn, message)
	}
	mutex.Lock()
	pool.DisconnectClient(conn)
	mutex.Unlock()
}

func broadcastMessage(sender net.Conn, message string) {
	fmt.Println("msg: ", message)
	clients := pool.GetClients()
	for conn := range clients {
		if conn != sender {
			_, err := conn.Write([]byte(message))
			if err != nil {
				fmt.Println("Error broadcasting message to client:", err.Error())
			}
		}
	}
}
