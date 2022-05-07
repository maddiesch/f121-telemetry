package packet

import (
	"github.com/maddiesch/telemetry/internal/scanner"
)

type Motion struct {
	header Header

	Data [22]MotionData

	// Player Car Only Data
	SuspensionPosition     WheelArray
	SuspensionVelocity     WheelArray
	SuspensionAcceleration WheelArray
	WheelSpeed             WheelArray
	WheelSlip              WheelArray
	LocalVelocityX         float32
	LocalVelocityY         float32
	LocalVelocityZ         float32
	AngularVelocityX       float32
	AngularVelocityY       float32
	AngularVelocityZ       float32
	AngularAccelerationX   float32
	AngularAccelerationY   float32
	AngularAccelerationZ   float32
	FrontWheelsAngle       float32
}

type MotionData struct {
	WorldPositionX     float32
	WorldPositionY     float32
	WorldPositionZ     float32
	WorldVelocityX     float32
	WorldVelocityY     float32
	WorldVelocityZ     float32
	WorldForwardDirX   int16
	WorldForwardDirY   int16
	WorldForwardDirZ   int16
	WorldRightDirX     int16
	WorldRightDirY     int16
	WorldRightDirZ     int16
	GForceLateral      float32
	GForceLongitudinal float32
	GForceVertical     float32
	Yaw                float32
	Pitch              float32
	Roll               float32
}

func (m *Motion) Header() Header { return m.header }

func (m *Motion) ID() ID { return PacketTypeMotion }

var _ Packet = (*Motion)(nil)

func decodeMotion(s *scanner.Scanner, header Header) (*Motion, error) {
	packet := &Motion{
		header: header,
	}

	err := s.ReadMemoryMappedStruct(packet)
	if err != nil {
		return nil, err
	}

	return packet, nil
}
