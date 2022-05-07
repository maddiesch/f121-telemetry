package packet

import (
	"github.com/maddiesch/telemetry/internal/scanner"
)

type CarStatus struct {
	header Header

	Data [22]CarStatusData
}

type CarStatusData struct {
	TractionControl         uint8   // Traction control - 0 = off, 1 = medium, 2 = full
	AntiLockBrakes          uint8   // 0 (off) - 1 (on)
	FuelMix                 uint8   // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	FrontBrakeBias          uint8   // Front brake bias (percentage)
	PitLimiterStatus        uint8   // Pit limiter status - 0 = off, 1 = on
	FuelInTank              float32 // Current fuel mass
	FuelCapacity            float32 // Fuel capacity
	FuelRemainingLaps       float32 // Fuel remaining in terms of laps (value on MFD)
	MaxRPM                  uint16  // Cars max RPM, point of rev limiter
	IdleRPM                 uint16  // Cars idle RPM
	MaxGears                uint8   // Maximum number of gears
	DRSAllowed              uint8   // 0 = not allowed, 1 = allowed
	DRSActivationDistance   uint16  // 0 = DRS not available, non-zero - DRS will be available in [X] metres
	ActualTyreCompound      uint8   // F1 Modern - 16 = C5, 17 = C4, 18 = C3, 19 = C2, 20 = C1, 7 = inter, 8 = wet | F1 Classic - 9 = dry, 10 = wet | F2 – 11 = super soft, 12 = soft, 13 = medium, 14 = hard | 15 = wet
	VisualTyreCompound      uint8   // F1 16 = soft, 17 = medium, 18 = hard, 7 = inter, 8 = wet | F1 Classic 16 = soft, 17 = medium, 18 = hard, 7 = inter, 8 = wet | F2 ‘19, 15 = wet, 19 – super soft, 20 = soft, 21 = medium , 22 = hard
	TyresAgeLaps            uint8
	VehicleFIAFlags         int8    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
	ERSStoreEnergy          float32 // ERS energy store in Joules
	ERSDeployMode           uint8   // ERS deployment mode, 0 = none, 1 = medium, 2 = hotlap, 3 = overtake
	ERSHarvestedThisLapMGUK float32
	ERSHarvestedThisLapMGUH float32
	ERSDeployedThisLap      float32
	NetworkPaused           uint8
}

func decodeCarStatus(s *scanner.Scanner, h Header) (*CarStatus, error) {
	p := &CarStatus{
		header: h,
		Data:   [22]CarStatusData{},
	}

	for i := 0; i < len(p.Data); i++ {
		data := CarStatusData{}

		err := decodeStructWithMemoryAlignedFields(s, &data)
		if err != nil {
			return nil, err
		}

		p.Data[i] = data
	}

	return p, nil
}

func (m *CarStatus) Header() Header { return m.header }

func (m *CarStatus) ID() ID { return PacketTypeCarStatus }

var _ Packet = (*CarStatus)(nil)
