package main

import "fmt"

/* solution1 */
// func main() {
// 	var k, s int
// 	fmt.Scan(&k, &s)
//
// 	count := 0
// 	double := 0
// 	triple := 0
// 	for x := 0; x <= k; x++ {
// 		for y := x; y <= k; y++ {
// 			for z := y; z <= k; z++ {
// 				if x+y+z == s {
// 					count++
// 					if x == y && y == z {
// 						triple++
// 					} else if x == y || y == z || z == x {
// 						double++
// 					}
// 				}
// 			}
// 		}
// 	}
// 	count = (count-(double+triple))*6 + double*3 + triple
// 	fmt.Println(count)
// }

/* solution2 */
func main() {
	var k, s int
	fmt.Scan(&k, &s)

	count := 0
	for x := 0; x <= k; x++ {
		for y := 0; y <= k; y++ {
			z := s - x - y
			if z >= 0 && z <= k {
				count++
			}
		}
	}
	fmt.Println(count)
}
