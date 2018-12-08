package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../inputs/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := bufio.NewScanner(file)
	var str string
	for input.Scan() {
		str = input.Text()
	}

	alphabet := "abcdefghijklmnopqrstuvwxyz"

	badMap := make(map[string]bool)
	for _, L := range alphabet {
		l := string(L)
		Aa := strings.ToUpper(l) + l
		aA := l + strings.ToUpper(l)
		badMap[Aa] = true
		badMap[aA] = true
	}
	didReplace := true
	for {
		if didReplace {
			for i := range str {
				if i != len(str)-1 {
					if _, found := badMap[str[i:i+2]]; found {
						str = strings.Replace(str, str[i:i+2], "", -1)
						break
					}
				} else {
					didReplace = false

				}
			}

		} else {
			break
		}
	}

	fmt.Println(len(str))

}
