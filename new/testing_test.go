package main

import (
	"regexp"
	"testing"

	"github.com/m-rishikesh/go-lang/mymodule/utils"
)

func TestSayHello(t *testing.T) {
	name := "rohit"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := utils.SayHello(name)
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("raj") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
