package telemetry

import (
	"encoding/binary"
	"errors"
	"math"
)

var ErrInvalidPacketSize = errors.New("invalid packet size")

func Parse(data []byte) (Packet, error) {
	if len(data) < 40 {
		return Packet{}, ErrInvalidPacketSize
	}
	var p Packet
	p.ID = binary.LittleEndian.Uint32(data[0:4])
	p.Timestamp = int64(binary.LittleEndian.Uint64(data[4:12]))
	p.Lat = math.Float64frombits(binary.LittleEndian.Uint64(data[12:20]))
	p.Lon = math.Float64frombits(binary.LittleEndian.Uint64(data[20:28]))
	p.Alt = math.Float64frombits(binary.LittleEndian.Uint64(data[28:36]))
	p.Battery = data[36]
	return p, nil
}
