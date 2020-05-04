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

func main() {
	n, a, b, c := nextInt(), nextInt(), nextInt(), nextInt()
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = nextString()
	}

	nums := []int{a, b, c}
	ans := make([]rune, 0)
	for i, v := range s {
		left := rune(v[0]) - 'A'
		right := rune(v[1]) - 'A'
		if nums[left] == 0 && nums[right] == 0 {
			fmt.Println("No")
			return
		}
		if i != len(s)-1 && v != s[i+1] && nums[left] == 1 && nums[right] == 1 {
			next := s[i+1]
			nextLeft := rune(next[0]) - 'A'
			nextRight := rune(next[1]) - 'A'
			if left == nextLeft || left == nextRight {
				nums[left]++
				nums[right]--
				ans = append(ans, 'A'+rune(left))
			} else {
				nums[left]--
				nums[right]++
				ans = append(ans, 'A'+rune(right))
			}
		} else if nums[left] >= nums[right] {
			nums[left]--
			nums[right]++
			ans = append(ans, 'A'+rune(right))
		} else {
			nums[left]++
			nums[right]--
			ans = append(ans, 'A'+rune(left))
		}
	}
	fmt.Println("Yes")
	for _, v := range ans {
		fmt.Printf("%c\n", v)
	}
}
