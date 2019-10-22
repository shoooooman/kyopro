package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	status := make([]string, h)
	for i := 0; i < h; i++ {
		var s string
		fmt.Scan(&s)
		status[i] = s
	}

	count := make([][]int, h)
	for i := 0; i < h; i++ {
		count[i] = make([]int, w)
	}
	x := []int{-1, 0, 1}
	y := []int{-1, 0, 1}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if status[i][j] == '#' {
				continue
			}

			cnt := 0
			for _, dx := range x {
				for _, dy := range y {
					if dx == 0 && dy == 0 {
						continue
					}
					if i+dx >= 0 && i+dx < h &&
						j+dy >= 0 && j+dy < w &&
						status[i+dx][j+dy] == '#' {
						cnt++
					}
				}
			}
			count[i][j] = cnt
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if status[i][j] == '#' {
				fmt.Printf("#")
			} else {
				fmt.Printf("%d", count[i][j])
			}
		}
		fmt.Println()
	}
}
