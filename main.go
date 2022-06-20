package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("./all.txt")
	if err != nil {
		panic("A big error here my guy")
	}
	fileContent := string(file)
	fileLines := strings.Split(fileContent, "\n")

	for _, line := range fileLines {
		ioutil.WriteFile(
			fmt.Sprintf("./lib/html/elements/%s", line),
			func() []byte {
				f, err := ioutil.ReadFile("./all.txt")
				if err != nil {
					panic("Ai muchacho")
				}
				fContent := string(f)
				fContent = strings.ReplaceAll(fContent, "#", strings.ToLower(line))
				fContent = strings.ReplaceAll(fContent, "$", strings.ToUpper(line))
				return []byte(fContent)
			}(),
			fs.ModeAppend,
		)
	}
}
