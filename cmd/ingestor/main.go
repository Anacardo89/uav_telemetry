package main

import (
	"fmt"
	"net"

	"github.com/Anacardo89/uav_telemetry/internal/telemetry"
)

func main() {
	conn, _ := net.ListenUDP("udp", &net.UDPAddr{Port: 8080})
	packetChan := make(chan []byte, 10000)

	for range 100 {
		go worker(packetChan)
	}

	for {
		buf := make([]byte, 40)
		n, _, _ := conn.ReadFromUDP(buf)
		packetChan <- buf[:n]
	}
}

func worker(ch chan []byte) {
	for raw := range ch {
		p, err := telemetry.Parse(raw)
		if err != nil {
			fmt.Println("failed to parse packet")
		}
		fmt.Printf("Drone %d at Alt: %.2f\n", p.ID, p.Alt)
	}
}
