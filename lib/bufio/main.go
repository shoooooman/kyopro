package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

/* ----------------------------------------------- */

// 長くない行(< 64*1024)を一気に読み込む
func getStdin() []string {
	stdin := bufio.NewScanner(os.Stdin)
	buf := make([]byte, initialBufSize)
	stdin.Buffer(buf, maxBufSize)
	lines := []string{}
	for stdin.Scan() {
		if err := stdin.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		lines = append(lines, stdin.Text())
	}
	if err := stdin.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return lines
}

/* ----------------------------------------------- */

// defaultだと64*1024しかない
// var sc = bufio.NewScanner(os.Stdin)
var (
	sc = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		sc.Split(bufio.ScanWords) // 行単位で読みたいときは消す
		return sc
	}()
)

// たくさん(>10^5)読み込みたいとき
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

/* ----------------------------------------------- */

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

// 長い行(>64*1024)を読み込みたいとき
func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

/* ----------------------------------------------- */

func main() {
	lines := getStdin()
	for i, line := range lines {
		fmt.Println(i, line)
	}

	s, t := nextString(), nextString()
	fmt.Printf("s=%s\n", s)
	fmt.Printf("t=%s\n", t)

	n, m := nextInt(), nextInt()
	fmt.Printf("n=%d\n", n)
	fmt.Printf("m=%d\n", m)

	f, g := nextFloat(), nextFloat()
	fmt.Printf("f=%f\n", f)
	fmt.Printf("g=%f\n", g)

	// $ go run input.go | go run main.go
	str := readLine()
	fmt.Println(str)
}
