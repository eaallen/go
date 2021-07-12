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
	batch         string
	template_name string
	path          string
	name          string
}

type Config struct {
	Template map[string]map[string]string `json:"template"`
	Batch    map[string][]string          `json:"batch"`
}

func init() {
	flag.BoolVar(&options.help, "help", false, "Enter the number of posts you would love to see")
	flag.StringVar(&options.batch, "batch", "", "expect many templates")
	flag.StringVar(&options.template_name, "template", "", "name of the template you want generated")
	flag.StringVar(&options.path, "path", "", "path other than working directory, if non empty this path will be used instead of working directory")
	flag.StringVar(&options.name, "name", "", "custom name for genrated file")
	flag.Parse()
}

func main() {
	if options.help {
		Usage()
		return
	} else {
		var used_flag string
		config := GetConfig()
		templates := []map[string]string{}

		// handle getting template data
		if options.batch != "" {
			used_flag = "batch"
			// get batch data
			for _, el := range config.Batch[options.batch] {
				templates = append(templates, config.Template[el])
			}
		} else {
			used_flag = "template"
			templates = append(templates, config.Template[options.template_name])
		}

		// if tempaltes are empty then the user did not enter a valid key
		if len(templates) == 0 {
			color.Red("Error: " + used_flag + " key not found!")
			return
		}

		// create template for each template data
		for _, template_data := range templates {
			// eventualy direct path to template location
			data, err := ioutil.ReadFile(
				"C:/go/github.com/eaallen/go/codegen/templates/" + template_data["path"],
			)
			Check(err)
			name := GetName(options.name, template_data["default_name"])
			path := GetPath(options.path)
			fullpath := path + GetEnviromentSlash() + name + template_data["ext"]
			err = WriteToFile(fullpath, string(data)) // create the template
			Check(err)
			color.Green("Successfully built " + name + " at " + fullpath)
		}
		color.Green("Finished")
		return
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

func GetConfig() Config {
	// Open our jsonFile
	jsonFile, err := os.Open("C:/go/github.com/eaallen/go/codegen/templates/config.json")
	Check(err)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	data := UnmarshallJSON(jsonFile)

	return data
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

func UnmarshallJSON(jsonFile *os.File) Config {
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var config Config

	// we unmarshal our byteArray which contains our
	// jsonFile's content
	json.Unmarshal(byteValue, &config)

	return config
}
