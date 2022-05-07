package packet

import (
	"errors"

	"github.com/maddiesch/telemetry/internal/scanner"
)

type LobbyInfo struct {
	header Header
}

func (m *LobbyInfo) Header() Header { return m.header }

func (m *LobbyInfo) ID() ID { return PacketTypeLobbyInfo }

var _ Packet = (*LobbyInfo)(nil)

func decodeLobbyInfo(s *scanner.Scanner, header Header) (*LobbyInfo, error) {
	return nil, errors.New("decodeLobbyInfo - not implemented")
}
