package poker

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	a := "AsAhAdAc5s6c9d"
	b := "AsAhAdAcJs6c5d"
	w, r := Compare(a, b)
	fmt.Println(a, b, r, w)
}
