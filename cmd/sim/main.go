package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"net"
	"time"

	"github.com/Anacardo89/uav_telemetry/internal/telemetry"
)

func main() {
	drones := flag.Int("drones", 1000, "Number of concurrent UAVs to simulate")
	freq := flag.Int("hz", 50, "Update frequency per drone in Hz")
	flag.Parse()
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	fmt.Printf("🚀 Starting swarm: %d drones at %dHz...\n", *drones, *freq)

	for i := 1; i <= *drones; i++ {
		go startDroneSim(uint32(i), addr, *freq)
	}
	select {}
}

func startDroneSim(id uint32, addr *net.UDPAddr, hz int) {
	conn, _ := net.DialUDP("udp", nil, addr)
	defer conn.Close()
	ticker := time.NewTicker(time.Second / time.Duration(hz))
	defer ticker.Stop()

	p := telemetry.NewPacket(id, 0, 0, 0)
	buf := make([]byte, 40)

	for range ticker.C {
		p.Timestamp = time.Now().UnixNano()
		p.Lat += (rand.Float64() - 0.5) * 0.001
		p.Lon += (rand.Float64() - 0.5) * 0.001
		p.Alt += (rand.Float64() - 0.5) * 0.1

		binary.LittleEndian.PutUint32(buf[0:4], p.ID)
		binary.LittleEndian.PutUint64(buf[4:12], uint64(p.Timestamp))
		binary.LittleEndian.PutUint64(buf[12:20], math.Float64bits(p.Lat))
		binary.LittleEndian.PutUint64(buf[20:28], math.Float64bits(p.Lon))
		binary.LittleEndian.PutUint64(buf[28:36], math.Float64bits(p.Alt))
		buf[36] = p.Battery

		_, err := conn.Write(buf)
		if err != nil {
			fmt.Printf("Drone %d connection lost: %s\n", id, err)
			return
		}
	}
}
