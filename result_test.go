package shine

import (
	"errors"
	"os"
	"strconv"
	"testing"
)

type closeableValue struct {
	closed bool
}

func (v *closeableValue) Close() error {
	v.closed = true
	return nil
}

type closeableError struct {
	closed bool
}

func (v *closeableError) Close() error {
	v.closed = true
	return nil
}

func (v *closeableError) Error() string {
	return ""
}

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

func TestErrorIs(t *testing.T) {
	switch r := NewResult(strconv.Atoi("not a number")).(type) {
	case Ok[int]:
		t.Fail()

	case Err[int]:
		if !errors.Is(r, strconv.ErrSyntax) {
			t.Fail()
		}
	}
}

func TestErrorAs(t *testing.T) {
	err := NewErr[struct{}](new(os.PathError))

	var pe *os.PathError
	if !errors.As(err, &pe) {
		t.Fail()
	}
}

func TestOk_Close(t *testing.T) {
	v := new(closeableValue)
	r := NewOk(v)
	_ = r.Close()

	if !v.closed {
		t.Fail()
	}
}

func TestErr_Close(t *testing.T) {
	v := new(closeableError)
	r := NewErr[struct{}](v)
	_ = r.Close()

	if !v.closed {
		t.Fail()
	}
}
