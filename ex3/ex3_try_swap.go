package main

func swap(x, y int)(int, int) {
	return y, x
}

func swap2(x, y *int)() {
	*x, *y = *y, *x
}

func main() {
	n, m := swap(10, 20)
	println(n, m)
	o, p := 10, 20
	swap2(&o, &p)
	println(o, p)
}
