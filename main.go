package main

import (
	"fmt"
)

func validate(key string, mp map[string]float32) bool {
	if val, ok := mp[key]; ok {
		fmt.Println(val)
		return true
	} else {
		return false
	}
}

func solve(amount int, src string, target string, mp map[string]float32) {
	key := src + target

	// fmt.Println(key)
	if !validate(key, mp) {
		fmt.Println("The Currency is invalid or in Lower Letter , please Follow Capital letters")
	} else {
		fmt.Println(amount, src, "is equivalent to ", float32(amount)*mp[key], target)
	}
}

func main() {
	var mp map[string]float32     // map has been Created but it points to NULL right Now
	mp = make(map[string]float32) //The make function allocates and initializes a hash map data structure
	// and returns a map value that points to

	mp["USDINR"] = 83.12
	mp["USDEUR"] = 0.93
	mp["USDJPY"] = 156.82
	mp["INRUSD"] = 0.012
	mp["EURUSD"] = 1.07
	mp["JPYUSD"] = 0.0064

	var amount int
	var src string
	var target string

	fmt.Println("Enter the Amount , source , target in this Format one By One")
	fmt.Scan(&amount, &src, &target)

	//fmt.Println("These are Your Details", amount, src, target)

	solve(amount, src, target, mp)
}
