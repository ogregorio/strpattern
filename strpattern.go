package main

import (
	"files"
	"os"
	"read"
	"replacer"
)

func main() {

	var origin string
	var destiny string
	var regex string

	args := os.Args

	if args[1] == "--help" {
		help()
	} else {
		switch args[1] {
		case "--origin":
			origin, destiny, regex = isValidCommand(args)
			if origin == "" {
				os.Exit(-1)
			} else {
				program(origin, destiny, regex)
			}
		default:
			println("Invalid command.")
		}
	}
}

func program(origin string, destiny string, regex string) {
	contentFromFile, _ := files.Read(origin)
	var words []string

	//for each line
	for i := 0; i < len(contentFromFile); i++ {
		//concat actual string vectors with generated string vector from line
		words = read.AppendLine(words, contentFromFile[i], regex)
	}

	println("Define your pattern:")
	pattern := read.Read()

	//Add to content the new line with pattern implemented
	var content string = ""
	for i := 0; i < len(words); i++ {
		content += replacer.Replace(pattern, words[i])
	}

	//Write on file
	files.WriteOneLine(content, destiny)
}

//Validate options and arguments
func isValidCommand(args []string) (string, string, string) {

	//If hasn't origin
	if len(args) == 0 {
		println("Invalid command!")
		return "", "", ""
		//If has just origin
	} else if args[1] == "--origin" && len(args) <= 3 {
		return args[2], args[2] + ".out", " "
		//If has just origin and destiny
	} else if args[1] == "--origin" && args[3] == "--destiny" && len(args) <= 5 {
		return args[2], args[4], " "
		//If has origin, destiny and regex
	} else if args[1] == "--origin" && args[3] == "--destiny" && args[5] == "--regex" {
		return args[2], args[4], args[6]
	}
	return "", "", ""
}

//Help
func help() {
	println("\nUSE:")
	println("strpattern --origin [path] --destiny [path] --regex [regex]\n")
	println("INFO:")
	println("The words found in your file will be allocated in the template defined by you, respecting the {?} icon.\n")
	println("EXAMPLE:")
	println("Origin: apple banana grape")
	println("Pattern: {?} is a fruit.")
	println("Destiny: \n    apple is a fruit\n    banana is a fruit\n    grape is a fruit")
}
