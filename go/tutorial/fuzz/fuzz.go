package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)


func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("invalid UTF-8 string")
    }
    a := []rune(s)
    b := []rune{}
    for i := len(a) - 1; i >= 0; i-- {
        b = append(b, a[i])
    }
    return string(b), nil
}

func main() {
    s := "hello world fuu"
    r, _ := Reverse(s)
    fmt.Printf("r :%v\n", r)
}
