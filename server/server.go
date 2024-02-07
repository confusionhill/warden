package main

import (
	"fmt"
	"net"
	"server/loop"
	"server/network"
)

func main() {
	const (
		NET_TYPE = "tcp"
		PORT     = ":8081"
	)
	net, err := net.Listen(NET_TYPE, PORT)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}
	fmt.Printf("listening on %s \n", PORT)
	defer net.Close()
	go network.StartServer(net)

	loop.MainLoop(net)
}
