package packet

import (
	"github.com/maddiesch/telemetry/internal/scanner"
)

type Lap struct {
	header Header

	Data [22]LapData
}

type LapData struct {
	LastLapTimeMS        uint32
	CurrentLapTimeMS     uint32
	SectorOneTimeMS      uint16
	SectorTwoTimeMS      uint16
	LapDistance          float32
	TotalDistance        float32
	SafetyCarDelta       float32
	CarPosition          uint8
	CurrentLapNumber     uint8
	PitStatus            uint8
	PitStopCount         uint8
	Sector               uint8
	CurrentLapInvalid    uint8
	Penalties            uint8
	Warnings             uint8
	PendingDriveThroughs uint8
	PendingStopGo        uint8
	GridPosition         uint8
	DriverStatus         uint8
	ResultStatus         uint8
	PitLaneTimerActive   uint8
	PitLaneTimeMS        uint16
	PitStopTimeMS        uint16
	PitStopServePenalty  uint8
}

func (m *Lap) Header() Header { return m.header }

func (m *Lap) ID() ID { return PacketTypeLapData }

var _ Packet = (*Lap)(nil)

func decodeLapData(s *scanner.Scanner, header Header) (*Lap, error) {
	p := &Lap{
		header: header,
	}

	for i := 0; i < len(p.Data); i++ {
		d := LapData{}

		err := decodeStructWithMemoryAlignedFields(s, &d)
		if err != nil {
			return nil, err
		}

		p.Data[i] = d
	}

	return p, nil
}
