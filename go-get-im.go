package main

import (
	"fmt"
	"os"

	"github.com/chrisgreg/go-get-im/utils"
)

func main() {

	userName := os.Args[1]
	userUrlToCheck := fmt.Sprintf("https://www.reddit.com/user/%s/submitted.json", userName)

	submittedPosts := utils.GetPosts(userUrlToCheck)

	utils.DownloadImages(submittedPosts, userName)
	fmt.Println("Finished")
}
