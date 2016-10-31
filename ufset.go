// Author: Vinhthuy Phan
// 11/2016

package ufset

type Element struct {
	Parent *Element
	Value  interface{}
	Rank   int
}

func Makeset(v int) *Element {
	e := new(Element)
	e.Rank = 0
	e.Parent = e
	e.Value = v
	return e
}

func Find(e *Element) *Element {
	if e.Parent != e {
		e.Parent = Find(e.Parent)
	}
	return e.Parent
}

func Union(e1, e2 *Element) {
	r1 := Find(e1)
	r2 := Find(e2)
	if r1 != r2 {
		if r1.Rank < r2.Rank {
			r1.Parent = r2
		} else if r2.Rank < r1.Rank {
			r2.Parent = r1
		} else {
			r1.Parent = r2
			r2.Rank++
		}
	}
}
