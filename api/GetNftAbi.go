package api

import (
	"io"
	"os"
)

func getNftAbi() (string, error) {
	file, err := os.Open("contract.json")
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	output := make([]byte, fileSize)
	_, err = io.ReadFull(file, output)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
