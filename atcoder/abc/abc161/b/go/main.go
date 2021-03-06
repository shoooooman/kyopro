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
	n, m := nextInt(), nextInt()
	a := make([]float64, n)
	sum := 0.0
	for i := 0; i < n; i++ {
		a[i] = nextFloat()
		sum += a[i]
	}

	cnt := 0
	for _, v := range a {
		if v < sum/float64((4*m)) {
			continue
		}
		cnt++
		if cnt >= m {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
