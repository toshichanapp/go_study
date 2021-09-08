package main

import (
	"math/rand"
	"time"
)

func randNum() int {
	t := time.Now().UnixNano()
	rand.Seed(t)

	return rand.Intn(6)
}

func main(){
	switch randNum() {
	case 6:
		println("大吉")
	case 5, 4:
		println("中吉")
	case 3, 2:
		println("吉")
	case 1:
		println("凶")
	}
}
