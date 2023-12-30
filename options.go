package main

import (
	"errors"
	"flag"
)

const (
	digits    = 8
	outputMax = 50
)

type options struct {
	digits int
	symbol bool
	num    int
}

func parseOptions() (*options, error) {
	var l int
	var sym bool
	var n int

	flag.IntVar(&l, "l", digits, "digits")
	flag.BoolVar(&sym, "b", false, "use special character")
	flag.IntVar(&n, "n", 1, "output number")
	flag.Parse()

	if n < 1 {
		return nil, errors.New("option \"n\" must be greater than 0")
	} else if n > 100 {
		return nil, errors.New("option \"n\" must be less than 101")
	}

	return &options{
		digits: l,
		symbol: sym,
		num:    n,
	}, nil
}
