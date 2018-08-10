package utils

import (
	"fmt"
	"os"
)

func Write(filePath string, content string) {
	fout, err := os.Create(filePath)
	if err != nil {
		fmt.Println(filePath, err)
		return
	}
	defer fout.Close()
	fout.WriteString(content)
}
