package scanner_test

import (
	"errors"
	"io"
	"testing"

	"github.com/maddiesch/telemetry/internal/scanner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScannerRead(t *testing.T) {
	t.Run("when there is enough data to read", func(t *testing.T) {
		s := scanner.New([]byte{'a', 'b', 'c', 'd'})

		s1 := make([]byte, 2)

		_, err := s.Read(s1)

		require.NoError(t, err)

		assert.Equal(t, []byte{'a', 'b'}, s1)

		s2 := make([]byte, 2)

		_, err = s.Read(s2)

		require.NoError(t, err)

		assert.Equal(t, []byte{'c', 'd'}, s2)
	})

	t.Run("when there is not enough data", func(t *testing.T) {
		s := scanner.New([]byte{'a', 'b', 'c', 'd'})

		s1 := make([]byte, 5)

		_, err := s.Read(s1)

		assert.True(t, errors.Is(io.EOF, err))
	})
}

func TestScannerReadValue(t *testing.T) {
	t.Run("int8", func(t *testing.T) {
		s := scanner.New([]byte{0x4f})

		var i int8

		if assert.NoError(t, s.ReadScalar(&i)) {
			assert.Equal(t, int8(79), i)
		}
	})

	t.Run("uint8", func(t *testing.T) {
		s := scanner.New([]byte{0xf3})

		var i uint8

		if assert.NoError(t, s.ReadScalar(&i)) {
			assert.Equal(t, uint8(243), i)
		}
	})

	t.Run("int16", func(t *testing.T) {
		s := scanner.New([]byte{0x4f, 0x00})

		var i int16

		if assert.NoError(t, s.ReadScalar(&i)) {
			assert.Equal(t, int16(79), i)
		}
	})

	t.Run("uint16", func(t *testing.T) {
		s := scanner.New([]byte{0xf3, 0x00})

		var i uint16

		if assert.NoError(t, s.ReadScalar(&i)) {
			assert.Equal(t, uint16(243), i)
		}
	})

	t.Run("int32", func(t *testing.T) {
		s := scanner.New([]byte{0x4f, 0x00, 0x00, 0x00})

		var i int32

		if assert.NoError(t, s.ReadScalar(&i)) {
			assert.Equal(t, int32(79), i)
		}
	})

	t.Run("uint32", func(t *testing.T) {
		s := scanner.New([]byte{0xf3, 0x00, 0x00, 0x00})

		var i uint32

		if assert.NoError(t, s.ReadScalar(&i)) {
			assert.Equal(t, uint32(243), i)
		}
	})

	t.Run("int64", func(t *testing.T) {
		s := scanner.New([]byte{0x4f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

		var i int64

		if assert.NoError(t, s.ReadScalar(&i)) {
			assert.Equal(t, int64(79), i)
		}
	})

	t.Run("uint64", func(t *testing.T) {
		s := scanner.New([]byte{0xf3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

		var i uint64

		if assert.NoError(t, s.ReadScalar(&i)) {
			assert.Equal(t, uint64(243), i)
		}
	})
}
