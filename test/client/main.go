package main

import (
	"log"
	"net"
)

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp4", "255.255.255.255:5555")
	log.Println(ServerAddr)
	if err != nil {
		log.Panic(err)
	}

	socket, err := net.DialUDP("udp4", nil, ServerAddr)
	if err != nil {
		log.Panic(err)
	}
	go func() {
		for {
			data := make([]byte, 4096)
			read, remoteAddr, err := socket.ReadFromUDP(data)
			if err != nil {
				log.Panic(err)
			}
			log.Println("S-IN ] " + string(data[:read]) + " from " + remoteAddr.String())
		}
	}()

	defer socket.Close()
	data := "hello"
	socket.Write([]byte(data))
	log.Println("C-OUT " + data)
}
