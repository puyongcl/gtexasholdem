package main

import (
	"fmt"
	"gtexasholdem/compare"
	"time"
)

func main() {
	//total  10000 hands, spending 20346183648 nano seconds
	beginTime := time.Now()
	m := compare.ReadMatches("/home/qydev/go/src/gtexasholdem/test/seven_cards_result.json")
	m.ExcuteCompare()
	finishTime := time.Now()
	fmt.Printf("共耗时：%d ns\n", finishTime.Sub(beginTime).Nanoseconds())

	beginTime = time.Now()
	m = compare.ReadMatches("/home/qydev/go/src/gtexasholdem/test/result.json")
	m.ExcuteCompare()
	finishTime = time.Now()
	fmt.Printf("共耗时：%d ns\n", finishTime.Sub(beginTime).Nanoseconds())

	beginTime = time.Now()
	m = compare.ReadMatches("/home/qydev/go/src/gtexasholdem/test/seven_cards_with_ghost_result.json")
	m.ExcuteCompare()
	finishTime = time.Now()
	fmt.Printf("共耗时：%d ns\n", finishTime.Sub(beginTime).Nanoseconds())
}
