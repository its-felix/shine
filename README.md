# shine
<small>like metal prior oxidation</small>

[![Go Reference](https://pkg.go.dev/badge/github.com/its-felix/shine/v4.svg)](https://pkg.go.dev/github.com/its-felix/shine/v4)
[![Go Report](https://goreportcard.com/badge/github.com/its-felix/shine/v4?style=flat-square)](https://goreportcard.com/report/github.com/its-felix/shine/v4)

Rust inspired implementation of `Option[T]` and `Result[T]` for Go.


---

## Get started
`go get github.com/its-felix/shine/v4`

## Examples
### Interoperability with `(T, error)` receivers
#### Parsing numbers

```golang
package example

import (
	"github.com/its-felix/shine/v4"
	"strconv"
)

func exampleVanilla() int {
	// vanilla
	v, err := strconv.Atoi("1")
	if err != nil {
		v = 1000
	}

	return v
}

func exampleShine() int {
	return shine.NewResult(strconv.Atoi("1")).UnwrapOr(1000)
}
```

#### Parsing URLs

```golang
package example

import (
	"github.com/its-felix/shine/v4"
	"net/url"
)

func exampleVanilla() string {
	u, err := url.Parse("https://github.com/")
	if err != nil {
		return ""
	}

	return u.Hostname()
}

func exampleShine() string {
	if r, ok := shine.NewResult(url.Parse("https://github.com/")).(shine.Ok[*url.URL]); ok {
		return r.Value().Hostname()
    }
	
	return ""
}
```

#### Closeable Support

```golang
package example

import (
	"github.com/its-felix/shine/v4"
	"os"
)

func example() {
	r := shine.NewResult(os.Open("file.txt"))
	// will call Close() on the contained value if it implements io.Closer
	defer r.Close()
}
```

#### Type switch

```golang
package example

import (
	"errors"
	"github.com/its-felix/shine/v4"
	"os"
)

func example() {
	r := shine.NewResult(os.Open("file.txt"))
	defer r.Close()

	switch r := r.(type) {
	case shine.Ok[*os.File]:
		f := r.Value()
		// do something with file

	case shine.Err[*os.File]:
		// handle error
		var pe *os.PathError
		if errors.As(r, &pe) {
			// handle path error 
		}
	}
}
```

```golang
package example

import "github.com/its-felix/shine/v4"

func example() {
	someMap := make(map[string]string)
	o := shine.NewOptionFromMap(someMap, "test")

	switch o := o.(type) {
	case shine.Some[string]:
		v := o.Value()
	    // do something with value

	case shine.None[string]:
		// handle none
	}
}
```

#### JSON support for Option

```golang
package example

import (
	"encoding/json"
	"github.com/its-felix/shine/v4"
)

type MyStruct struct {
	Value shine.JSONOption[string] `json:"value"`
}

func example() {
	some := shine.NewSome("hello world")
	b, err := json.Marshal(MyStruct{Value: shine.JSONOption[string]{some}})
	// {"value":"hello world"}
	
	none := shine.NewNone[string]()
	b, err := json.Marshal(MyStruct{Value: shine.JSONOption[string]{none}})
	// {"value":null}
}
```