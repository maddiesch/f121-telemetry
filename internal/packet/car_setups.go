package packet

import (
	"github.com/maddiesch/telemetry/internal/scanner"
)

type CarSetups struct {
	header Header

	Data [22]CarSetupData
}

type CarSetupData struct {
	FrontWing              uint8
	RearWing               uint8
	OnThrottle             uint8
	OffThrottle            uint8
	FrontCamber            float32
	RearCamber             float32
	FrontToe               float32
	RearToe                float32
	FrontSuspension        uint8
	RearSuspension         uint8
	FrontAntiRollBar       uint8
	RearAntiRollBar        uint8
	FrontSuspensionHeight  uint8
	RearSuspensionHeight   uint8
	BrakePressure          uint8
	BrakeBias              uint8
	RearLeftTyrePressure   float32
	RearRightTyrePressure  float32
	FrontLeftTyrePressure  float32
	FrontRightTyrePressure float32
	Ballast                uint8
	FuelLoad               float32
}

func (m *CarSetups) Header() Header { return m.header }

func (m *CarSetups) ID() ID { return PacketTypeCarSetups }

var _ Packet = (*CarSetups)(nil)

func decodeCarSetups(s *scanner.Scanner, header Header) (*CarSetups, error) {
	p := &CarSetups{
		header: header,
	}

	for i := 0; i < len(p.Data); i++ {
		data := CarSetupData{}

		err := decodeStructWithMemoryAlignedFields(s, &data)
		if err != nil {
			return nil, err
		}

		p.Data[i] = data
	}

	return p, nil
}
