package main

func is_even(num int) bool {
	return num % 2 == 0
}

func main()  {
	for i := 1; i <= 100; i++ {
		print(i)
		if is_even(i) {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
}
