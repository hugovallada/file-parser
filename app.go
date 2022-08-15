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

func validateArgs(argsSize int) {
	if argsSize%2 == 0 {
		panic("Os replaces devem ser enviados em pares")
	}
}

func getArgs() map[string]string {
	replaces := make(map[string]string)
	if len(os.Args) > 1 {

		validateArgs(len(os.Args))

		for i := 1; i < len(os.Args); i += 2 {
			replaces[os.Args[i]] = os.Args[i+1]
		}
	}
	return replaces
}

func main() {
	executor(func() {
		replaces := getArgs()
		file, err := os.OpenFile("acesso.txt", os.O_CREATE, os.ModeAppend)

		if err != nil {
			return
		}

		defer file.Close()

		fileScanner := bufio.NewScanner(file)

		fileScanner.Split(bufio.ScanLines)

		newFile, err := os.Create("lista_acessos.txt")

		if err != nil {
			return
		}

		defer newFile.Close()

		for fileScanner.Scan() {
			newFile.WriteString(formatText(fileScanner.Text(), replaces) + "\n")
		}
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
