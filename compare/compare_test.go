package compare

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	a := "AsAhAdAc5s"
	b := "AsAhAdAcJs"
	rr := 2
	w, r := Compare(a, b)
	fmt.Println(a, b, r, w, rr)
}

func LoopA() {
	a, b := 2, 1
	for i := 0; i < 14; i++ {
		if a > b {
			c := a
			b = a + b
			a = c * b
			if a > b {
				c = a
				b = a + b
				a = c * b
				if a > b {
					c = a
					b = a + b
					a = c * b
				}
			}
		}

		if a > b {
			c := a
			b = a + b
			a = c * b
			if a > b {
				c = a
				b = a + b
				a = c * b
				if a > b {
					c = a
					b = a + b
					a = c * b
				}
			}
		}

		if a > b {
			c := a
			b = a + b
			a = c * b
			if a > b {
				c = a
				b = a + b
				a = c * b
				if a > b {
					c = a
					b = a + b
					a = c * b
				}
			}
		}

		if a > b {
			c := a
			b = a + b
			a = c * b
			if a > b {
				c = a
				b = a + b
				a = c * b
				if a > b {
					c = a
					b = a + b
					a = c * b
				}
			}
		}

		if a > b {
			c := a
			b = a + b
			a = c * b
			if a > b {
				c = a
				b = a + b
				a = c * b
				if a > b {
					c = a
					b = a + b
					a = c * b
				}
			}
		}

		if a > b {
			c := a
			b = a + b
			a = c * b
			if a > b {
				c = a
				b = a + b
				a = c * b
				if a > b {
					c = a
					b = a + b
					a = c * b
				}
			}
		}
	}
}

func LoopB() {
	a, b := 2, 1
	for i := 0; i < 14; i++ {
		if a > b {
			c := a
			b = a + b
			a = c * b
			if a > b {
				c = a
				b = a + b
				a = c * b
				if a > b {
					c = a
					b = a + b
					a = c * b
				}
			}
		}
	}
	for i := 0; i < 14; i++ {
		if a > b {
			c := a
			b = a + b
			a = c * b
			if a > b {
				c = a
				b = a + b
				a = c * b
				if a > b {
					c = a
					b = a + b
					a = c * b
				}
			}
		}
	}
	for i := 0; i < 14; i++ {
		if a > b {
			c := a
			b = a + b
			a = c * b
			if a > b {
				c = a
				b = a + b
				a = c * b
				if a > b {
					c = a
					b = a + b
					a = c * b
				}
			}
		}
	}
	for i := 0; i < 14; i++ {
		if a > b {
			c := a
			b = a + b
			a = c * b
			if a > b {
				c = a
				b = a + b
				a = c * b
				if a > b {
					c = a
					b = a + b
					a = c * b
				}
			}
		}
	}
	for i := 0; i < 14; i++ {
		if a > b {
			c := a
			b = a + b
			a = c * b
			if a > b {
				c = a
				b = a + b
				a = c * b
				if a > b {
					c = a
					b = a + b
					a = c * b
				}
			}
		}
	}
	for i := 0; i < 14; i++ {
		if a > b {
			c := a
			b = a + b
			a = c * b
			if a > b {
				c = a
				b = a + b
				a = c * b
				if a > b {
					c = a
					b = a + b
					a = c * b
				}
			}
		}
	}
}

func BenchmarkA(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopA()
	}
}

func BenchmarkB(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoopB()
	}
}
