package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	words := strings.Split(input, " ")

	//Можно также использовать stringbuilder и записывать слова с конца

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	reversed := strings.Join(words, " ")

	fmt.Println(reversed)
}
