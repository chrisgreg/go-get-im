package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SubredditApiResponse struct {
	Data struct {
		Children []SubredditChildren
	}
}

type SubredditChildren struct {
	Data struct {
		Title     string `json:"title"`
		Url       string `json:"url"`
		IsAlbum   bool
		ImgurHash string
	}
}

func parseRedditJson(body []byte) (*SubredditApiResponse, error) {
	var s = new(SubredditApiResponse)
	err := json.Unmarshal(body, &s)

	if err != nil {
		fmt.Println(err.Error())
	}
	return s, err
}

func GetRedditPosts(url string) *SubredditApiResponse {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "go-get-im")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	subredditJson, err := parseRedditJson(body)
	if err != nil {
		panic(err.Error())
	}

	return subredditJson
}
