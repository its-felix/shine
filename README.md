# shine
<small>like metal prior oxidation</small>

[![Go Reference](https://pkg.go.dev/badge/github.com/its-felix/shine.svg)](https://pkg.go.dev/github.com/its-felix/shine)
[![Go Report](https://goreportcard.com/badge/github.com/its-felix/shine?style=flat-square)](https://goreportcard.com/report/github.com/its-felix/shine)

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
v := shine.NewResult(strconv.Atoi("1")).UnwrapOr(1000)
```

#### Parsing URLs
```golang
func ExampleVanilla() string {
    u, err := url.Parse("https://github.com/")
    if err != nil {
        return ""
    }

    return u.Hostname()
}

func ExampleShine() string {
    return shine.ResMap(shine.NewResult(url.Parse("https://github.com/")), (*url.URL).Hostname).UnwrapOrDefault()
}
```
