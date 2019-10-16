package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

	//cuidado porque o slice pode substituir o valor do array
	a := [2]int{1, 2}
	fmt.Println(a)

	slice := a[:1]
	fmt.Println(slice)

	slice = append(slice, 4)
	fmt.Println(slice)
	fmt.Println(a)

}
