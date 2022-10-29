package shine

import (
	"testing"
)

func TestNewOptionForNil(t *testing.T) {
	var s *string
	opt := NewOption(s)

	if opt.IsSome() {
		t.Log("NewOption(nil) should return None")
		t.FailNow()
	}
}

func TestNewOptionForNonNil(t *testing.T) {
	opt := NewOption(&struct{}{})

	if opt.IsNone() {
		t.Log("NewOption(not nil) should return Some")
		t.FailNow()
	}
}
