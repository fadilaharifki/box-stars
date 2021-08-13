package main

import "fmt"

func main() {

	result := ""

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result += "*"
		}
		fmt.Println(result)
		result = ""
	}
}
