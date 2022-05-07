package packet

import (
	"github.com/maddiesch/telemetry/internal/scanner"
)

type CarTelemetry struct {
	header Header

	Data [22]CarTelemetryData

	MFDPanelIndex                uint8
	MFDPanelIndexSecondaryPlayer uint8
	SuggestedGear                int8
}

type CarTelemetryData struct {
	Speed            uint16
	Throttle         float32
	Steer            float32
	Brake            float32
	Clutch           uint8
	Gear             int8
	EngineRPM        uint16
	DRS              uint8
	RevLightPercent  uint8
	RevLightBitValue uint16
	BrakeTempC       [4]uint16
	TireSurfaceTempC [4]uint8
	TireInnerTempC   [4]uint8
	EngineTempC      uint16
	TirePressurePSI  [4]float32
	SurfaceType      [4]uint8
}

func (m *CarTelemetry) Header() Header { return m.header }

func (m *CarTelemetry) ID() ID { return PacketTypeCarTelemetry }

var _ Packet = (*CarTelemetry)(nil)

func decodeCarTelemetry(s *scanner.Scanner, header Header) (*CarTelemetry, error) {
	p := &CarTelemetry{
		header: header,
	}

	err := s.ReadMemoryMappedStruct(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
