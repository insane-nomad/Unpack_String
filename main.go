package main

import (
	"fmt"
	"strconv"
)

var data = `a4bc2d5e`

func main() {
	unpack(data)
}

func IsLetter(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return false
	}
	return true
}

func unpack(target string) {
	if !IsLetter(target) {
		fmt.Println("")
		return
	}

	byteArray := make([]byte, 0, len(target)*2)
	compare := 0

	for i := 0; i < len(target); i++ {
		if int(target[i]-'0') <= 9 {
			if target[i-1] == '\\' && target[i-2] != '\\' {
				byteArray = append(byteArray, target[i])
			} else if target[i-1] == '\\' && target[i-2] == '\\' {
				compare = int(target[i] - '0')
			} else {
				compare = int(target[i] - '0' - 1)
			}
			for count := 0; count < compare; count++ {
				byteArray = append(byteArray, target[i-1])
			}
		} else {
			if target[i] != '\\' {
				byteArray = append(byteArray, target[i])
			}
		}
	}
	fmt.Printf("%v\n", string(byteArray))
}
