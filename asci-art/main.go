package main

import (
	"fmt"
	"os"
)

func main() {
	lent := len(os.Args)
	if lent < 2 || lent > 3 {
		fmt.Println("use: go run . os.Args[1] os.Args[2]")
		return
	}
	 banners := "standard.txt"

	fmt.Println(generate(banners))
}
