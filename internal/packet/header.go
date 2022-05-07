package packet

import (
	"github.com/maddiesch/telemetry/internal/scanner"
)

// Header
type Header struct {
	PacketFormat         uint16
	GameMajorVersion     uint8
	GameMinorVersion     uint8
	PacketVersion        uint8
	PacketID             uint8
	SessionID            uint64
	SessionTime          float32
	FrameID              uint32
	PlayerCarIndex       uint8
	SecondPlayerCarIndex uint8
}

func decodeHeader(s *scanner.Scanner) (Header, error) {
	h := Header{}

	err := scannerReadAll(s,
		&h.PacketFormat,
		&h.GameMajorVersion,
		&h.GameMinorVersion,
		&h.PacketVersion,
		&h.PacketID,
		&h.SessionID,
		&h.SessionTime,
		&h.FrameID,
		&h.PlayerCarIndex,
		&h.SecondPlayerCarIndex,
	)
	if err != nil {
		return Header{}, err
	}

	return h, nil
}
