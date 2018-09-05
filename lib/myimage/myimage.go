package myimage

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

//获取图片的宽高
func GetImgWidthHeight(imagePath string) (int, int) {
	file, _ := os.Open(imagePath)
	c, _, err := image.DecodeConfig(file)
	if err != nil {
		return 100, 100
	}
	file.Close()
	return c.Width, c.Height
}
