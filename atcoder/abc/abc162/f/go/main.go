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

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func calc(a []int, begin, passRemain, remain int) int {
	if v, ok := dp[key{begin, passRemain}]; ok {
		return v
	}

	if remain <= 0 {
		return 0
	}

	// if len(a) <= 2 {
	if len(a[begin:]) <= 2 {
		// return a[0]
		dp[key{begin, passRemain}] = a[begin]
		return dp[key{begin, passRemain}]
	}

	if passRemain > 0 {
		// sum1 := calc(a[2:], passRemain)
		sum1 := calc(a, begin+2, passRemain, remain-1)
		m := sum1
		// if len(a) >= 4 {
		if len(a[begin:]) >= 4 {
			// sum2 := calc(a[3:], passRemain-1)
			sum2 := calc(a, begin+3, passRemain-1, remain-1)
			if sum2 > m {
				m = sum2
			}
			// if len(a) >= 5 && passRemain >= 2 {
			if len(a[begin:]) >= 5 && passRemain >= 2 {
				// sum3 := calc(a[4:], passRemain-2)
				sum3 := calc(a, begin+4, passRemain-2, remain-1)
				if sum3 > m {
					m = sum3
				}
			}
		}
		// return a[0] + m
		dp[key{begin, passRemain}] = a[begin] + m
		return dp[key{begin, passRemain}]
	}
	// return a[0] + calc(a[2:], passRemain)
	dp[key{begin, passRemain}] = a[begin] + calc(a, begin+2, passRemain, remain-1)
	return dp[key{begin, passRemain}]
}

type key struct {
	Begin      int
	PassRemain int
}

var dp map[key]int

func main() {
	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	dp = make(map[key]int)

	var ans int
	if n%2 == 0 {
		// ans = max(calc(a, 1), calc(a[1:], 1))
		ans = max(calc(a, 0, 1, n/2), calc(a, 1, 1, n/2))
	} else {
		// s1 := calc(a, 2)
		// s2 := calc(a[1:], 1)
		// s3 := calc(a[2:], 0)
		s1 := calc(a, 0, 2, n/2)
		s2 := calc(a, 1, 1, n/2)
		s3 := calc(a, 2, 0, n/2)
		s4 := max(s1, s2)
		ans = max(s3, s4)
	}
	fmt.Println(ans)
}
