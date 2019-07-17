package main

import (
	"fmt"

	"github.com/mfdeux/pushshift/pushshift"
)

func main() {
	client := pushshift.NewClient("testClient/0.1.0")
	submissions := client.StreamSubmissions(nil)
	for submission := range submissions {
		fmt.Println(submission.Subreddit)
		fmt.Println(submission.Author)
		fmt.Println(submission.URL)
	}
}
