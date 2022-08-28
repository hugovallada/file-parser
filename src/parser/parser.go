package parser

import (
	"flag"
	"strings"
)

func ParseArgs() ([]string, []string, map[string]string, bool) {
	fileNames := flag.String("fileNames", "", "Names of the files to be parsed.")
	newFileNames := flag.String("newFileNames", "", "Names of the new files to be created")
	parsers := flag.String("parsers", "", "Values to be parsed.")
	deleteOldFiles := flag.Bool("deleteOld", false, "Delete old files")

	flag.Parse()

	return parseStringToSlice(*fileNames), parseStringToSlice(*newFileNames), generateMapOfParsers(*parsers), *deleteOldFiles
}

func parseStringToSlice(stringToBeParsed string) []string {
	var parsedString []string
	switch {
	case strings.Contains(stringToBeParsed, ","):
		parsedString = append(parsedString, strings.Split(stringToBeParsed, ",")...)

	case strings.Contains(stringToBeParsed, " "):
		parsedString = append(parsedString, strings.Split(stringToBeParsed, " ")...)
	default:
		parsedString = append(parsedString, stringToBeParsed)
	}
	return parsedString
}

func generateMapOfParsers(stringToBeParsed string) map[string]string {
	mapOfParsers := make(map[string]string)
	for _, parser := range parseStringToSlice(stringToBeParsed) {
		values := strings.Split(parser, "=")
		if values[1] == "spc" {
			values[1] = " "
		}
		mapOfParsers[values[0]] = values[1]
	}
	return mapOfParsers
}
