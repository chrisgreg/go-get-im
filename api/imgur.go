package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const ImgurClientID = "374afdc264065e7"

type ImgurAlbumApiResponse struct {
	Data    []ImgurAlbumChild
	Success bool
	Status  int32
}

type ImgurImageApiResponse struct {
	Data    ImgurAlbumChild
	Success bool
	Status  int32
}

type ImgurAlbumChild struct {
	Link string `json:"link"`
	ID   string `json:"id"`
}

func ParseAlbumJson(body []byte) (*ImgurAlbumApiResponse, error) {
	var s = new(ImgurAlbumApiResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s, err
}

func ParseImageJson(body []byte) (*ImgurImageApiResponse, error) {
	var s = new(ImgurImageApiResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println(err.Error())
	}
	return s, err
}

func GetImageLink(post SubredditChildren) *ImgurImageApiResponse {
	imgurAPIUrl := CreateImageAPIUrl(post.Data.ImgurHash)
	imgurAPIRequest := CreateImgurRequest(imgurAPIUrl)

	client := &http.Client{}
	res, err := client.Do(imgurAPIRequest)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	imageData, err := ParseImageJson(body)
	if err != nil {
		panic(err.Error())
	}

	return imageData
}

func CreateImgurRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err.Error())
	}

	authHeader := fmt.Sprintf("Client-ID %s", ImgurClientID)
	req.Header.Add("Authorization", authHeader)
	req.Header.Add("Content-Type", "application/json")
	return req
}

func CreateAlbumAPIUrl(hash string) string {
	const url = "https://api.imgur.com/3/album/%s/images"
	return fmt.Sprintf(url, hash)
}

func CreateImageAPIUrl(hash string) string {
	const url = "https://api.imgur.com/3/image/%s"
	return fmt.Sprintf(url, hash)
}
