package main

import (
	"bufio"
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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextFloat() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

func main() {
	_, m := nextInt(), nextInt()
	ac := make(map[int]bool)
	pn := make(map[int]int)
	ps := 0
	for i := 0; i < m; i++ {
		p, s := nextInt(), nextString()
		if !ac[p] && s == "AC" {
			ac[p] = true
			ps += pn[p]
		} else if !ac[p] && s == "WA" {
			pn[p]++
		}
	}
	fmt.Printf("%d %d\n", len(ac), ps)
}
