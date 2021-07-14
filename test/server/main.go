package main

import (
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp4", "0.0.0.0:5555")
	log.Println(addr)
	if err != nil {
		log.Panic(err)
	}
	socket, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Panic(err)
	}

	defer socket.Close()
	for {
		data := make([]byte, 4096)
		read, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			log.Panic(err)
		}
		log.Println("S-IN ] " + string(data[:read]) + " from " + remoteAddr.String())
		socket.Write([]byte("hi"))
	}

}
