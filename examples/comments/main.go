package main

import (
	"fmt"
	"log"

	"github.com/mfdeux/pushshift/pushshift"
)

func main() {
	client := pushshift.NewClient("testClient/0.1.0")
	q := &pushshift.CommentQuery{Author: "SlinkToTheDink"}
	comments, err := client.GetComments(q)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(comments)
}
