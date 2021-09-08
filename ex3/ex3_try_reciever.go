package main

type MyInt int
func (n MyInt) Inc() MyInt {
	return MyInt(int(n) + 1)
}

func main()  {
	var n MyInt
	println(n)
	n.Inc()
	println(n.Inc())
}
