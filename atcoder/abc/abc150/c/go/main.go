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

func exclude(str string, i int) string {
	ret := str[:i]
	ret += str[i+1:]
	return ret
}

func perm(pat, str string, ret *[]string) {
	if len(str) == 0 {
		*ret = append(*ret, pat)
		return
	}

	for i := 0; i < len(str); i++ {
		perm(pat+string(str[i]), exclude(str, i), ret)
	}
}

func fact(n int) int {
	if n <= 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	n := nextInt()
	a := 0
	b := 0
	for i := 0; i < n; i++ {
		a *= 10
		a += nextInt()
	}
	for i := 0; i < n; i++ {
		b *= 10
		b += nextInt()
	}

	str := ""
	for i := 0; i < n; i++ {
		str += strconv.Itoa(i + 1)
	}
	l := fact(n)
	ret := make([]string, 0, l)
	perm("", str, &ret)

	as := strconv.Itoa(a)
	bs := strconv.Itoa(b)
	var ai, bi int
	for i := 0; i < l; i++ {
		if as == ret[i] {
			ai = i
		}
		if bs == ret[i] {
			bi = i
		}
	}
	if ai >= bi {
		fmt.Println(ai - bi)
	} else {
		fmt.Println(bi - ai)
	}
}
