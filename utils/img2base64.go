package utils

import (
	"encoding/base64"
	"io"
	"os"
)

func Img2Base64Url(imgPath string) (string, error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	imgBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	imgBase64 := base64.StdEncoding.EncodeToString(imgBytes)
	return "data:image/png;base64," + imgBase64, nil
}
