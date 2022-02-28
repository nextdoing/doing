package main

import (
	"fmt"
	"sort"
	"strings"

	"gobase/base"
)

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(base.Index(strs, "pear"))
	fmt.Println(base.Include(strs, "grape"))
	fmt.Println(base.Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "b")
	}))
	fmt.Println(base.All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(base.Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	fmt.Println(base.Map(strs, strings.ToUpper))

	si := []string{"apple", "pear", "aims"}
	sort.Sort(base.Str(si))
	fmt.Println("Sort str type: ", si)
}
