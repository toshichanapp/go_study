package main

import "fmt"

func main(){
	str := "hello, 世界"
	println(str)

	const strConst = "hello world"
	println(strConst)

	const (
		c = 1 << iota
		d
		e
		f
	)
	println(c, d, e, f)

	for i :=1; i <= 100; i++ {
		if i % 2 == 0 {
			fmt.Printf( "%v-偶数", i)
		} else {
			fmt.Printf( "%v-奇数", i)
		}
	}
}
