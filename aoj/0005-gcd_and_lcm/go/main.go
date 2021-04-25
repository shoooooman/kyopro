package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 10000
	maxBufSize     = 10000000
)

var (
	sc = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		sc.Split(bufio.ScanWords)
		return sc
	}()
)

func nextInt() (int, error) {
	if !sc.Scan() {
		return 0, errors.New("EOF")
	}
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i, nil
}

func main() {
	for {
		a, err := nextInt()
		if err != nil {
			return
		}
		b, err := nextInt()
		if b > a {
			a, b = b, a
		}
		fmt.Printf("%d %d\n", gcd(a, b), lcm(a, b))
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	g := gcd(a, b)
	return a * b / g
}
