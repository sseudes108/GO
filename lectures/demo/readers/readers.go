package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	sum := 0

	for {
		input, inputErr := r.ReadString(' ')

		n := strings.TrimSpace(input)
		if n == "" {
			continue
		}

		num, convError := strconv.Atoi(n)
		if convError != nil {
			fmt.Println(convError)
		} else {
			sum += num
		}

		if inputErr == io.EOF {
			break
		}

		if inputErr != nil {
			fmt.Println("Error reading Stdin: ", inputErr)
		}
	}
	fmt.Println(sum)
}
