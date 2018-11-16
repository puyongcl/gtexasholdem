package compare

import (
	"fmt"
	"github.com/json-iterator/go"
	"gtexasholdem/poker"
	"io/ioutil"
)

type Matches struct {
	Matches []*Match `json:"matches"`
}

type Match struct {
	PlayerA string `json:"alice"`
	PlayerB string `json:"bob"`
	Result  int    `json:"result"`
}

// 获取牌组（必然获取，否则抛出异常）
func ReadMatches(path string) *Matches {
	var file []byte
	var err error

	if file, err = ioutil.ReadFile(path); err != nil {
		panic("panic: " + err.Error())
	}

	matches := Matches{}
	if err := jsoniter.Unmarshal(file, &matches); err != nil {
		panic("panic: " + err.Error())
	}
	return &matches
}

// 打印牌组比较结果
func (matches *Matches) ExcuteCompare() {
	var failCnt int
	for _, v := range matches.Matches {
		winner, res := Compare(v.PlayerA, v.PlayerB)

		if winner != v.Result {
			fmt.Printf("%s, %s : %s, %d, %d\n", v.PlayerA, v.PlayerB, res, winner, v.Result)
			failCnt ++
		}
	}
	fmt.Printf("合计：%d 条,failcnt : %d\n", len(matches.Matches), failCnt)
}

func Compare(a, b string) (int, string) {
	w1, err, res1 := poker.CalcWeight(a)
	if err != nil {
		return 0, ""
	}
	w2, err, res2 := poker.CalcWeight(b)
	if err != nil {
		return 0, ""
	}
	win := winner(w1, w2)

	s := fmt.Sprintf("%s:%07d--%s:%07d", res1, w1, res2, w2)
	return win, s
}

func winner(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return 2
	} else {
		return 0
	}
}
