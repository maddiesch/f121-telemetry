// Package packet contains the telemetry packet format and decoding for F1 2021 gameplay telemetry
package packet

import (
	"fmt"

	"github.com/maddiesch/telemetry/internal/scanner"
)

const (
	HeaderLength = 24
)

type ID uint8

const (
	PacketTypeMotion ID = iota
	PacketTypeSession
	PacketTypeLapData
	PacketTypeEvent
	PacketTypeParticipants
	PacketTypeCarSetups
	PacketTypeCarTelemetry
	PacketTypeCarStatus
	PacketTypeFinalClassification
	PacketTypeLobbyInfo
	PacketTypeCarDamage
	PacketTypeSessionHistory
)

type Packet interface {
	ID() ID
	Header() Header
}

type WheelArray [4]float32

func (w WheelArray) FrontLeft() float32 {
	return w[2]
}

func (w WheelArray) FrontRight() float32 {
	return w[3]
}

func (w WheelArray) RearLeft() float32 {
	return w[0]
}

func (w WheelArray) RearRight() float32 {
	return w[1]
}

type StrictDecodeError struct {
	Message string
}

func (e *StrictDecodeError) Error() string {
	return fmt.Sprintf("strict decoder error: %s", e.Message)
}

func DecodeStrict(data []byte, requireFullBuffer bool) (Packet, error) {
	s := scanner.New(data)
	header, err := decodeHeader(s)
	if err != nil {
		return nil, err
	}

	defer func() {
		if requireFullBuffer && s.Available() != 0 {
			panic(&StrictDecodeError{
				Message: fmt.Sprintf("the packet decoder failed to consume all available bytes in packet (%d)", s.Available()),
			})
		}
	}()

	switch ID(header.PacketID) {
	case PacketTypeMotion:
		return decodeMotion(s, header)
	case PacketTypeSession:
		return decodeSession(s, header)
	case PacketTypeLapData:
		return decodeLapData(s, header)
	case PacketTypeEvent:
		return decodeEvent(s, header)
	case PacketTypeParticipants:
		return decodeParticipants(s, header)
	case PacketTypeCarSetups:
		return decodeCarSetups(s, header)
	case PacketTypeCarTelemetry:
		return decodeCarTelemetry(s, header)
	case PacketTypeCarStatus:
		return decodeCarStatus(s, header)
	case PacketTypeFinalClassification:
		return decodeFinalClassification(s, header)
	case PacketTypeLobbyInfo:
		return decodeLobbyInfo(s, header)
	case PacketTypeCarDamage:
		return decodeCarDamage(s, header)
	case PacketTypeSessionHistory:
		return decodeSessionHistory(s, header)
	default:
		return nil, &ErrUnknownPacketType{ID: header.PacketID}
	}
}

func Decode(data []byte) (Packet, error) {
	return DecodeStrict(data, false)
}

func scannerReadAll(s *scanner.Scanner, values ...any) error {
	for _, v := range values {
		if err := s.ReadScalar(v); err != nil {
			return err
		}
	}
	return nil
}

type ErrUnknownPacketType struct {
	ID uint8
}

func (e *ErrUnknownPacketType) Error() string {
	return fmt.Sprintf("unknown packet type: %d", e.ID)
}
