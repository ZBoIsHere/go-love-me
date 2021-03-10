package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	const (
		INTERFACE_NAMT = "eth0"
		CLOUD_PUBLIC_IP = "119.8.112.94"
		CLOUD_PUBLIC_PORT int = 9981
		EDGE_PRIVATE_PORT int = 0000
	)

	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: CLOUD_PUBLIC_PORT})
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("Local Address: <%s> \n", listener.LocalAddr().String())
	peers := make([]net.UDPAddr, 0, 2)
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		log.Printf("<%s> %s\n", remoteAddr.String(), data[:n])
		peers = append(peers, *remoteAddr)
		if len(peers) == 2 {

			log.Printf("Begin udp hole punch %s <--> %s conn\n", peers[0].String(), peers[1].String())
			listener.WriteToUDP([]byte(peers[1].String()), &peers[0])
			listener.WriteToUDP([]byte(peers[0].String()), &peers[1])
			time.Sleep(time.Second * 8)
			log.Println("relay quit, but not influence p2p conn")
			return
		}
	}
}
