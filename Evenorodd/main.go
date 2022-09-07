package main

import "fmt"

func main() {
	lis := newlist()
	for j := 1; j <= len(lis); j++ {
		fmt.Println(evenorodd(j))
	}
}
