package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/chrisgreg/go-get-im/api"
)

func IsAlbum(url string) bool {
	const albumStringRegex = ".*imgur.com/a/.*"
	matched, err := regexp.MatchString(albumStringRegex, url)
	if err != nil {
		panic(err.Error())
	}
	return matched
}

func GetAlbumImageLinks(post api.SubredditChildren) *api.ImgurAlbumApiResponse {
	imgurAPIUrl := api.CreateAlbumAPIUrl(post.Data.ImgurHash)
	imgurAPIRequest := api.CreateImgurRequest(imgurAPIUrl)

	client := &http.Client{}
	res, err := client.Do(imgurAPIRequest)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	albumData, err := api.ParseAlbumJson(body)
	if err != nil {
		panic(err.Error())
	}

	return albumData
}

func DownloadImages(posts []api.SubredditChildren, username string) {

	CreateFolder(username)

	for _, post := range posts {

		fmt.Println(post.Data)

		if post.Data.IsAlbum {

			albumResponse := GetAlbumImageLinks(post)

			for _, image := range albumResponse.Data {
				DownloadFile("go-get-im-results/"+username+"/"+image.ID+".jpg", image.Link)
			}
		} else {
			imageResponse := api.GetImageLink(post)
			if imageResponse.Data.Link == "" {
				continue
			}
			DownloadFile("go-get-im-results/"+username+"/"+imageResponse.Data.ID+".jpg", imageResponse.Data.Link)
		}
	}
}

func GetImgurHash(url string, album bool) string {
	const imageHashRegex = ".*imgur.com/(.*.)"
	const albumHashRegex = ".*imgur.com/a/(.*.)"

	var regexToMatch *regexp.Regexp
	var err error

	if album == true {
		regexToMatch, err = regexp.Compile(albumHashRegex)
	} else {
		regexToMatch, err = regexp.Compile(imageHashRegex)
	}

	if err != nil {
		panic(err.Error())
	}

	match := regexToMatch.FindStringSubmatch(url)
	return CleanUrl(match[1])
}
