package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const (
	letters      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>/?"
)

func run() int {
	opts, err := parseOptions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}

	s := ""
	if opts.symbol {
		s = letters + specialChars
	} else {
		s = letters
	}

	c := int64(len(s))
	buf := make([][]byte, opts.num)

	for i := 0; i < opts.num; i++ {
		// Print password delimeter
		if i != 0 {
			if opts.delimeter != "" {
				unquotedDelimeter, err := strconv.Unquote(`"` + opts.delimeter + `"`)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
					return 1
				}
				fmt.Print(unquotedDelimeter)
			} else {
				fmt.Println()
			}
		}

		buf[i] = make([]byte, opts.digits)
		for j := 0; j < opts.digits; j++ {
			r, err := rand.Int(rand.Reader, big.NewInt(c))
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				return 1
			}
			buf[i][j] = s[r.Int64()]
		}

		fmt.Print(string(buf[i]))
	}
	fmt.Println()

	return 0
}

func main() {
	code := run()
	os.Exit(code)
}
