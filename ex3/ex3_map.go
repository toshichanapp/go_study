package main

func main() {
	m := map[string]int{"x": 10, "y": 20}

	println(m["z"])
	n, ok := m["z"]
	println(n, ok)
}