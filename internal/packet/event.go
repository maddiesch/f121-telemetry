package packet

import (
	"fmt"

	"github.com/maddiesch/telemetry/internal/scanner"
)

type Event struct {
	header Header

	EventID [4]uint8

	FastestLap                *EventDetailFastestLap                `json:",omitempty"`
	Retirement                *EventDetailRetirement                `json:",omitempty"`
	TeamMateInPits            *EventDetailTeamMateInPits            `json:",omitempty"`
	RaceWinner                *EventDetailRaceWinner                `json:",omitempty"`
	Penalty                   *EventDetailPenalty                   `json:",omitempty"`
	SpeedTrap                 *EventDetailSpeedTrap                 `json:",omitempty"`
	StartLights               *EventDetailStartLights               `json:",omitempty"`
	DriveThroughPenaltyServed *EventDetailDriveThroughPenaltyServed `json:",omitempty"`
	StopGoPenaltyServed       *EventDetailStopGoPenaltyServed       `json:",omitempty"`
	Flashback                 *EventDetailFlashback                 `json:",omitempty"`
	Button                    *EventDetailButton                    `json:",omitempty"`
}

type EventDetailFastestLap struct {
	VehicleIndex uint8
	LapTime      float32
}

type EventDetailRetirement struct {
	VehicleIndex uint8
}

type EventDetailTeamMateInPits struct {
	VehicleIndex uint8
}

type EventDetailRaceWinner struct {
	VehicleIndex uint8
}

type EventDetailPenalty struct {
	PenaltyType       uint8
	VehicleIndex      uint8
	OtherVehicleIndex uint8
	Time              uint8
	LapNumber         uint8
	PlacesGained      uint8
}

type EventDetailSpeedTrap struct {
	VehicleIndex            uint8
	Speed                   float32
	OverallFastestInSession uint8
	DriverFastestInSession  uint8
}

type EventDetailStartLights struct {
	NumberLights uint8
}

type EventDetailDriveThroughPenaltyServed struct {
	VehicleIndex uint8
}

type EventDetailStopGoPenaltyServed struct {
	VehicleIndex uint8
}

type EventDetailFlashback struct {
	FrameID     uint32
	SessionTime float32
}

type EventDetailButton struct {
	ButtonStatus uint32
}

type EventCode [4]uint8

var (
	EventCodeSessionStarted     = EventCode{'S', 'S', 'T', 'A'}
	EventCodeSessionEnded       = EventCode{'S', 'E', 'N', 'D'}
	EventCodeFastestLap         = EventCode{'F', 'T', 'L', 'P'}
	EventCodeRetirement         = EventCode{'R', 'T', 'M', 'T'}
	EventCodeDrsEnabled         = EventCode{'D', 'R', 'S', 'E'}
	EventCodeDrsDisabled        = EventCode{'D', 'R', 'S', 'D'}
	EventCodeTeamMatePit        = EventCode{'T', 'M', 'P', 'T'}
	EventCodeCheckqueredFlag    = EventCode{'C', 'H', 'Q', 'F'}
	EventCodeRaceWinner         = EventCode{'R', 'C', 'W', 'N'}
	EventCodePenaltyIssued      = EventCode{'P', 'E', 'N', 'A'}
	EventCodeSpeedTrapTriggered = EventCode{'S', 'P', 'T', 'P'}
	EventCodeStartLights        = EventCode{'S', 'T', 'L', 'G'}
	EventCodeLightsOut          = EventCode{'L', 'G', 'O', 'T'}
	EventCodeDriveThroughServed = EventCode{'D', 'T', 'S', 'V'}
	EventCodeStopGoServed       = EventCode{'S', 'G', 'S', 'V'}
	EventCodeFlashback          = EventCode{'F', 'L', 'B', 'K'}
	EventCodeButtonStatus       = EventCode{'B', 'U', 'T', 'N'}
)

func (m *Event) Header() Header { return m.header }

func (m *Event) ID() ID { return PacketTypeEvent }

var _ Packet = (*Event)(nil)

func decodeEvent(s *scanner.Scanner, header Header) (*Event, error) {
	p := &Event{
		header: header,
	}

	if err := s.ReadScalar(&p.EventID); err != nil {
		return nil, err
	}

	switch p.EventID {
	case EventCodeSessionStarted:
		// no-op
	case EventCodeSessionEnded:
		// no-op
	case EventCodeFastestLap:
		p.FastestLap = &EventDetailFastestLap{}
		if err := s.ReadMemoryMappedStruct(p.FastestLap); err != nil {
			return nil, err
		}
	case EventCodeRetirement:
		p.Retirement = &EventDetailRetirement{}
		if err := s.ReadMemoryMappedStruct(p.Retirement); err != nil {
			return nil, err
		}
	case EventCodeDrsEnabled:
		// no-op
	case EventCodeDrsDisabled:
		// no-op
	case EventCodeTeamMatePit:
		p.TeamMateInPits = &EventDetailTeamMateInPits{}
		if err := s.ReadMemoryMappedStruct(p.TeamMateInPits); err != nil {
			return nil, err
		}
	case EventCodeCheckqueredFlag:
		// no-op
	case EventCodeRaceWinner:
		p.RaceWinner = &EventDetailRaceWinner{}
		if err := s.ReadMemoryMappedStruct(p.RaceWinner); err != nil {
			return nil, err
		}
	case EventCodePenaltyIssued:
		p.Penalty = &EventDetailPenalty{}
		if err := s.ReadMemoryMappedStruct(p.Penalty); err != nil {
			return nil, err
		}
	case EventCodeSpeedTrapTriggered:
		p.SpeedTrap = &EventDetailSpeedTrap{}
		if err := s.ReadMemoryMappedStruct(p.SpeedTrap); err != nil {
			return nil, err
		}
	case EventCodeStartLights:
		p.StartLights = &EventDetailStartLights{}
		if err := s.ReadMemoryMappedStruct(p.StartLights); err != nil {
			return nil, err
		}
	case EventCodeLightsOut:
		// no-op
	case EventCodeDriveThroughServed:
		p.DriveThroughPenaltyServed = &EventDetailDriveThroughPenaltyServed{}
		if err := s.ReadMemoryMappedStruct(p.DriveThroughPenaltyServed); err != nil {
			return nil, err
		}
	case EventCodeStopGoServed:
		p.StopGoPenaltyServed = &EventDetailStopGoPenaltyServed{}
		if err := s.ReadMemoryMappedStruct(p.StopGoPenaltyServed); err != nil {
			return nil, err
		}
	case EventCodeFlashback:
		p.Flashback = &EventDetailFlashback{}
		if err := s.ReadMemoryMappedStruct(p.Flashback); err != nil {
			return nil, err
		}
	case EventCodeButtonStatus:
		p.Button = &EventDetailButton{}
		if err := s.ReadMemoryMappedStruct(p.Button); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown event subtype: %s", string(p.EventID[:]))
	}

	return p, nil
}
