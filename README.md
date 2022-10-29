# shine
like metal before it oxidizes

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