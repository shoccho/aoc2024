package utils

import (
	"os"
	"runtime"
	"strings"
	"time"
)

func ReadFile(fileName string) string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func GetEndl() string {
	os := runtime.GOOS
	splitBy := "\n"
	if os == "windows" {
		splitBy = "\r\n"
	}
	return splitBy
}

func ReadLines(fileName string) []string {

	lines := strings.Split(ReadFile(fileName), GetEndl())
	n := len(lines)
	if lines[n-1] == "" {
		return lines[:n-1]
	}
	return lines
}

func MeasureAvgRuntime(fn func(), iterations int) time.Duration {
	var totalTime time.Duration

	for i := 0; i < iterations; i++ {
		start := time.Now()
		fn()
		elapsed := time.Since(start)
		totalTime += elapsed
	}
	return totalTime / time.Duration(iterations)
}
