package qrcodeGen

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

var (
	folderLoc = "./codeimg/"
	filePrefix = "code_"
	extension = ".png"
)

func GenerateCode(id, content string) (string, error) {
	filename := filePrefix + id + extension
	if err := qrcode.WriteFile(
		content,
		qrcode.High,
		256,
		folderLoc + filename,
	); err != nil {
		return "", fmt.Errorf("GenerateCode, %w", err)
	}
	return filename ,nil
}
