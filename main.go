package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

const (
	letters           = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	specialCharacters = "!@#$%^&*()-_=+[]{}|;:,.<>/?"
)

var digits = 8

func run() int {
	s := letters + specialCharacters
	c := len(s)

	buf := make([]byte, c)

	for i := 0; i < digits; i++ {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(c)))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return 1
		}
		buf[i] = s[r.Int64()]
	}

	fmt.Println(string(buf))
	return 0
}

func main() {
	code := run()
	os.Exit(code)
}
