package main

import (
	"fmt"
	"mymath"
)

func main()  {
	c, error := mymath.Add(1, -2)
	fmt.Println(c, error)
}
