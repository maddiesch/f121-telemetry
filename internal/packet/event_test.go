package packet_test

import (
	"encoding/base64"
	"testing"

	"github.com/maddiesch/telemetry/internal/packet"
	"github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {
	packetData, _ := base64.RawStdEncoding.DecodeString(``)

	p, err := packet.DecodeStrict(packetData, true)

	if assert.NoError(t, err) {
		assert.NotZero(t, p.Header())
		assert.Equal(t, packet.PacketTypeEvent, p.ID())
	}
}
