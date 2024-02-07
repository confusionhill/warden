package handler

import (
	"bufio"
	"fmt"
	"net"
	"server/modules/handler/broadcaster"
	"server/network/pool"
	"sync"
)

type Handlers struct {
	conn   net.Conn
	mutex  *sync.Mutex
	reader *bufio.Reader
}

func New(conn net.Conn, mutex *sync.Mutex) Handlers {
	return Handlers{
		conn:   conn,
		mutex:  mutex,
		reader: bufio.NewReader(conn),
	}
}

func (h *Handlers) MainConnHandler() {
	defer func() {
		h.conn.Close()
		fmt.Println("conn closed")
	}()
	h.mutex.Lock()
	pool.AddClient(h.conn)
	h.mutex.Unlock()

	for {
		message, err := h.reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err.Error())
			break
		}
		broadcaster.BroadcastMessage(h.conn, h.mutex, message)
	}

	h.mutex.Lock()
	pool.DisconnectClient(h.conn)
	h.mutex.Unlock()
}
