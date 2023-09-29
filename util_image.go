package main

import (
	"net/http"
	"os"
	"strings"
)

func IsImage(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false, err
	}

	contentType := http.DetectContentType(buffer)
	if strings.Split(contentType, "/")[0] == "image" {
		return true, nil
	}

	return false, nil

}
