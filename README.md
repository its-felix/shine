# shine
<small>like metal prior oxidation</small>

[![Go Reference](https://pkg.go.dev/badge/github.com/its-felix/shine.svg)](https://pkg.go.dev/github.com/its-felix/shine)
[![Go Report](https://goreportcard.com/badge/github.com/labstack/echo?style=flat-square)](https://goreportcard.com/report/github.com/its-felix/shine)

Rust inspired implementation of `Option[T]` and `Result[T]` for Go.


---

## Get started
`go get github.com/its-felix/shine`

## Examples
### Interoperability with `(T, error)` receivers
#### Parsing numbers
```golang
// vanilla
v, err := strconv.Atoi("1")
if err != nil {
	v = 1000
}

// shine
r := shine.NewResult(strconv.Atoi("1")) // Result[int]
v := r.UnwrapOr(1000)
```

#### Parsing URLs
```golang
// vanilla
u, err := url.Parse("https://github.com/")
if err != nil {
	return err
}
hostname := u.Host

// shine
r := shine.NewResult(url.Parse("https://github.com/")) // Result[*url.URL]
return shine.ResultMap(r, func(u *url.URL) string{
	return u.Host
})
```

#### Reading JSON
```golang
// vanilla
func jsonRead(fname string) (map[string]any, error) {
    f, err := os.Open(fname)
    if err != nil {
        return nil, err
    }
    
    defer f.Close()
    dec := json.NewDecoder(f)
    var res map[string]any
    
    return res, dec.Decode(&res)
}


// shine
func jsonRead(fname string) Result[map[string]any] {
    r := shine.NewResult(os.Open(fname))
    
    return shine.ResultAndThen(r, func(f *os.File) Result[map[string]any] {
        defer f.Close()
        
        dec := json.NewDecoder(f)
        var res map[string]any
        
        return shine.NewResult(res, dec.Decode(&res))
    })
}
```

### Interoperability with `(T, error)` returns
```golang
func something() (string, error) {
	return shine.NewOk("hello").UnwrapBoth() // (string, error)
}
```

## Performance
Since both `Result` and `Option` are wrapper-types and not natively built into the language like in Rust, they come with some overhead.

Given the following benchmark:
```golang
// vanilla
something, err := getSomethingErrornous()
if err == nil {
	doSomething(something)// no-op
} else {
	doSomethingElse(err)// no-op
}

// shine
r := shine.NewResult(getSomethingErrornous())
if r.IsOk() {
    doSomething(r.Unwrap())// no-op
} else {
    doSomethingElse(r.UnwrapErr())// no-op
}
```

shine is 3 times slower than vanilla, with vanilla being at `1.248 ns/op` and shine at `3.617 ns/op` (MacBook Pro 2021 M1 Pro)