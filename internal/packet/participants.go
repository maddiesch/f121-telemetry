package packet

import (
	"github.com/maddiesch/telemetry/internal/scanner"
)

type Participants struct {
	header Header

	ActiveCarCount uint8

	Data [22]ParticipantData
}

type ParticipantData struct {
	AIControlled  uint8
	DriverID      uint8
	NetworkID     uint8
	TeamID        uint8
	MyTeam        uint8
	RaceNumber    uint8
	Nationality   uint8
	Name          [48]byte
	YourTelemetry uint8
}

func (m *Participants) Header() Header { return m.header }

func (m *Participants) ID() ID { return PacketTypeParticipants }

var _ Packet = (*Participants)(nil)

func decodeParticipants(s *scanner.Scanner, header Header) (*Participants, error) {
	p := &Participants{
		header: header,
	}
	err := s.ReadMemoryMappedStruct(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
