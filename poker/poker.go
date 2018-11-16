package poker

import (
	"errors"
	"regexp"
	"strings"
)

/*
扑克牌52张，花色黑桃spades，红心hearts，方块diamonds，草花clubs各13张，2-10，J，Q，K，A
Face：即2-10，J，Q，K，A，其中10用T来表示。
Color：即S(spades)、H(hearts)、D(diamonds)、C(clubs)
用 Face字母+小写Color字母表示一张牌，比如As表示黑桃A，其中A为牌面，s为spades，即黑桃，Ah即红心A，
以此类推
现分别给定任意两手牌各有5张，例如：AsAhQsQhQc，vs KsKhKdKc2c，请按德州扑克的大小规则来判断双方大
小。代码要求有通用性，可以任意输入一手牌或几手牌来进行比较。
*/

// 权值放大系数，防止累加抵消大的对子牌
const PairScale = 10000
const SecondPairScale = 10

// 比较顺子牌面A为最小牌面
var faceStraightAMin = map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10,
	"J": 11, "Q": 12, "K": 13, "A": 1, "X": 15}
// 比较顺子牌面A为大派面
var faceStraightAMax = map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10,
	"J": 11, "Q": 12, "K": 13, "A": 14, "X": 16}

// 权值比较
var faceWeight = map[string]int{"2": 1, "3": 2, "4": 4, "5": 8, "6": 16, "7": 32, "8": 64, "9": 128, "T": 256,
	"J": 512, "Q": 1024, "K": 2048, "A": 4096, "X": 8192}
// 权值比较
var faceWeightA = map[string]int{"2": 1, "3": 2, "4": 4, "5": 8, "6": 16, "7": 32, "8": 64, "9": 128, "T": 256,
	"J": 512, "Q": 1024, "K": 2048, "A": 0, "X": 8192}

// 牌面
var pokerFace = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var pokerColor = []string{"s", "h", "d", "c"}

const WeightScale = 10000000
const (
	// 判断皇家同花顺 同花色的A, K, Q, J和10.
	WeightRoyalStraightFlush = 9 * WeightScale
	// 判断同花顺 五张同花色的连续数字牌
	WeightStraightFlush = 8 * WeightScale
	// 判断四条 其中四张是相同点数但不同花的扑克牌，第五张是随意的一张牌
	WeightFourOfAKind = 7 * WeightScale
	// 判断满堂彩 由三张相同点数及任何两张其他相同点数的扑克牌组成
	WeightFullHouse = 6 * WeightScale
	// 判断同花 此牌由五张不按顺序但相同花的扑克牌组成
	WeightFlush = 5 * WeightScale
	// 判断顺子 此牌由五张顺序扑克牌组成
	WeightStraight = 4 * WeightScale
	// 判断三条 由三张相同点数和两张不同点数的扑克组成
	WeightThreeOfAKind = 3 * WeightScale
	// 判断两对 两对点数相同但两两不同的扑克和随意的一张牌组成
	WeightTwoPair = 2 * WeightScale
	// 判断一对 由两张相同点数的扑克牌和另三张随意的牌组成
	WeightOnePair = 1 * WeightScale
	// 单张大牌 既不是同一花色也不是同一点数的五张牌组成。
	//WeightHighCard = 0
)

// 是TJQKA这种牌面,A在判断顺子时，认为是大牌，否则认为是1
var isAMax = false

// 计算单纯牌面大小
func CalcWeightByPokerFace(in string) int {
	var weight int
	for i := 0; i < len(in); i += 2 {
		weight += faceWeight[string(in[i])]
	}
	return weight
}

// 计算单纯牌面大小
func CalcWeightByPokerFaceB(in string) int {
	var weight int
	for i := 0; i < len(in); i += 2 {
		weight += faceWeightA[string(in[i])]
	}
	return weight
}

func IsValidPokerFace(in string) (bool, error) {
	ok, err := regexp.MatchString(`^[2-9TJQKAX]+$`, in)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, errors.New("invalid pokerFace faceWeight")
	}
	return true, nil
}

func IsValidPokerColor(in string) (bool, error) {
	ok, err := regexp.MatchString(`^[shdcn]+$`, in)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, errors.New("invalid pokerFace color")
	}
	return true, nil
}

func IsValidPoker(in string) (bool, error) {
	l := len(in)
	if l > 14 {
		return false, errors.New("pokerFace cnt more than 7")
	}
	if l%2 != 0 {
		return false, errors.New("invalid pokerFace input")
	}
	for i := 0; i < len(in); i += 2 {
		if ok, err := IsValidPokerFace(string(in[i])); !ok {
			return false, err
		}
		if ok, err := IsValidPokerColor(string(in[i+1])); !ok {
			return false, err
		}
	}
	return true, nil
}

// 权值牌面排序
//func SortPoker(in string) string {
//	data := []byte(in)
//	for i := 0; i < len(in); i += 2 {
//		for j := 2; j < len(in)-i; j += 2 {
//			if faceWeight[string(data[j])] < faceWeight[string(data[j-2])] {
//				data[j], data[j-2] = data[j-2], data[j]
//				data[j+1], data[j-1] = data[j-1], data[j+1]
//			}
//		}
//	}
//	return string(data)
//}
// 权值牌面排序
func SortPoker(in string) string {
	base := [15][]string{}
	for i := 0; i < len(in); i += 2 {
		index := faceStraightAMax[in[i:i+1]]
		base[index] = append(base[index], in[i:i+2])
	}
	var result string
	for _, v := range base {
		for _, value := range v {
			result += value
		}
	}
	return result
}

// 顺子牌面排序
//func SortPokerStraight(in string) string {
//	data := []byte(in)
//	for i := 0; i < len(in); i += 2 {
//		for j := 2; j < len(in)-i; j += 2 {
//			if faceStraightAMin[string(data[j])] < faceStraightAMin[string(data[j-2])] {
//				data[j], data[j-2] = data[j-2], data[j]
//				data[j+1], data[j-1] = data[j-1], data[j+1]
//			}
//		}
//	}
//	return string(data)
//}
func SortPokerStraight(in string) string {
	base := [15][]string{}
	for i := 0; i < len(in); i += 2 {
		index := faceStraightAMin[in[i:i+1]]
		base[index] = append(base[index], in[i:i+2])
	}
	var result string
	for _, v := range base {
		for _, value := range v {
			result += value
		}
	}
	return result
}

// 有几张相同花色的牌
func CountSameColorPoker(in string) (int, string) {
	colorCnt := make(map[string]int, 5)
	for i := 0; i < len(in); i += 2 {
		colorCnt[string(in[i+1])]++
	}

	var max int
	var color string
	for c, cnt := range colorCnt {
		if cnt > max {
			max = cnt
			color = c
		}
	}
	return max, color
}

// 判断皇家同花顺 同花色的A, K, Q, J和10.
func IsRoyalStraightFlush(in string) (bool, string) {
	// 判断是否有5张花色相同的牌
	// 判断是否包含A K Q J 10
	for {
		if !strings.Contains(in, "T") {
			isAMax = false
			break
		}
		if !strings.Contains(in, "J") {
			isAMax = false
			break
		}
		if !strings.Contains(in, "Q") {
			isAMax = false
			break
		}
		if !strings.Contains(in, "K") {
			isAMax = false
			break
		}
		if !strings.Contains(in, "A") {
			isAMax = false
			break
		}
		isAMax = true
		break
	}

	if isAMax {
		maxPoker := string([]byte(in)[len(in)-10:])
		cnt, _ := CountSameColorPoker(maxPoker)
		if cnt < 5 {
			return false, ""
		}
		return true, maxPoker
	} else {
		return false, ""
	}
}

// 判断同花顺 五张同花色的连续数字牌
func IsStraightFlush(in string) (bool, string) {
	cnt, color := CountSameColorPoker(in)
	if cnt < 5 {
		return false, ""
	}

	cnt, _, res := CountStraight(in, color)
	if cnt < 5 {
		return false, ""
	}
	cnt, color = CountSameColorPoker(res)
	if cnt < 5 {
		return false, ""
	}
	return true, res
}

// 判断四条 其中四张是相同点数但不同花的扑克牌，第五张是随意的一张牌
func IsFourOfAKind(in string) (ok bool, maxPoker, face string) {
	data := []byte(in)
	sameFaceCnt := 1
	l := len(in)
	// 找出有四张相同的牌
	for i := 2; i < l; i += 2 {
		if in[i] != in[i-2] {
			sameFaceCnt = 1
			continue
		}
		sameFaceCnt++
		if 4 == sameFaceCnt {
			if i == l-2 { // 最后四张
				maxPoker = string(data[l-10:])
			} else if i == 6 { // 前四张
				maxPoker = string(data[:8]) + string(data[l-2:])
			} else {
				maxPoker = string(data[i-6:i+2]) + string(data[l-2:])
			}
			ok = true
			face = string(data[i])
			return
		}
	}
	return
}

// 判断满堂彩 由三张相同点数及任何两张其他相同点数的扑克牌组成
func IsFullHouse(in string) (ok bool, maxPoker, threeFace string) {
	data := []byte(in)
	sameFaceCnt, threeSameCnt := 1, 0
	threeSame, twoSame := false, false
	var three, two string
	l := len(in)
	// 找出有3张相同的牌和2张相同的牌
	for i := 2; i < l; i += 2 {
		if in[i] != in[i-2] {
			if sameFaceCnt >= 3 {
				three += string(data[i-6 : i])
				threeSame = true
				threeSameCnt++
				threeFace = string(data[i-2])
			} else if 2 == sameFaceCnt {
				two = string(data[i-4 : i])
				twoSame = true
			}
			sameFaceCnt = 1
			continue
		}
		sameFaceCnt++
	}
	if sameFaceCnt >= 3 {
		three += string(data[l-6:])
		threeSame = true
		threeSameCnt++
		threeFace = string(data[l-2])
	} else if 2 == sameFaceCnt {
		two = string(data[l-4:])
		twoSame = true
	}

	if threeSameCnt > 1 {
		maxPoker = string([]byte(three)[2:])
		twoSame = true
	} else if threeSameCnt == 1 {
		maxPoker = three + two
		maxPoker = SortPoker(maxPoker)
	} else {
		return
	}

	if threeSame && twoSame {
		ok = true
	} else {
		maxPoker = ""
	}
	return
}

// 判断同花 此牌由五张不按顺序但相同花的扑克牌组成
func IsFlush(in string) (ok bool, maxPoker string) {
	cnt, c := CountSameColorPoker(in)
	if cnt < 5 {
		return false, ""
	}

	data := []byte(in)
	for i := 0; i < len(in); i += 2 {
		if string(in[i+1]) == c {
			maxPoker += string(data[i : i+2])
		}
	}
	if len(maxPoker) > 10 {
		maxPoker = string([]byte(maxPoker)[len(maxPoker)-10:])
	}
	ok = true
	return
}

// 统计顺子的长度和牌面值和
func CountStraight(in, color string) (maxLen int, maxStraightWeight int, maxPoker string) {
	tmpFace := faceStraightAMax
	if !isAMax {
		in = SortPokerStraight(in)
		tmpFace = faceStraightAMin
	}
	weight := faceWeight[string(in[0])]
	sameCnt := 1
	l := len(in)
	data := []byte(in)
	straight := string(data[:2])
	var isSkipBehind bool // 第一个相同的是否已经跳过了

	for i := 2; i < l; i += 2 {
		if in[i] == in[i-2] {
			if !isSkipBehind {
				continue
			}
			isSkipBehind = false
		} else {
			c1 := color != "" && string(data[i+1]) != color
			c2 := tmpFace[string(data[i])]-tmpFace[string(data[i-2])] != 1
			if c1 || c2 {
				if c1 && !c2 {
					if i+2 < l && in[i] == in[i+2] { // 如果颜色不对，再看下一张相同的牌
						isSkipBehind = true
						continue
					}
				}

				if maxLen < sameCnt {
					maxLen = sameCnt
					maxStraightWeight = weight
					maxPoker = straight
				}
				weight = 0
				sameCnt = 1
				weight += faceWeight[string(in[i-2])]
				straight = string(data[i : i+2])
				continue
			}
		}
		straight += string(data[i : i+2])
		weight += faceWeight[string(in[i])]
		sameCnt++
	}

	if maxLen < sameCnt {
		maxLen = sameCnt
		maxStraightWeight = weight
		maxPoker = straight
	}

	if maxLen > 5 {
		maxPoker = string([]byte(maxPoker)[maxLen*2-10:])
	}
	return
}

// 判断三条 由三张相同点数和两张不同点数的扑克组成
func IsThreeOfAKind(in string) (ok bool, maxPoker, face string) {
	left, sameFaceCnt := 0, 1
	l := len(in)
	data := []byte(in)
	// 找出有3张相同的牌和2张相同的牌
	for i := 2; i < l; i += 2 {
		if in[i] != in[i-2] {
			if sameFaceCnt >= 3 {
				ok = true
				maxPoker = string(data[i-6 : i])
				left = i - 6
				face = string(data[i-2])
			}
			sameFaceCnt = 1
			continue
		}
		sameFaceCnt++
	}

	if sameFaceCnt >= 3 {
		ok = true
		maxPoker = string(data[l-6:])
		left = l - 6
		face = string(data[l-2])
	}

	if left >= l-10 {
		maxPoker = string(data[l-10:])
	} else {
		maxPoker += string(data[l-4:])
	}

	if !ok {
		maxPoker = ""
	}

	return
}

// 判断两对 两对点数相同但两两不同的扑克和随意的一张牌组成
func IsTwoPair(in string) (ok bool, maxPoker, pairFace string) {
	sameFaceCnt := 1
	sameCnt := 0
	l := len(in)
	locate := make(map[int]int)
	data := []byte(in)
	// 找出2张相同的牌
	for i := 2; i < l; i += 2 {
		if in[i] != in[i-2] {
			if sameFaceCnt >= 2 {
				maxPoker += string(data[i-4 : i])
				sameCnt++
				locate[sameCnt] = i - 4
				pairFace += string(in[i-4])
			}
			sameFaceCnt = 1
			continue
		}
		sameFaceCnt++
	}

	if sameFaceCnt >= 2 {
		maxPoker += string(data[l-4:])
		sameCnt++
		locate[sameCnt] = l - 4
		pairFace += string(in[l-4])
	}

	if sameCnt < 2 {
		maxPoker = ""
	} else if sameCnt == 2 {
		if locate[2]+4 != l {
			maxPoker += string(data[l-2:])
		} else if locate[2]+4 == l && locate[1]+4 != l-4 {
			maxPoker += string(data[l-6 : l-4])
		} else {
			maxPoker += string(data[locate[1]-2 : locate[1]])
		}
		maxPoker = SortPoker(maxPoker)
	} else {
		maxPoker = string(data[l-10:])
		pairFace = string([]byte(pairFace)[1:])
	}
	ok = sameCnt >= 2
	return
}

// 判断一对 由两张相同点数的扑克牌和另三张随意的牌组成
func IsOnePair(in string) (ok bool, maxPoker string, pairFace string) {
	sameFaceCnt, sameCnt, left := 1, 0, 0
	l := len(in)
	data := []byte(in)
	// 找出2张相同的牌
	for i := 2; i < l; i += 2 {
		if in[i] != in[i-2] {
			if sameFaceCnt >= 2 {
				sameCnt++
				left = i - 4
				maxPoker = string(data[i-4 : i])
				pairFace = string(in[i-4])
			}
			sameFaceCnt = 1
			continue
		}
		sameFaceCnt++
	}

	if sameFaceCnt >= 2 {
		sameCnt++
		left = l - 4
		maxPoker = string(data[l-4:])
		pairFace = string(in[l-4])
	}

	if sameCnt >= 1 {
		if left <= l-10 {
			maxPoker += string(data[l-6:]) // 2H3C4H5C6S6H7S
		} else if left > l-10 {
			maxPoker = string(data[left-(6-(l-(left+4))):left]) + maxPoker + string(data[left+4:])
			maxPoker = SortPoker(maxPoker)
		}
	}
	ok = sameCnt >= 1
	return
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 计算扑克牌权值
func CalcWeight(in string) (weight int, err error, maxPoker string) {
	// 计算赖子的个数
	xCnt := strings.Count(in, "X")
	if xCnt > 0 {
		return calcWeightByPokerFaceAndColorX(in, xCnt)
	}
	return calcWeightByPokerFaceAndColor(in)
}

// 替换癩子为普通牌
//func replaceX(in string, pokerPos int, color byte) (res string) {
//	l := len(in)
//	data := []byte(in)
//	var xPos int
//	for i := 0; i < l; i += 2 {
//		if data[i] == 'X' {
//			data[i] = []byte(pokerFace[pokerPos+xPos])[0]
//			data[i+1] = color
//			xPos++
//		}
//	}
//	return string(data)
//}

// 按扑克牌面大小和花色规则来增加权值,带癩子的牌
func calcWeightByPokerFaceAndColorX(in string, xCnt int) (maxWeight int, err error, maxPoker string) {
	if xCnt > 1 {
		panic("not implemented")
		return
	}
	var weight int
	l := len(pokerFace)
	for j := 0; j < 4; j++ {
		for i := 0; i < l; i++ {
			var onePoker, pokers string
			onePoker = pokerFace[i] + pokerColor[j]
			if strings.Contains(in, onePoker) {
				continue
			}

			pokers = strings.Replace(in, "Xn", onePoker, 1)
			weight, err, pokers = calcWeightByPokerFaceAndColor(pokers)
			if weight > maxWeight {
				maxWeight = weight
				maxPoker = pokers
			}
		}
	}
	return
}

// 按扑克牌面大小和花色规则来增加权值
func calcWeightByPokerFaceAndColor(in string) (weight int, err error, maxPoker string) {
	isAMax = false
	var ok bool
	var cnt int
	var pairFace string
	if ok, err = IsValidPoker(in); !ok {
		return
	}
	// sort
	in = SortPoker(in)

	if err != nil {
		return
	}

	if ok, maxPoker = IsRoyalStraightFlush(in); ok { // 判断皇家同花顺
		weight += CalcWeightByPokerFace(maxPoker)
		weight += WeightRoyalStraightFlush
		//return
	} else if ok, maxPoker = IsStraightFlush(in); ok { // 判断同花顺
		weight += CalcWeightByPokerFace(maxPoker)
		weight += WeightStraightFlush
		//return
	} else if ok, maxPoker, pairFace = IsFourOfAKind(in); ok { // 判断四条
		weight += CalcWeightByPokerFace(maxPoker)
		weight += WeightFourOfAKind
		weight += faceStraightAMax[pairFace] * PairScale * 4
		//return
	} else if ok, maxPoker, pairFace = IsFullHouse(in); ok { // 判断三带二
		weight += CalcWeightByPokerFace(maxPoker)
		weight += WeightFullHouse
		weight += faceStraightAMax[pairFace] * PairScale * 3
		//return
	} else if ok, maxPoker = IsFlush(in); ok { // 判断同花
		weight += CalcWeightByPokerFace(maxPoker)
		weight += WeightFlush
		//return
	} else if cnt, _, maxPoker = CountStraight(in, ""); cnt >= 5 { // 判断顺子
		if isAMax {
			weight += CalcWeightByPokerFace(maxPoker)
		} else {
			weight += CalcWeightByPokerFaceB(maxPoker)
		}
		weight = weight + WeightStraight
		//return
	} else if ok, maxPoker, pairFace = IsThreeOfAKind(in); ok { // 判断三条
		weight += CalcWeightByPokerFace(maxPoker)
		weight += WeightThreeOfAKind
		weight += faceStraightAMax[pairFace] * PairScale * 3
		//return
	} else if ok, maxPoker, pairFace = IsTwoPair(in); ok { // 判断两对
		weight += CalcWeightByPokerFace(maxPoker)
		weight += faceStraightAMax[string(pairFace[0])] * PairScale
		weight += faceStraightAMax[string(pairFace[1])] * PairScale * SecondPairScale
		weight += WeightTwoPair
		//return
	} else if ok, maxPoker, pairFace = IsOnePair(in); ok { // 判断一对
		weight += CalcWeightByPokerFace(maxPoker)
		weight += faceStraightAMax[pairFace] * PairScale
		weight += WeightOnePair
		//return
	} else { // 单张大牌
		maxPoker = string([]byte(in)[len(in)-10:])
		weight += CalcWeightByPokerFace(maxPoker)
	}
	return
}

// 按扑克牌面大小和花色规则来增加权值
func BCalcWeightByPokerFaceAndColor(in string) (weight int, err error, maxPoker string) {
	isAMax = false
	ok, err := IsValidPoker(in)
	if !ok {
		return
	}
	// sort
	in = SortPoker(in)

	data := []byte(in)
	l := len(in)

	// 判断皇家同花顺
	royal, _ := IsRoyalStraightFlush(in)
	if royal {
		maxPoker = string(data[l-10:])
		weight = 60
		weight += WeightRoyalStraightFlush
		return
	}

	// 判断同花顺
	sameColorCnt, color := CountSameColorPoker(in)
	straightCnt, _, maxStraight := CountStraight(in, color)
	if straightCnt >= 5 && sameColorCnt >= 5 {
		maxPoker = maxStraight
		weight = CalcWeightByPokerFace(maxPoker)
		weight = weight + WeightStraightFlush
		return
	}

	// 判断四条
	sameFaceCnt, sameCnt, threeSameCnt := 1, 0, 0
	threeSame, twoSame := false, false
	// 找出相同的牌个数
	for i := 2; i < l; i += 2 {
		if in[i] != in[i-2] {
			if sameFaceCnt >= 3 {
				threeSame = true
				sameCnt++
				threeSameCnt++
			} else if 2 == sameFaceCnt {
				twoSame = true
				sameCnt++
			}
			sameFaceCnt = 1
			continue
		}
		sameFaceCnt++
		if 4 == sameFaceCnt {
			if i-6 > l-10 {
				maxPoker = string(data[l-10:])
			} else if i-6 == 0 {
				maxPoker = string(data[:8]) + string(data[l-2:])
			} else {
				maxPoker = string(data[i-8:8]) + string(data[l-2:])
			}
			weight = CalcWeightByPokerFace(maxPoker)
			weight += WeightFourOfAKind
			return
		}
	}

	if sameFaceCnt >= 3 {
		maxPoker += string(data[l-6:])
		threeSame = true
		sameCnt++
		threeSameCnt++
	} else if 2 == sameFaceCnt {
		maxPoker += string(data[l-4:])
		twoSame = true
		sameCnt++
	}

	if threeSameCnt > 1 {
		twoSame = true
	}

	// 判断三带二
	if threeSame && twoSame {
		var pairFace string
		_, maxPoker, pairFace = IsFullHouse(in)
		weight = CalcWeightByPokerFace(maxPoker)
		weight += WeightFullHouse
		weight += faceStraightAMax[pairFace] * PairScale * 3
		return
	}

	// 判断同花
	sameColorCnt, _ = CountSameColorPoker(in)
	if sameColorCnt >= 5 {
		_, maxPoker = IsFlush(in)
		weight = CalcWeightByPokerFace(maxPoker)
		weight += WeightFlush
		return
	}

	// 判断顺子
	if straightCnt >= 5 {
		maxPoker = maxStraight
		if isAMax {
			weight = CalcWeightByPokerFace(maxPoker)
		} else {
			weight = CalcWeightByPokerFaceB(maxPoker)
		}
		weight += WeightStraight
		return
	}

	// 判断三条
	if threeSame == true && twoSame == false {
		var pairFace string
		_, maxPoker, pairFace = IsThreeOfAKind(in)
		weight = CalcWeightByPokerFace(maxPoker)
		weight += WeightThreeOfAKind
		weight += faceStraightAMax[pairFace] * PairScale * 3
		return
	}

	// 判断两对
	if sameCnt >= 2 {
		var pairFace string
		_, maxPoker, pairFace = IsTwoPair(in)
		weight = CalcWeightByPokerFace(maxPoker)
		weight += faceStraightAMax[string(pairFace[0])] * PairScale
		weight += faceStraightAMax[string(pairFace[1])] * PairScale * SecondPairScale
		weight += WeightTwoPair
		return
	}

	// 判断一对
	if sameCnt >= 1 {
		var pairFace string
		_, maxPoker, pairFace = IsOnePair(in)
		weight = CalcWeightByPokerFace(maxPoker)
		weight += faceStraightAMax[pairFace] * PairScale
		weight += WeightOnePair
		return
	}
	// 单张大牌
	maxPoker = string(data[len(in)-10:])
	weight = CalcWeightByPokerFace(maxPoker)
	return
}
