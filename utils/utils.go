package utils

import (
	"bufio"
	"os"
)

func ReadFile(name string) []string {
	file, _ := os.Open(name)
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}
