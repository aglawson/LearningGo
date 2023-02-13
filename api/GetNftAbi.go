package api

import (
	"fmt"
	"io"
	"os"
)

func getNftAbi() string {
	file, err := os.Open("contract.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err.Error()
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	output := make([]byte, fileSize)
	_, err = io.ReadFull(file, output)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err.Error()
	}

	return string(output)
}
