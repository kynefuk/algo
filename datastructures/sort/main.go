package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAgeAsc []Person

func (a ByAgeAsc) Len() int {
	return len(a)
}

func (a ByAgeAsc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAgeAsc) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type ByAgeDesc []Person

func (a ByAgeDesc) Len() int {
	return len(a)
}

func (a ByAgeDesc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAgeDesc) Less(i, j int) bool {
	return a[i].Age > a[j].Age
}

func main() {
	people := []Person{
		{"Bob", 30},
		{"John", 31},
		{"Michael", 32},
		{"Jenny", 33},
	}

	// ascending by sort.Interface
	sort.Sort(ByAgeAsc(people))
	fmt.Println(people)

	// descending by sort.Interface
	sort.Sort(ByAgeDesc(people))
	fmt.Println(people)

	// sort.Sliceはsort.Interfaceを満たさなくても使える
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})

	fmt.Println(people)
}
