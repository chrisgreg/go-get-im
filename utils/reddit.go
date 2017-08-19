package utils

import (
	"strings"

	"github.com/chrisgreg/go-get-im/api"
)

func GetPosts(url string) []api.SubredditChildren {
	var results []api.SubredditChildren

	subredditJson := api.GetRedditPosts(url)

	for _, post := range subredditJson.Data.Children {

		if !strings.Contains(post.Data.Url, "imgur") {
			continue
		}

		post.Data.IsAlbum = IsAlbum(post.Data.Url)
		post.Data.ImgurHash = GetImgurHash(post.Data.Url, post.Data.IsAlbum)
		results = append(results, post)
	}

	return results
}
