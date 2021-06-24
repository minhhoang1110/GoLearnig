package main

import (
	"fmt"
)

func input(mp map[string]int) {
	defer fmt.Println("Input completed !!!!")
	isInput := true
	for isInput {
		var key string
		var val int
		var checkContinue int
		fmt.Print("Input key and value: ")
		fmt.Scan(&key, &val)
		mp[key] = val
		fmt.Print("Do you wanna continue input ? (1/0): ")
		fmt.Scan(&checkContinue)
		if checkContinue == 0 {
			isInput = false
		}
	}
}
func output(mp map[string]int) {
	for i, element := range mp {
		fmt.Printf("Key \"%s\" have value: %d", i, element)
	}
	fmt.Println()
}
func isContain(mp map[string]int, index string) (int, bool) {
	val, ok := mp[index]
	return val, ok
}
func retrunFunc(x string) func(string) {
	return func(str string) {
		fmt.Println(x + str)
	}
}
func changeVal(x *int) {
	*x += 10
}

type Point struct {
	x int
	y int
}

func (point *Point) changePoint() {
	point.x += 10
}

func main() {
	// mp := map[string]int{}
	// input(mp)
	// output(mp)
	// val, ok := isContain(mp, "apple")
	// fmt.Println(val, ok)
	// test := func(x int) int {
	// 	return x
	// }(1)
	// fmt.Println(test)
	// x := retrunFunc("hello")
	// x("!!!")
	//mutable and immutable data type
	// a1 := []int{1, 2, 3}
	// a2 := a1
	// a2[0] = 10
	// fmt.Println(a1, a2)
	//pointer
	// x := 2
	// fmt.Println(x)
	// changeVal(&x)
	// fmt.Println(x)
	point := Point{1, 2}
	point.changePoint()
	fmt.Println(point)
}
