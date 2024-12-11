package utils

import (
	"os"
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

func ReadLines(fileName string) []string {
	lines := strings.Split(ReadFile(fileName), "\n")
	n := len(lines)
	return lines[0:n]
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
