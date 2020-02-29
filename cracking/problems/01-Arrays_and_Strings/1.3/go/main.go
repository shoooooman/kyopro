package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve1(str string, l int) {
	var parsed bytes.Buffer
	cnt := 0
	for _, c := range str {
		if c == ' ' {
			parsed.Write([]byte("%20"))
			cnt++
		} else {
			parsed.WriteByte(byte(c))
			cnt++
		}
		if cnt == l {
			fmt.Println(parsed.String())
		}
	}
}

func main() {
	// read one line
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	si := strings.Split(input, ", ")

	str := si[0]
	if !(len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"') {
		fmt.Println("The string is wrong format.")
		os.Exit(1)
	}
	str = str[1 : len(str)-1]

	l, err := strconv.Atoi(si[1])
	if err != nil {
		fmt.Println("The length is wrong.")
		os.Exit(1)
	}

	solve1(str, l)
}
