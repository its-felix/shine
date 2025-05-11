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

func TestNewOptionFromMap(t *testing.T) {
	myMap := make(map[string]string)
	opt := NewOptionFromMap(myMap, "key")

	if _, ok := opt.(Some[string]); ok {
		t.Log("NewOptionFromMap(empty map, key) should return None")
		t.FailNow()
	}

	myMap["key"] = "value"
	opt = NewOptionFromMap(myMap, "key")

	if s, ok := opt.(Some[string]); ok {
		if s.Value() != "value" {
			t.Log("NewOptionFromMap(non empty map, existing key) should return Some[value]")
			t.FailNow()
		}
	} else {
		t.Log("NewOptionFromMap(non empty map, existing key) should return Some")
		t.FailNow()
	}
}
