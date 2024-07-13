package shine

import (
	"testing"
)

func TestNewOptionForNil(t *testing.T) {
	var s *string
	opt := NewOption(s)

	if _, ok := opt.(Some[*string]); ok {
		t.Log("NewOption(nil) should return None")
		t.FailNow()
	}
}

func TestNewOptionForNonNil(t *testing.T) {
	opt := NewOption(&struct{}{})

	if _, ok := opt.(None[*struct{}]); ok {
		t.Log("NewOption(not nil) should return Some")
		t.FailNow()
	}
}
