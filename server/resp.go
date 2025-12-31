package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := "$5\r\nAhmed\r\n"
	reader := bufio.NewReader(strings.NewReader(input))

	// read the first character
	b, _ := reader.ReadByte()
	// only strings supported
	if b != '$' {
		fmt.Println("Invalid command, expecting bulk strings only")
		os.Exit(1)
	}

	// read the second character which is number, determine the number of character to read
	size, _ := reader.ReadByte()
	strSize, _ := strconv.Atoi(string(size))

	// \r
	reader.ReadByte()
	// \n
	reader.ReadByte()

	// read *strSize* characters
	name := make([]byte, strSize)
	reader.Read(name)
	fmt.Println(string(name))

}
