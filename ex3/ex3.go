package main

import "fmt"

func main() {
	var f float64 = 10
	var n int = int(f)
	println(n)

	var sum int
	sum = 5 + 6 + 3
	avg := sum / 3
	if float64(avg) > 4.5 {
		println("good")
	}

	p := struct {
		name string
		age int
	}{ name: "gopher", age: 10 }

	p.age++
	println(p.name, p.age)

	ns := [...]int{1, 2, 3, 4, 5}
	println(ns[3])
	println(len(ns))
	fmt.Println(ns[1:3])

	ns3 := []int{5:50, 10:100}
	println(ns3)
}
