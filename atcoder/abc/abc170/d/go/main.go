package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	sort.Ints(a)
	flags := make(map[int]bool)
	dup := make(map[int]bool)
	ok := make(map[int]bool)
	for _, v := range a {
		flag := true
		for i := 1; i*i <= v; i++ {
			if v%i == 0 {
				if flags[i] {
					flag = false
				}
				if i*i != v {
					if flags[v/i] {
						flag = false
					}
				}
			}
		}
		if ok[v] {
			dup[v] = true
		}
		if flag && !flags[v] {
			ok[v] = true
		}
		flags[v] = true
	}
	fmt.Println(len(ok) - len(dup))
}
