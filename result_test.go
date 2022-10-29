package shine

import (
	"strconv"
	"testing"
)

func TestNewResultForErr(t *testing.T) {
	r := NewResult(strconv.Atoi("not a number"))

	if r.IsOk() {
		t.Log("NewResult(value, not nil) should return Err")
		t.FailNow()
	}
}

func TestNewResultForNonErr(t *testing.T) {
	r := NewResult(strconv.Atoi("1"))

	if r.IsErr() {
		t.Log("NewResult(value, nil) should return Ok")
		t.FailNow()
	}
}

func TestUnwrapOrDefaultForOk(t *testing.T) {
	r := NewResult(strconv.Atoi("1"))

	if r.UnwrapOrDefault() != 1 {
		t.Log("UnwrapOrDefault with Ok(1) should return 1")
		t.FailNow()
	}
}

func TestUnwrapOrDefaultForErr(t *testing.T) {
	r := NewResult(strconv.Atoi("not a number"))

	if r.UnwrapOrDefault() != 0 {
		t.Log("UnwrapOrDefault with Err[int] should return 0")
		t.FailNow()
	}
}
