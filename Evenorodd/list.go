package main

type ss []int

func newlist() ss {
	li := ss{}
	cardSuits := 10
	for i := 1; i <= cardSuits; i++ {
		li = append(li, i)
	}
	return li
}

func evenorodd(s int) string {
	if s%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
