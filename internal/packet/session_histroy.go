package packet

import (
	"github.com/maddiesch/telemetry/internal/scanner"
)

type SessionHistroy struct {
	header Header

	CarIndex              uint8
	LapCount              uint8
	TyreStintCount        uint8
	BestLapTimeLapNum     uint8
	BestSectorOneLapNum   uint8
	BestSectorTwoLapNum   uint8
	BestSectorThreeLapNum uint8

	LapHistroyData       [100]LapHistroyData
	TyreStintHistoryData [8]TyreStintHistoryData
}

type LapHistroyData struct {
	LapTimeMS         uint32
	SectorOneTimeMS   uint16
	SectorTwoTimeMS   uint16
	SectorThreeTimeMS uint16
	LapValidBitFlags  uint8
}

type TyreStintHistoryData struct {
	EndLap             uint8
	TyreActualCompound uint8
	TyreVisualCompound uint8
}

func (m *SessionHistroy) Header() Header { return m.header }

func (m *SessionHistroy) ID() ID { return PacketTypeSessionHistory }

var _ Packet = (*SessionHistroy)(nil)

func decodeSessionHistory(s *scanner.Scanner, header Header) (*SessionHistroy, error) {
	p := &SessionHistroy{
		header: header,
	}

	err := s.ReadMemoryMappedStruct(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
