package utils

import (
	"io"
	"net/http"
	"os"
	"regexp"
)

func CleanUrl(url string) string {
	regex := regexp.MustCompile(".jpg|.png|.gif|.gifv")
	return regex.ReplaceAllString(url, "")
}

func DownloadFile(filepath string, url string) {
	out, err := os.Create(filepath)
	if err != nil {
		panic(err.Error())
	}

	defer out.Close()

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		panic(err.Error())
	}
}

func CreateFolder(username string) {
	os.MkdirAll("go-get-im-results/"+username, os.ModePerm)
}
