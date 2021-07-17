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

// -------------------------------- Constants -------------------------
const TEMPLATE_PATH = "C:/go/github.com/eaallen/go/codegen/templates/"
const CONFIG_PATH = TEMPLATE_PATH + "config.json"

var options struct {
	help           bool
	show_templates bool
	batch          string
	template_name  string
	path           string
	name           string
}

type Config struct {
	Template map[string]map[string]string `json:"template"`
	Batch    map[string][]string          `json:"batch"`
}

// init() is called before main()
func init() {
	flag.BoolVar(&options.help, "help", false, "Enter the number of posts you would love to see")
	flag.BoolVar(&options.show_templates, "show", false, "Shows the templates you have avaliable")
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
	} else if options.show_templates {
		ShowTemplateOptions()
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
			// we know its a template
			var template_name string
			if options.template_name == "" {
				template_name = os.Args[1]
			} else {
				template_name = options.template_name
			}

			used_flag = "template"
			templates = append(templates, config.Template[template_name])
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
				TEMPLATE_PATH + template_data["path"],
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
	fmt.Fprintf(os.Stderr, "Usage: %s\n", os.Args[0])
	flag.PrintDefaults()
}

func GetConfig() Config {
	// Open our jsonFile
	jsonFile, err := os.Open(CONFIG_PATH)
	Check(err)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	data := UnmarshalJSON(jsonFile)

	return data
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetEnviromentSlash() string {
	if IsWindows() {
		return "\\"
	}
	return "/"
}

func UnmarshalJSON(jsonFile *os.File) Config {
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var config Config

	// we unmarshal our byteArray which contains our
	// jsonFile's content
	json.Unmarshal(byteValue, &config)

	return config
}

func ShowTemplateOptions() {
	const t = `Key (use to generate): %s 
Default Name: %s
Extension: %s
Location: %s
	
`
	config := GetConfig()
	// str := ""
	for key, item := range config.Template {
		fmt.Printf(t, key, item["default_name"], item["ext"], item["path"])
	}
}
