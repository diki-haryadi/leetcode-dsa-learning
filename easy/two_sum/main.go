package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	result1 := twoSum([]int{3, 2, 4}, 6)
	took := time.Since(now)
	fmt.Println("result2 took time:", took)
	fmt.Println(result1)

	nowV1 := time.Now()
	result2 := twoSumVariant1([]int{3, 2, 4}, 5)
	fmt.Println("result2 took time:", time.Since(nowV1))
	fmt.Println(result2)

}
