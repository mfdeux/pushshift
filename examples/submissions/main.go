package main

import (
	"fmt"
	"log"

	"github.com/mfdeux/pushshift"
)

func main() {
	client := pushshift.NewClient("testClient/0.1.0")
	q := &pushshift.SubmissionQuery{Author: "GallowBoob"}
	submissions, err := client.GetSubmissions(q)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(submissions)
}
