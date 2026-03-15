package telemetry

import "time"

type Packet struct {
	ID        uint32
	Timestamp int64
	Lat       float64
	Lon       float64
	Alt       float64
	Battery   uint8
	Padding   [3]byte
}

func NewPacket(id uint32, lat, lon, alt float64) Packet {
	return Packet{
		ID:        id,
		Timestamp: time.Now().UnixNano(),
		Lat:       lat,
		Lon:       lon,
		Alt:       alt,
		Battery:   100,
	}
}
