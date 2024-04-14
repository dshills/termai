package main

import (
	"bufio"
	"os"
	"strings"
)

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func getPipedData() string {
	if !isInputFromPipe() {
		return ""
	}
	builder := strings.Builder{}
	r := os.Stdin
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}
	return builder.String()
}
