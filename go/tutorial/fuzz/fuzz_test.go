package main

import (
	"testing"
    "unicode/utf8"
)

func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello World", " ", "!12344"}
    for _, tc := range testcases {
        f.Add(tc)
    }

    f.Fuzz(func(t *testing.T, a string) {
        rev, err1 := Reverse(a)
        if err1 != nil {
            return
        }
        doubleRev, err2 := Reverse(rev)
        if err2 != nil {
            return
        }
        t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(a), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
        if a != doubleRev {
            t.Errorf("Reverse(Reverse(%q)) = %q, want %q", a, doubleRev, a)
        }

        if utf8.ValidString(a) && !utf8.ValidString(rev) {
            t.Errorf("Reverse(%q) is not a valid UTF-8 string", rev)
        }
    })
}
