package input

import (
	"embed"
	"fmt"
	"strconv"
)

//go:embed day*.txt
var content embed.FS

func GetFileContents(day int) string {
	inputFile := fmt.Sprintf("day%d.txt", day)

	fileBytes, err := content.ReadFile(inputFile)
	if err != nil {
		panic(err.Error())
	}

	return string(fileBytes)
}

func GetInt(input string) int {
	return int(GetInt64(input))
}

func GetInt64(input string) int64 {
	n, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		panic(err)
	}

	return n
}

func GetFloat64(input string) float64 {
	return float64(GetInt64(input))
}
