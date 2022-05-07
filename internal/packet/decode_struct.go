package packet

import (
	"reflect"

	"github.com/maddiesch/telemetry/internal/scanner"
)

func decodeStructWithMemoryAlignedFields(s *scanner.Scanner, v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		panic("must pass a pointer")
	}
	sv := rv.Elem()

	return s.ReadValue(sv)
}
