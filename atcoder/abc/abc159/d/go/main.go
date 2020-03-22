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
	n := nextInt()
	a := make([]int, n)
	count := make(map[int]int)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		count[a[i]]++
	}

	combs := make(map[int]int)
	sum := 0
	for k, v := range count {
		comb := v * (v - 1) / 2
		combs[k] = comb
		sum += comb
	}

	for i := 0; i < n; i++ {
		ans := sum - combs[a[i]]
		c := count[a[i]] - 1
		ans += c * (c - 1) / 2
		fmt.Println(ans)
	}
}
