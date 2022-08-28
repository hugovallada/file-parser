package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func executor(fn func()) {
	startTime := time.Now()
	fn()
	fmt.Println("Duration:", time.Since(startTime))
}

func parseArgs() (string, string, map[string]string) {
	if len(os.Args) < 4 {
		panic("Not enough args")
	}
	fileName, newFileName, stringOfParsers := os.Args[1], os.Args[2], os.Args[3]
	return fileName, newFileName, getMapOfParsers(stringOfParsers)
}

func getMapOfParsers(stringOfParsers string) map[string]string {
	sliceOfParsers := strings.Split(stringOfParsers, ",")
	mapOfParsers := make(map[string]string)
	for _, parser := range sliceOfParsers {
		values := strings.Split(parser, "=")
		mapOfParsers[values[0]] = values[1]
	}
	return mapOfParsers
}

func main() {
	executor(func() {
		fileName, newFileName, replaces := parseArgs()
		file, err := os.OpenFile(fileName, os.O_CREATE, os.ModeAppend)
		if err != nil {
			return
		}
		defer file.Close()
		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanLines)
		newFile, err := os.Create(newFileName)
		if err != nil {
			return
		}
		defer newFile.Close()
		for fileScanner.Scan() {
			newFile.WriteString(formatText(fileScanner.Text(), replaces) + "\n")
		}
		os.Remove(fileName)
	})
}

func formatText(text string, replaces map[string]string) string {
	text = strings.TrimSpace(text)
	if len(replaces) >= 1 {
		for older, new := range replaces {
			text = strings.ReplaceAll(text, older, new)
		}
	}
	return text
}
