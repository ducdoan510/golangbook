package main

import (
	"bytes"
	"fmt"
)

const BitsetSize = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/BitsetSize, uint(x%BitsetSize)
	return word < len(s.words) && (s.words[word] & (1<<bit) != 0)
}

func (s *IntSet) Add(x int) {
	word, bit := x/BitsetSize, uint(x%BitsetSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1<<bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for word, bitset := range t.words {
		if word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= bitset
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for word, bitset := range t.words {
		if word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] &= bitset
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for word, bitset := range t.words {
		if word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] &^= bitset
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for word, bitset := range t.words {
		if word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] ^= bitset
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for word, bitset := range s.words {
		for bit := 0; bit < BitsetSize; bit++ {
			if bitset & (1 << bit) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", BitsetSize* word + bit)
			}
		}
	}
	buf.WriteByte('}')
	fmt.Fprintf(&buf, ": length %d", s.Len())
	return buf.String()
}

func (s *IntSet) Len() (len int) {
	for _, bitset := range s.words {
		for bit := 0; bit < BitsetSize; bit++ {
			if bitset & (1 << bit) != 0 {
				len += 1
			}
		}
	}
	return
}

func (s *IntSet) Remove(x int) {
	word, bit := x/BitsetSize, uint(x%BitsetSize)
	if word < len(s.words) {
		s.words[word] &^= 1<<bit
	}
}

func (s *IntSet) Clear() {
	s.words = []uint{}
}

func(s *IntSet) Copy() *IntSet {
	var ret IntSet
	ret.words = make([]uint, len(s.words))
	copy(ret.words, s.words)
	return &ret
}

func (s *IntSet) AddAll(nums ...int) {
	for _, num := range nums {
		s.Add(num)
	}
}

func (s *IntSet) Elems() []int {
	ans := make([]int, 0)
	for word, bitset := range s.words {
		for bit := 0; bit < BitsetSize; bit++ {
			if bitset & (1 << bit) != 0 {
				ans = append(ans, word *BitsetSize+ bit)
			}
		}
	}
	return ans
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(&x)

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())

	fmt.Println(x.Has(9))

	x.Remove(9)
	fmt.Println(x.Has(9))

	t := x.Copy()
	t.AddAll(9, 10, 9, 11)
	t.AddAll()
	fmt.Println("t: ", t)

	x.AddAll(12, 13)
	fmt.Println("x: ", &x)
	copy1 := x.Copy()
	copy1.IntersectWith(t)
	fmt.Println("Intersection", copy1)

	copy2 := x.Copy()
	copy2.SymmetricDifference(t)
	fmt.Println("Symmetric Difference", copy2)

	copy3 := x.Copy()
	copy3.DifferenceWith(t)
	fmt.Println("x Difference With t", copy3)

	fmt.Println(x.Elems())

}
