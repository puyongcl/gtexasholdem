package main

import (
	"fmt"
	"gtexasholdem/poker"
	"time"
)

func main() {
	//total  10000 hands, spending 20346183648 nano seconds
	beginTime := time.Now()
	m := poker.ReadMatches("/home/qydev/go/src/gtexasholdem/input/seven_cards_result.json")
	m.ExcuteCompare()
	m = poker.ReadMatches("/home/qydev/go/src/gtexasholdem/input/result.json")
	m.ExcuteCompare()
	finishTime := time.Now()
	fmt.Printf("共耗时：%d ns\n", finishTime.Sub(beginTime).Nanoseconds())
}
