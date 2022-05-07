package packet

import (
	"github.com/maddiesch/telemetry/internal/scanner"
)

type Session struct {
	header Header

	Weather                    uint8
	TrackTemperatureC          int8
	AirTemperatureC            int8
	TotalLaps                  uint8
	TrackLengthM               uint16
	SessionType                uint8
	TrackID                    int8
	Formula                    uint8
	SessionTimeLeftSec         uint16
	SessionDurationSec         uint16
	PitSpeedLimit              uint8
	GamePaused                 uint8
	IsSpectating               uint8
	SpectatorCarIndex          uint8
	SLIProNativeSupport        uint8
	MarshalZoneCount           uint8
	MarshalZones               [21]MarshalZone
	SafetyCarStatus            uint8
	NetworkGame                uint8
	WeatherForecastSampleCount uint8
	WeatherForecasts           [56]WeatherForecastSample
	ForecaseAccuracy           uint8
	AIDifficulty               uint8
	SeasonLinkID               uint32
	WeekendLinkIdentifier      uint32
	SessionLinkIdentifier      uint32
	PitStopWindowIdealStartLap uint8
	PitStopWindowLatestLap     uint8
	PitStopRejoinPosition      uint8
	SteeringAssist             uint8
	BrakingAssist              uint8
	GearboxAssist              uint8
	PitAssist                  uint8
	PitReleaseAssist           uint8
	ERSAssist                  uint8
	DRSAssist                  uint8
	DynamicRacingLine          uint8
	DynamicRacingLineType      uint8
}

type MarshalZone struct {
	ZoneStart float32
	ZoneFlag  int8
}

type WeatherForecastSample struct {
	SessionType            uint8
	TimeOffset             uint8
	Weather                uint8
	TrackTemperatureC      int8
	TrackTemperatureChange int8
	AirTemperatureC        int8
	AirTemperatureChange   int8
	RainPercent            uint8
}

func (m *Session) Header() Header { return m.header }

func (m *Session) ID() ID { return PacketTypeSession }

var _ Packet = (*Session)(nil)

func decodeSession(s *scanner.Scanner, header Header) (*Session, error) {
	p := &Session{
		header: header,
	}

	err := s.ReadMemoryMappedStruct(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
