// Author: Vinhthuy Phan
// 11/2016

package main

import (
	"fmt"
	"github.com/vtphan/ufset"
)

func main() {
	items := make([]*(ufset.Element), 0)
	for i := 0; i < 10; i++ {
		items = append(items, ufset.Makeset(i))
	}
	ufset.Union(items[0], items[1])
	ufset.Union(items[2], items[3])
	ufset.Union(items[4], items[5])
	ufset.Union(items[6], items[7])
	ufset.Union(items[8], items[9])
	ufset.Union(items[0], items[9])
	ufset.Union(items[0], items[9])
	ufset.Union(items[2], items[8])

	for i := range items {
		fmt.Println(i, ufset.Find(items[i]))
	}
}
