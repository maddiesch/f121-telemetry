package packet

import (
	"github.com/maddiesch/telemetry/internal/scanner"
)

type CarDamage struct {
	header Header

	Data [22]CarDamageData
}

type CarDamageData struct {
	TyresWear            [4]float32
	TyresDamage          [4]uint8
	BrakesDamage         [4]uint8
	FrontLeftWingDamage  uint8
	FrontRightWingDamage uint8
	RearWingDamage       uint8
	FloorDamage          uint8
	DiffuserDamage       uint8
	SidepodDamage        uint8
	DRSFault             uint8
	GearBoxDamage        uint8
	EngineDamage         uint8
	EngineMGUHWear       uint8
	EngineESWear         uint8
	EngineCEWear         uint8
	EngineICEWear        uint8
	EngineMGUKWear       uint8
	EngineTCWear         uint8
}

func (m *CarDamage) Header() Header { return m.header }

func (m *CarDamage) ID() ID { return PacketTypeCarDamage }

var _ Packet = (*CarDamage)(nil)

func decodeCarDamage(s *scanner.Scanner, header Header) (*CarDamage, error) {
	p := &CarDamage{
		header: header,
	}

	for i := 0; i < len(p.Data); i++ {
		d := CarDamageData{}

		err := decodeStructWithMemoryAlignedFields(s, &d)
		if err != nil {
			return nil, err
		}

		p.Data[i] = d
	}

	return p, nil
}
