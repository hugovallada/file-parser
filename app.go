package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	psr "com.github.hugovallada/text-parser/src/parser"
)

func timedExecution(fn func()) {
	startTime := time.Now()
	fn()
	fmt.Println("Duration:", time.Since(startTime))
}

func main() {
	timedExecution(func() {
		fileName, newFileName, replaces, deleteOldFiles := psr.ParseArgs()

		for index, file := range fileName {
			executeFileParser(file, newFileName[index], replaces, deleteOldFiles)
		}
	})
}

func executeFileParser(fileName string, newFileName string, replaces map[string]string, deleteFile bool) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	newFile, err := os.Create(newFileName)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	for fileScanner.Scan() {
		newFile.WriteString(formatText(fileScanner.Text(), replaces) + "\n")
	}
	if deleteFile {
		os.Remove(fileName)
	}
}

func formatText(text string, replaces map[string]string) string {
	if len(replaces) >= 1 {
		for older, new := range replaces {
			text = strings.ReplaceAll(text, older, new)
		}
	}
	return strings.TrimSpace(text)
}
