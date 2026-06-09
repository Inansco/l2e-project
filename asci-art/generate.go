package main

import (
	"fmt"
	"os"
	"strings"
)

func generate(input string) string {

	input = os.Args[1]

	banners := "standard.txt"
	if len(os.Args) == 3 {
		banners = os.Args[2]
	}

	file, err := os.ReadFile(banners)
	if err != nil {
		fmt.Println("Error")
		return ""
	}
	banner := strings.Split(string(file), "\n")

	line := strings.Split(input, "\\n")

	for _, char := range line {
		if char == "" {
			fmt.Println()
			continue
		}
		for i := 0; i <= 8; i++ {
			for j := 0; j < len(char); j++ {
				index := int(char[j]-32)*9 + i
				if index >= 0 && index < len(banner) {
					fmt.Print(banner[index])
				}

			}
			fmt.Println()
		}
	}
	return ""

}
