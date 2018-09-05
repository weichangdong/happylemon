package qrcode

import (
	"image"
	_ "image/png"
	"os"

	"github.com/divan/qrlogo"
)

func GenQRcodeImg(text string, qrImgName string, qrLogoFile string) (bool, string) {
	QrSize := 256
	file, err := os.Open(qrLogoFile)
	if err != nil {
		return false, "err-1"
	}
	defer file.Close()

	logo, _, err := image.Decode(file)
	if err != nil {
		return false, "err-2"
	}
	qr, err := qrlogo.Encode(text, logo, QrSize)
	if err != nil {
		return false, "err-3"
	}

	out, err := os.Create(qrImgName)
	if err != nil {
		return false, "err-4"
	}
	out.Write(qr.Bytes())
	defer out.Close()
	return true, qrImgName
}
