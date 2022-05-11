package main

import "fmt"

func ExampleNestedIterator() {
	a := NestedInteger{}
	b := NestedInteger{}
	c := NestedInteger{}
	d := NestedInteger{}
	e := NestedInteger{}
	a.SetInteger(1)
	b.SetInteger(1)
	c.SetInteger(2)
	d.SetInteger(1)
	e.SetInteger(1)
	list1 := NestedInteger{}
	list1.Add(a)
	list1.Add(b)
	list2 := NestedInteger{}
	list2.Add(d)
	list2.Add(e)
	ni := Constructor([]*NestedInteger{&list1, &c, &list2})
	for ni.HasNext() {
		fmt.Println(ni.Next())
	}
	f := NestedInteger{}
	g := NestedInteger{}
	h := NestedInteger{}
	f.SetInteger(1)
	g.SetInteger(4)
	h.SetInteger(6)
	list3 := NestedInteger{}
	list3.Add(h)
	list4 := NestedInteger{}
	list4.Add(g)
	list4.Add(list3)
	ni = Constructor([]*NestedInteger{&f, &list4})
	for ni.HasNext() {
		fmt.Println(ni.Next())
	}
	// Output:
	// 1
	// 1
	// 2
	// 1
	// 1
	// 1
	// 4
	// 6
}
