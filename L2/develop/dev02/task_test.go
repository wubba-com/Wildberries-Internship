package main

import (
	"testing"
)

func TestRepeatS(t *testing.T)  {
	m := make(map[string]string)
	m["a4bc2d5e"] = "aaaabccddddde"
	m["abcd"] = "abcd"
	m["45"] = ""
	m[""] = ""

	for k, v := range m {
		if s := RepeatS(k); s != v {
			t.Error("result != value")
		}
	}
}
