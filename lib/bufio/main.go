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
		sc.Split(bufio.ScanWords)
		return sc
	}()
)

// たくさん(>10^5)読み込みたいとき
func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// スペース区切りで読み込む
// sc.Split(bufio.ScanWords) が必要
func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
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

	s, t := nextLine(), nextLine()
	fmt.Printf("s=%v\n", s)
	fmt.Printf("t=%v\n", t)

	// sc.Split(bufio.ScanWords)
	// n, m := nextInt(), nextInt()
	// fmt.Printf("n=%v\n", n)
	// fmt.Printf("m=%v\n", m)

	// $ go run input.go | go run main.go
	str := readLine()
	fmt.Println(str)
}
