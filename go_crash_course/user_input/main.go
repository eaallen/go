package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, HandleLineBreak(), "", -1)
		fmt.Println(text)
		if text == "hi" {
			fmt.Println("hello, Yourself")
		}
	}
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func HandleLineBreak() string {
	if IsWindows() {
		return "\r\n"
	}
	return "\n"
}
