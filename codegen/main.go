package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"

	color "github.com/fatih/color"
	flag "github.com/spf13/pflag"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

var options struct {
	help          bool
	template_name string
	path          string
	name          string
}

func init() {
	flag.BoolVar(&options.help, "help", false, "Enter the number of posts you would love to see")
	flag.StringVar(&options.template_name, "template", "", "name of the template you want generated")
	flag.StringVar(&options.path, "path", "", "path other than working directory, if non empty this path will be used instead of working directory")
	flag.StringVar(&options.name, "name", "", "custom name for genrated file")
	flag.Parse()
}

func main() {
	if options.help {
		Usage()
	} else {
		template_data := GetConfig()[options.template_name]
		// eventualy direct path to template location
		data, err := ioutil.ReadFile("C:/go/github.com/eaallen/go/codegen/templates/" + template_data["path"])
		Check(err)
		name := GetName(options.name, template_data["default_name"])
		path := GetPath(options.path)
		fullpath := path + GetEnviromentSlash() + name + template_data["ext"]
		err = WriteToFile(fullpath, string(data))
		Check(err)
		color.Green("Successfully built " + name + " at " + fullpath)
	}
}

// get the right path if user inputs a path or not
func GetPath(user_path string) string {
	if user_path == "" {
		path, err := os.Getwd() // get working dir
		Check(err)
		return path
	}
	return user_path
}

func GetName(user_name string, default_name string) string {
	if user_name == "" {
		return default_name
	}
	return user_name
}

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	Check(err)
	defer file.Close()
	_, err = io.WriteString(file, data)
	Check(err)
	return file.Sync()
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s[post_count]\n", os.Args[0])
	flag.PrintDefaults()
}

func GetConfig() map[string]map[string]string {
	// Open our jsonFile
	jsonFile, err := os.Open("./templates/config.json")
	Check(err)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var res map[string]map[string]string
	json.Unmarshal([]byte(byteValue), &res)

	return res
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func GetEnviromentSlash() string {
	if IsWindows() {
		return "\\"
	}
	return "/"
}
