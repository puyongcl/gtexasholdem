package poker

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcWeightByPokerFace(t *testing.T) {
	w := CalcWeightByPokerFace("AsAhQsQhQc")
	fmt.Println(w)
	w = CalcWeightByPokerFace("TsJhQsKhAd")
	fmt.Println(w)
}

func TestCountSameColorPoker(t *testing.T) {
	w, c := CountSameColorPoker("QsAsJs3s4s5sAs")
	fmt.Println(w, c)
	fmt.Println(5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5 * 5)
	fmt.Println(math.MaxInt64)
}

func TestSortPoker(t *testing.T) {
	w := SortPoker("AsKsJsQsTs7h9d")
	fmt.Println(w)
}

func TestIsRoyalstraightFlush(t *testing.T) {
	ok, res := IsRoyalStraightFlush("TsJsQsKsAs")
	assert.Equal(t, true, ok, "error exist")
	fmt.Println(res)
	ok, res = IsRoyalStraightFlush("3sJsQsKsAs")
	assert.Equal(t, false, ok, "error exist")
	fmt.Println(res)
}

func TestIsstraightFlush(t *testing.T) {
	ok, res := IsStraightFlush("3s4s5s6s7s8sXn")
	assert.Equal(t, true, ok, "error exist")
	fmt.Println(res)
	ok, res = IsStraightFlush("3s3s5s6s7s8s9s")
	assert.Equal(t, true, ok, "error exist")
	fmt.Println(res)
}

func TestIsstraight(t *testing.T) {
	v, v2, res := CountStraight("TsJsQsKsAs", "")
	assert.Equal(t, 5, v, "error exist")
	fmt.Println(v2, res)

	v, v2, res = CountStraight("3sJsQsKsAs", "")
	assert.Equal(t, 4, v, "error exist")
	fmt.Println(v2, res)
}

func TestIsFourOfAKind(t *testing.T) {
	ok, res, _ := IsFourOfAKind("2h2d2c2s3s3h4s")
	assert.Equal(t, true, ok, "error exist")
	fmt.Println(res)
	ok, res, _ = IsFourOfAKind("2h3d3c3h3s4d4s")
	assert.Equal(t, true, ok, "error exist")
	fmt.Println(res)
	ok, res, _ = IsFourOfAKind("2h3d3c4h4c4d4s")
	assert.Equal(t, true, ok, "error exist")
	fmt.Println(res)
	ok, res, _ = IsFourOfAKind("2h2d2c3h3c6d7s")
	assert.Equal(t, false, ok, "error exist")
	fmt.Println(res)
}

func TestIsFullhouse(t *testing.T) {
	ok, res, _ := IsFullHouse("2h2d2c3s4h4s")
	assert.Equal(t, true, ok, "error exist")
	fmt.Println(res)

	ok, res, _ = IsFullHouse("2h2d2c3h4c5s")
	assert.Equal(t, false, ok, "error exist")
	fmt.Println(res)
}

func TestIsFlush(t *testing.T) {
	ok, res := IsFlush("2h3h4h5h8h9hTh")
	assert.Equal(t, true, ok, "error exist")
	fmt.Println(res)
	ok, res = IsFlush("2h2d2c3h4c5s6h7h")
	assert.Equal(t, false, ok, "error exist")
	fmt.Println(res)
}

func TestIsThreeOfAKind(t *testing.T) {
	ok, res, _ := IsThreeOfAKind("2h2s2d5h8h9s")
	assert.Equal(t, true, ok, "error exist")
	fmt.Println(res)
	ok, res, _ = IsThreeOfAKind("2h2c3h4c5s6h7s")
	assert.Equal(t, false, ok, "error exist")
	fmt.Println(res)
}

func TestIsTwoPair(t *testing.T) {
	ok, res, _ := IsTwoPair("2h2s4d4h8h9sAh")
	assert.Equal(t, true, ok, "error exist")
	fmt.Println(res)
	ok, res, _ = IsTwoPair("2h2c3h4c5s6h7s")
	assert.Equal(t, false, ok, "error exist")
	fmt.Println(res)
}

func TestIsOnePair(t *testing.T) {
	ok, res, _ := IsOnePair("2h2s3d4h8h9sAh")
	assert.Equal(t, true, ok, "error exist")
	fmt.Println(res)
	ok, res, _ = IsOnePair("2h3c4h5c6s6h7s")
	assert.Equal(t, true, ok, "error exist")
	fmt.Println(res)
}

func TestCalcWeihtByPokerFaceAndColor(t *testing.T) {
	p := []string{"8s9sTsJsQsKsAs", "7s8s9sTsJsQsKs", "7s7h7d7cJsQsKs", "7s7h7dJhJsQsKs",
		"3s7s9sTsJsQsQc", "3h4s5d6c7s8c9s", "3s3h3d4sTsJcQd", "3s3d9s9dJdQcKs", "3s3d4s5c7c9dKs"}
	for _, s := range p {
		w1, err, res1 := calcWeightByPokerFaceAndColor(s)
		assert.NoError(t, err, "error exist")
		w2, err, res2 := BCalcWeightByPokerFaceAndColor(s)
		assert.NoError(t, err, "error exist")
		assert.Equal(t, w1, w2, "error exist")
		assert.Equal(t, res1, res2, "error exist")

		fmt.Println(w1, w2, res1, res2)
	}
}

func TestBcalcWeihtByPokerFaceAndcolor(t *testing.T) {
	p := []string{"8s9sTsJsQsKsAs", "7s8s9sTsJsQsKs", "7s7h7d7cJsQsKs", "7s7h7dJhJsQsKs",
		"3s7s9sTsJsQsQc", "3h4s5d6c7s8c9s", "3s3h3d4sTsJcQd", "3s3d9s9dJdQcKs", "3s3d4s5c7c9dKs"}
	for _, s := range p {
		w, err, res := BCalcWeightByPokerFaceAndColor(s)
		assert.NoError(t, err, "error exist")
		fmt.Println(w, res)
	}
}

func BenchmarkCalcWeihtByPokerFaceAndcolor(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calcWeightByPokerFaceAndColor("8s9sTsJsQsKsAs") // 皇家同花顺
		calcWeightByPokerFaceAndColor("7s8s9sTsJsQsKs") // 同花顺
		calcWeightByPokerFaceAndColor("7s7h7d7cJsQsKs") // 四条
		calcWeightByPokerFaceAndColor("7s7h7dJhJsQsKs") // 三带二
		calcWeightByPokerFaceAndColor("3s7s9sTsJsQsQc") // 同花
		calcWeightByPokerFaceAndColor("3h4s5d6c7s8c9s") // 顺子
		calcWeightByPokerFaceAndColor("3s3h3d4sTsJcQd") // 三条
		calcWeightByPokerFaceAndColor("3s3d9s9dJdQcKs") // 两对牌
		calcWeightByPokerFaceAndColor("3s3d4s5c7c9dKs") // 一对
	}
}

func BenchmarkBcalcWeihtByPokerFaceAndcolor(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BCalcWeightByPokerFaceAndColor("8s9sTsJsQsKsAs") // 皇家同花顺
		BCalcWeightByPokerFaceAndColor("7s8s9sTsJsQsKs") // 同花顺
		BCalcWeightByPokerFaceAndColor("7s7h7d7cJsQsKs") // 四条
		BCalcWeightByPokerFaceAndColor("7s7h7dJhJsQsKs") // 三带二
		BCalcWeightByPokerFaceAndColor("3s7s9sTsJsQsQc") // 同花
		BCalcWeightByPokerFaceAndColor("3h4s5d6c7s8c9s") // 顺子
		BCalcWeightByPokerFaceAndColor("3s3h3d4sTsJcQd") // 三条
		BCalcWeightByPokerFaceAndColor("3s3d9s9dJdQcKs") // 两对牌
		BCalcWeightByPokerFaceAndColor("3s3d4s5c7c9dKs") // 一对
	}
}

func BenchmarkSortPoker(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i ++ {
		SortPoker("AsKsJsQsTs7h9d")
	}
}
