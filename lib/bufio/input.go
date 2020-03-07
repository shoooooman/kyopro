package main

import (
	"bytes"
	"fmt"
)

// 長い行を生成
func main() {
	var buf bytes.Buffer
	for i := 0; i < 100000; i++ {
		buf.Write([]byte("a"))
	}
	fmt.Printf("%s:end\n", buf.String())
}
