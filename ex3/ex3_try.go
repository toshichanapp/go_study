package main

func main() {
	n1 := 19
	n2 := 86
	n3 := 1
	n4 := 12

	sum := n1 + n2 + n3 + n4
	println(sum)

	ns := []int{19, 86, 1, 12}
	var sum1 int
	for i:=0; i < len(ns); i++ {
		sum1 += ns[i]
	}

	println(sum1)
}
