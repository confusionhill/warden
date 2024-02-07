package loop

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"server/network/pool"
	"sync"
)

func MainLoop(net net.Listener) {
	var mutex sync.Mutex
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("press 1 to check connected clients \n press 2 to shutdown server")
	for scanner.Scan() {
		message := scanner.Text()
		if message == "1" {
			mutex.Lock()
			fmt.Println("connected clients", pool.GetConnectedClients())
			mutex.Unlock()
		}
		if message == "2" {
			net.Close()
			fmt.Println("server shutting down")
			return
		}
		fmt.Println("cli :", message)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading standard input:", err)
		return
	}
}
