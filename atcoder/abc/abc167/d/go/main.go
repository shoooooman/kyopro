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
	n, k := nextInt(), nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt() - 1
	}

	visited := make(map[int]int)
	visited[0] = 0
	t, start := n, 0
	s := []int{0}
	next := a[0]
	for i := 0; i < n; i++ {
		if v, ok := visited[next]; ok {
			t = i + 1 - v
			start = v
			break
		}
		s = append(s, next)
		visited[next] = i + 1
		next = a[next]
	}

	var ans int
	if k < start {
		ans = s[k] + 1
	} else {
		mod := (k - start) % t
		ans = s[start+mod] + 1
	}
	fmt.Println(ans)
}
