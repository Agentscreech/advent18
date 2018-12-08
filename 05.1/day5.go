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

	badMap := make([]string, 0)

	for _, char := range alphabet {
		badMap = append(
			badMap,
			strings.ToUpper(string(char))+string(char),
			string(char)+strings.ToUpper(string(char)),
		)
	}

	for {
		didReplace := false
		for _, letters := range badMap {
			replaced := strings.Replace(str, letters, "", -1)
			if str != replaced {
				didReplace = true
				str = replaced
			}
		}
		if !didReplace {
			break
		}
	}

	fmt.Println(len(str))

}
