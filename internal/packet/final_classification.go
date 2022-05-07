package packet

import (
	"github.com/maddiesch/telemetry/internal/scanner"
)

type FinalClassification struct {
	header Header

	NumberCars uint8
	Data       [22]FinalClassificationData
}

type FinalClassificationData struct {
	Position         uint8
	NumLaps          uint8
	GridPosition     uint8
	Points           uint8
	PitStopCount     uint8
	ResultStatus     uint8
	BestLapTimeMS    uint32
	TotalRaceTime    float64
	PenaltiesTimeSec uint8
	PenaltiesCount   uint8
	TyreStintCount   uint8
	TyreUsedActual   [8]uint8
	TyreUsedVisual   [8]uint8
}

func (m *FinalClassification) Header() Header { return m.header }

func (m *FinalClassification) ID() ID { return PacketTypeFinalClassification }

var _ Packet = (*FinalClassification)(nil)

func decodeFinalClassification(s *scanner.Scanner, header Header) (*FinalClassification, error) {
	p := &FinalClassification{
		header: header,
	}

	if err := s.ReadScalar(&p.NumberCars); err != nil {
		return nil, err
	}

	for i := 0; i < len(p.Data); i++ {
		err := scannerReadAll(s,
			&p.Data[i].Position,
			&p.Data[i].NumLaps,
			&p.Data[i].GridPosition,
			&p.Data[i].Points,
			&p.Data[i].PitStopCount,
			&p.Data[i].ResultStatus,
			&p.Data[i].BestLapTimeMS,
			&p.Data[i].TotalRaceTime,
			&p.Data[i].PenaltiesTimeSec,
			&p.Data[i].PenaltiesCount,
			&p.Data[i].TyreStintCount,
			&p.Data[i].TyreUsedActual[0],
			&p.Data[i].TyreUsedActual[1],
			&p.Data[i].TyreUsedActual[2],
			&p.Data[i].TyreUsedActual[3],
			&p.Data[i].TyreUsedActual[4],
			&p.Data[i].TyreUsedActual[5],
			&p.Data[i].TyreUsedActual[6],
			&p.Data[i].TyreUsedActual[7],
			&p.Data[i].TyreUsedVisual[0],
			&p.Data[i].TyreUsedVisual[1],
			&p.Data[i].TyreUsedVisual[2],
			&p.Data[i].TyreUsedVisual[3],
			&p.Data[i].TyreUsedVisual[4],
			&p.Data[i].TyreUsedVisual[5],
			&p.Data[i].TyreUsedVisual[6],
			&p.Data[i].TyreUsedVisual[7],
		)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}
