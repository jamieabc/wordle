// Wordle
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("list.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close();

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
