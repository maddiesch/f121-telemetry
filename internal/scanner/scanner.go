package scanner

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"io"
	"math"
	"reflect"
)

type Scanner struct {
	data   []byte
	offset int
}

var (
	// The network byte order used by F1 2021
	ByteOrder = binary.LittleEndian
)

var (
	ErrInvalidPointer  = errors.New("must pass a pointer")
	ErrInvalidSettable = errors.New("must pass a settable value")
	ErrInvalidType     = errors.New("invalid type for scanner")
)

// New returns a new Scanner
func New(b []byte) *Scanner {
	return &Scanner{data: b}
}

func (s *Scanner) Available() int {
	return len(s.data) - s.offset
}

// Read consumes the number of bytes into the given byte slice
func (s *Scanner) Read(out []byte) (int, error) {
	c := len(out)
	if len(s.data) < s.offset+c {
		return 0, io.EOF
	}

	r := copy(out, s.data[s.offset:s.offset+c])

	s.offset += r

	return r, nil
}

func (s *Scanner) ReadMemoryMappedStruct(v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		panic("must pass a pointer")
	}
	return s.ReadValue(rv.Elem())
}

func (s *Scanner) ReadValue(sv reflect.Value) error {
	if !sv.CanSet() {
		return ErrInvalidSettable
	}

	get := func(c int) ([]byte, error) {
		b := make([]byte, c)
		_, err := s.Read(b)
		return b, err
	}

	switch sv.Kind() {
	case reflect.Int8:
		v, err := get(1)
		if err != nil {
			return err
		}
		sv.SetInt(int64(v[0]))
		return nil
	case reflect.Uint8:
		v, err := get(1)
		if err != nil {
			return err
		}
		sv.SetUint(uint64(v[0]))
		return nil
	case reflect.Int16:
		v, err := get(2)
		if err != nil {
			return err
		}
		i := ByteOrder.Uint16(v)
		sv.SetInt(int64(i))
		return nil
	case reflect.Uint16:
		v, err := get(2)
		if err != nil {
			return err
		}
		i := ByteOrder.Uint16(v)
		sv.SetUint(uint64(i))
		return nil
	case reflect.Int32:
		v, err := get(4)
		if err != nil {
			return err
		}
		i := ByteOrder.Uint32(v)
		sv.SetInt(int64(i))
		return nil
	case reflect.Uint32:
		v, err := get(4)
		if err != nil {
			return err
		}
		i := ByteOrder.Uint32(v)
		sv.SetUint(uint64(i))
		return nil
	case reflect.Int64:
		v, err := get(8)
		if err != nil {
			return err
		}
		i := ByteOrder.Uint64(v)
		sv.SetInt(int64(i))
		return nil
	case reflect.Uint64:
		v, err := get(8)
		if err != nil {
			return err
		}
		i := ByteOrder.Uint64(v)
		sv.SetUint(uint64(i))
		return nil
	case reflect.Float32:
		v, err := get(4)
		if err != nil {
			return err
		}
		i := ByteOrder.Uint32(v)
		f := math.Float32frombits(i)
		sv.SetFloat(float64(f))
		return nil
	case reflect.Float64:
		v, err := get(8)
		if err != nil {
			return err
		}
		i := ByteOrder.Uint64(v)
		f := math.Float64frombits(i)
		sv.SetFloat(f)
		return nil
	case reflect.Array:
		for i := 0; i < sv.Len(); i++ {
			subV := sv.Index(i)
			if err := s.ReadValue(subV); err != nil {
				return err
			}
		}
		return nil
	case reflect.Struct:
		for i := 0; i < sv.NumField(); i++ {
			f := sv.Field(i)

			if !f.CanSet() {
				continue
			}

			if err := s.ReadValue(f); err != nil {
				return err
			}
		}

		return nil
	default:
		return ErrInvalidType
	}
}

func (s *Scanner) ReadScalar(v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return ErrInvalidPointer
	}
	sv := rv.Elem()

	return s.ReadValue(sv)
}

func (s *Scanner) Export() string {
	return base64.RawStdEncoding.EncodeToString(s.data)
}
