package shine

import (
	"encoding/json"
	"testing"
)

type Outer struct {
	Value JSONOption[string] `json:"value"`
}

func TestJSONOption_Some(t *testing.T) {
	var jsonStr []byte
	{
		opt := NewSome("test")
		v := Outer{Value: JSONOption[string]{opt}}

		var err error
		jsonStr, err = json.Marshal(v)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if string(jsonStr) != `{"value":"test"}` {
			t.Log("invalid json string")
			t.FailNow()
		}
	}

	var v Outer
	if err := json.Unmarshal(jsonStr, &v); err != nil {
		t.Log(err)
		t.FailNow()
	}

	opt := v.Value.Option()
	if s, ok := opt.(Some[string]); ok {
		if s.Value() != "test" {
			t.Log("expected Some[value]")
			t.FailNow()
		}
	} else {
		t.Log("expected Some[value], got None")
		t.FailNow()
	}
}

func TestJSONOption_None(t *testing.T) {
	var jsonStr []byte
	{
		opt := NewNone[string]()
		v := Outer{Value: JSONOption[string]{opt}}

		var err error
		jsonStr, err = json.Marshal(v)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		if string(jsonStr) != `{"value":null}` {
			t.Log("invalid json string")
			t.FailNow()
		}
	}

	var v Outer
	if err := json.Unmarshal(jsonStr, &v); err != nil {
		t.Log(err)
		t.FailNow()
	}

	opt := v.Value.Option()
	if _, ok := opt.(None[string]); !ok {
		t.Log("expected None, got Some")
		t.FailNow()
	}
}
