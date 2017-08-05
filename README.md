uintvar
=======

This package allows to parse a slice of bytes represeting a [uintvar](https://en.wikipedia.org/wiki/Variable-length_quantity#/media/File:Uintvar_coding.svg) which is a varible length encoding of a uint to a uint64 in Go.

## Example
```go
bytes, _ := hex.DecodeString("850c")
number, bytesUsed, err := uintvar.Parse(bytes)
// = 652, 2, nil
```