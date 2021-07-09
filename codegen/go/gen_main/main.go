package main

import (
	"io"
	"io/ioutil"
	"os"

	flag "github.com/spf13/pflag"
)

// flags
var (
	user string
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}

func main() {
	template_path := "codegen/go/templates/"
	data, err := ioutil.ReadFile(template_path + "main.go")
	check(err)

	err = WriteToFile("codegen/go/main.go", string(data))
	check(err)
}

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	check(err)
	defer file.Close()
	_, err = io.WriteString(file, data)
	check(err)
	return file.Sync()
}
