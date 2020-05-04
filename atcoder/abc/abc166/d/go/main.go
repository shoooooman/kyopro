package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	const max = 1000
	mp := make(map[int]int)
	for i := 0; i <= max; i++ {
		k := i * i * i * i * i
		mp[k] = i
		if v, ok := mp[x+k]; ok {
			fmt.Println(i, v)
			return
		}
	}
	for i := 1; i >= -max; i-- {
		k := i * i * i * i * i
		mp[k] = i
		if v, ok := mp[x+k]; ok {
			fmt.Println(v, i)
			return
		}
	}
}
