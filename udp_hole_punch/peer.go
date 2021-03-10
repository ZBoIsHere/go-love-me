package main

import (
	"fmt"
	"go-love-me/utils"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var tag string

const HAND_SHAKE_MSG = "UDP HOLE PUNCH MESSAGE"

func main() {

	const (
		INTERFACE_NAMT = "eth0"
		CLOUD_PUBLIC_IP = "119.8.112.94"
		CLOUD_PUBLIC_PORT int = 9981
		EDGE_PRIVATE_PORT int = 0000
	)

	// 当前进程标记字符串,便于显示
	tag, _ := os.Hostname()
	srcAddr := &net.UDPAddr{IP: utils.GetInterfaceIP(INTERFACE_NAMT), Port: EDGE_PRIVATE_PORT} // 注意端口必须固定
	dstAddr := &net.UDPAddr{IP: net.ParseIP(CLOUD_PUBLIC_IP), Port: CLOUD_PUBLIC_PORT}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	if _, err = conn.Write([]byte("hello, I'm new peer:" + tag)); err != nil {
		log.Panic(err)
	}
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Printf("error during read: %s", err)
	}
	conn.Close()
	anotherPeer := parseAddr(string(data[:n]))
	fmt.Printf("local:%s server:%s another:%s\n", srcAddr, remoteAddr, anotherPeer.String())

	// 开始打洞
	bidirectionHole(srcAddr, &anotherPeer)

}

func parseAddr(addr string) net.UDPAddr {
	t := strings.Split(addr, ":")
	port, _ := strconv.Atoi(t[1])
	return net.UDPAddr{
		IP:   net.ParseIP(t[0]),
		Port: port,
	}
}

func bidirectionHole(srcAddr *net.UDPAddr, anotherAddr *net.UDPAddr) {
	conn, err := net.DialUDP("udp", srcAddr, anotherAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	// 向另一个peer发送一条udp消息(对方peer的nat设备会丢弃该消息,非法来源),用意是在自身的nat设备打开一条可进入的通道,这样对方peer就可以发过来udp消息
	if _, err = conn.Write([]byte(HAND_SHAKE_MSG)); err != nil {
		log.Println("send handshake:", err)
	}
	go func() {
		for {
			time.Sleep(10 * time.Second)
			if _, err = conn.Write([]byte("from [" + tag + "]")); err != nil {
				log.Println("send msg fail", err)
			}
		}
	}()
	for {
		data := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Printf("error during read: %s\n", err)
		} else {
			log.Printf("Receive date:%s\n", data[:n])
		}
	}
}

