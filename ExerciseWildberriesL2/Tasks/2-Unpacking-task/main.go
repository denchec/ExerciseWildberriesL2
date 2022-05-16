package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var newString string
	fmt.Scanln(&newString)

	if newString == "" {
		fmt.Println(newString)
		return
	}

	var symbol string
	var endString string
	escapeOrder := false

	for _, i := range newString {
		if escapeOrder {
			if symbol == "" {
				symbol += string(i)
				escapeOrder = false
			} else {
				endString += symbol
				symbol = string(i)
				escapeOrder = false
			}
		} else {
			if num, err := strconv.Atoi(string(i)); err == nil {
				if symbol == "" {
					continue
				}
				endString += strings.Repeat(symbol, num)
				symbol = ""
				continue
			}
			if string(i) == "\\" {
				escapeOrder = true
				continue
			}
			if symbol == "" {
				symbol += string(i)
			} else {
				endString += symbol
				symbol = string(i)
			}
		}
	}
	endString += symbol

	fmt.Println(endString)
}
