package main

import (
	"fmt"
	"io"
	"math"
	"runtime"
	"strconv"
)

const (
	name                   = "factor"
	version                = "0.0.1"
	ExitCodeOK             = 0
	ExitCodeParseFlagError = 1
)

var revision = "HEAD"

type CLI struct {
	outStream, errStream io.Writer
}

func factor(n int) []int {

	var primeFactors []int

	if n < 2 {
		return []int{n}
	}

	for i := 2; i < int(math.Pow(float64(n), 0.5))+1; i++ {
		for n%i == 0 {
			primeFactors = append(primeFactors, i)
			n /= i
		}
	}

	if n > 1 {
		primeFactors = append(primeFactors, n)
	}

	return primeFactors
}

func (c *CLI) Run(args []string) int {
	if len(args) < 2 {
		fmt.Printf("%s %s (factor: %s/%s)\n", name, version, revision, runtime.Version())
		fmt.Println("$ factor 12")
		return ExitCodeParseFlagError
	}

	i, _ := strconv.Atoi(args[1])
	fmt.Fprintf(c.outStream, "%v\n", factor(i))

	return ExitCodeOK
}
