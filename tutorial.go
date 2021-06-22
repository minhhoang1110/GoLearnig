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
func main() {
	mp := map[string]int{}
	input(mp)
	output(mp)
	val, ok := isContain(mp, "apple")
	fmt.Println(val, ok)
}
