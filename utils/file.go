package utils

import (
	"os"
	"strings"
)

func ReadFile(fileName string) string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func ReadLines(fileName string) []string {
	lines := strings.Split(ReadFile(fileName), "\n")
	n := len(lines)
	return lines[0:n]
}
