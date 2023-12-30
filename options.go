package main

import (
	"flag"
)

const digits = 8

type options struct {
	digits int
	symbol bool
}

func parseOptions() *options {
	var l int
	var sym bool

	flag.IntVar(&l, "l", digits, "digits")
	flag.BoolVar(&sym, "b", false, "use special character")
	flag.Parse()

	return &options{
		digits: l,
		symbol: sym,
	}
}
