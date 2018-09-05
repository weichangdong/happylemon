package download

import (
	"io"
	"net/http"
	"os"
)

func DownFile(url string, saveFileName string) (bool, string) {
	res, err := http.Get(url)
	if err != nil {
		return false, "http-get-error:" + err.Error()
	}
	f, err := os.Create(saveFileName)
	if err != nil {
		return false, "os-create-error:" + err.Error()
	}
	io.Copy(f, res.Body)

	return true, saveFileName
}
