package main

import (
	"fmt"

	"github.com/mfdeux/pushshift/pushshift"
)

func main() {
	client := pushshift.NewClient("testClient/0.1.0")
	things := client.StreamFirehose(nil)
	for thing := range things {
		switch t := thing.(type) {
		case *pushshift.Comment:
			fmt.Println(t.Subreddit)
			fmt.Println(t.Author)
			fmt.Println(t.Body)
			break
		case *pushshift.Submission:
			fmt.Println(t.Subreddit)
			fmt.Println(t.Author)
			fmt.Println(t.URL)
			break
		default:
			fmt.Printf("I don't know about type %T!\n", t)
		}

	}
}
