// Password generator which creates random passwords and runs it through
// a test to determine which is the most complex.
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Define variables for usage in generation and strength test
const lower = "abcdefghijklmnopqrstuvwxyz"
const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const spechars = "`~!@#$%^&*()-_=+[{}];:'\"\\/?.>,<"
const chars = lower + upper + spechars

// rPassword returns a randomly generated string
func rPassword() string {
	var randSeed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 20+rand.Intn(10))
	for i := range b {
		b[i] = chars[randSeed.Intn(len(chars))]
	}
	return string(b)
}

// strengthTest determines strength of password and returns
// boolean (true false) if password is strong enough for usage.
func strengthTest(password string) bool {
	var lowercount int
	var uppercount int
	var speccount int
	for _, ch := range password {
		switch true {
		case len(password) < 21:
			return false
		case speccount > 9:
			return true
		case lowercount > 10:
			return true
		case uppercount > 10:
			return true
		case strings.Contains(lower, string(ch)):
			lowercount++
		case strings.Contains(upper, string(ch)):
			uppercount++
		case strings.Contains(spechars, string(ch)):
			speccount++
		}
	}
	return false
}

func main() {
	var passwords [10]string
	fmt.Println("GoLang Password Generator")
	fmt.Println("--- Generated Passwords ---")
	for i := 0; i < 10; i++ {
		passwords[i] = rPassword()
		time.Sleep(10000) // Sleep to ensure randomization
		if !strengthTest(passwords[i]) {
			passwords[i] = ""
		} else if passwords[i] != "" {
			fmt.Println(passwords[i])
		}
	}
	fmt.Println("--- Generated Passwords ---")
}
