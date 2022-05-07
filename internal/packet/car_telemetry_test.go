package packet_test

import (
	"encoding/base64"
	"testing"

	"github.com/maddiesch/telemetry/internal/packet"
	"github.com/stretchr/testify/assert"
)

func TestCarTelemetry(t *testing.T) {
	packetData, _ := base64.RawStdEncoding.DecodeString(`5QcBEgEGD9GPo5li73jlB8hCbAgAABP/4gAAAIA/VH5RPAAAAAAABoAoAAAAABQDEgP9AvoCWVFWTVlWWVdaAAAArEEAAKxBAAC4QQAAuEEAAAAAgwDzbws/vm2JPgAAAAAAA54nAAAAAIkDiAObA5kDVkxdUFdVV1ZaAAAArEEAAKxBAAC4QQAAuEEAAAAApAAoPSY/6GMSPgAAAAAABNwnAAAAAF8DXgNhA18DWk9eTVdVWFZaAAAArEEAAKxBAAC4QQAAuEEAAAAAdQAW+KM+m4uDPgAAAAAAA08iAAAAAK4DrAPIA8UDVUxdUldWV1ZaAAAArEEAAKxBAAC4QQAAuEEAAAAAAQEAAIA/cBPYuwAAAAAAB5YtAEzgf6wCqwKBAn4CVlFSTFlWWVdaAAAArEEAAKxBAAC4QQAAuEEAAAAAxQAAAIA/6oEePQAAAAAABcwoAAAAADcDNgMuAysDW05cTFhWWVZaAAAArEEAAKxBAAC4QQAAuEEAAAAAlwCu1UY/1ocdPgAAAAAAAy0tAD/gA3gDdgOCA38DWk9eTldVV1VaAAAArEEAAKxBAAC4QQAAuEEAAAAACQEAAIA/AAAAgAAAAAAAB5kpAAAAAJECkAJgAl4CVlJRTFlXWldaAAAArEEAAKxBAAC4QQAAuEEAAAAA+gAAAAAARHZ0OyYwYj8ABnosACxgAOIC4QIOAw0DU1NRUVdVVlVaAAAArEEAAKxBAAC4QQAAuEEBAAEArgAl6kU/BJfzPQAAAAAABBYqAAAAAFgDVgNZA1YDW05eTVhVWFZaAAAArEEAAKxBAAC4QQAAuEEAAAAA3gAAAIA/F2+DPAAAAAAABewtAFbgfxIDEQP9AvoCWU9YSlhWWVdaAAAArEEAAKxBAAC4QQAAuEEBAAEAiAAAAAAAJQU2Pj5WFj8AAxonAAAAANADzwP7A/kDUU9VVFdWV1VaAAAArEEAAKxBAAC4QQAAuEEAAAAAsgAAAAAAsjsUPXrwvj4ABHYqAAAAAJ4DnQPSA9EDUVFSUldWWVVaAAAArEEAAKxBAAC4QQAAuEEBAAEAKAEAAIA/qLOVOwAAAAAAB24uAGTgf5EBjwFHAUQBVFVQUFZVVlRaAAAArEEAAKxBAAC4QQAAuEEAAAAAzQAAAAAAtFi3OwAAAAAABRMqAAAAAG8DbgOmA6YDUVFSUldVVlVaAAAArEEAAKxBAAC4QQAAuEEBAAEA3AAAAIA/qLMVuwAAAAAABXctAEngfxADDwP8AvkCV09USlhWWldaAAAArEEAAKxBAAC4QQAAuEEBAAEAcgAMRQI+teamPgAAAAAAA2ohAAAAALkDuAPVA9MDVExbUldWV1VaAAAArEEAAKxBAAC4QQAAuEEAAAAAjQDSTB4/tKpDPgAAAAAAA3YqAAAAAHkDdwOHA4UDWE1eT1dVV1ZaAAAArEEAAKxBAAC4QQAAuEEAAAAAdgAAAAAAa+qYPikbOz0AAzYiAAAAAM8DzQP1A/IDUk1XVFdVVlVaAAAArEEAAKxBAAC4QQAAuEEAAAAAGQEAAIA/AAAAAAAAAAAABy8sACQAAIACgAI6AjgCWlhTUF9dYF5pACWBukFkM7lB7W+zQT+rskEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA//8A`)

	p, err := packet.DecodeStrict(packetData, true)

	if assert.NoError(t, err) {
		assert.NotZero(t, p.Header())
		assert.Equal(t, packet.PacketTypeCarTelemetry, p.ID())
	}
}
