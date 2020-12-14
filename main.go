package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []string{"Alice", "Bob", "Dave"}
	fruits := []string{"Banana", "Orange", "Apple"}
	ex := []string{"bnklbnkl", "vjff", "kgf"}

	fmt.Println(sortByLen(people))
	sortByLenAndPrint(fruits)
	sortByLenAndPrint(ex)
}

// sort and return sorted data
func sortByLen(list []string) []string {
	sort.Slice(list, func(i, j int) bool {
		return len(list[i]) < len(list[j])
	})
	return list
}

// Print on screen sorted items
func sortByLenAndPrint(list []string) {
	fmt.Println(sortByLen(list))
}
