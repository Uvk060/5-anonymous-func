package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []string{"Alice", "Bob", "Dave"}
	sort.Slice(people, func(i, j int) bool {
		return len(people[i]) < len(people[j])
	})

	fmt.Println(people)
}
