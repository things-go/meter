package main

import (
	"fmt"

	"github.com/things-go/meter"
)

func main() {
	value, _ := meter.ParseBytes("2.99TB")
	s := meter.HumanSize(value)
	fmt.Println(value)
	fmt.Println(s)

	// output:
	// 3287539767050
	// 2.99TB
}
