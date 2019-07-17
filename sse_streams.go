package pushshift

import (
	"encoding/json"
	"fmt"

	"github.com/r3labs/sse"
)

// url = "stream.pushshift.io"
// from: https://github.com/pushshift/reddit_sse_stream

type FirehoseQuery struct {
	Type               string `url:"type"`
	Author             string `url:"author"`
	Domain             string `url:"domain"`
	Subreddit          string `url:"subreddit"`
	SubmissionBackfill int    `url:"submission_backfil"`
	CommentBackfill    int    `url:"comment_backfill"`
	SubmissionStartID  int    `url:"submission_start_id"`
	CommentStartID     int    `url:"comment_start_id"`
	IsOver18           bool   `url:"over_18"`
	IsSelf             bool   `url:"is_self"`
	Filter             string `url:"filter"`
}

func (c *Client) StreamFirehose(q *FirehoseQuery) chan interface{} {
	things := make(chan interface{})
	client := sse.NewClient(sseEndpoint)
	go func() {
		client.Subscribe("messages", func(event *sse.Event) {
			switch string(event.Event) {
			case "rc":
				comment := &Comment{}
				if err := json.Unmarshal(event.Data, &comment); err != nil {
					fmt.Println(err)
				}
				things <- comment
				break
			case "rs":
				submission := &Submission{}
				if err := json.Unmarshal(event.Data, &submission); err != nil {
					fmt.Println(err)
				}
				things <- submission
				break
			default:
				break
			}
		})
	}()
	return things
}

func (c *Client) StreamSubmissions(q *FirehoseQuery) chan *Submission {
	submissions := make(chan *Submission)
	client := sse.NewClient(sseEndpoint)
	go func() {
		client.Subscribe("messages", func(event *sse.Event) {
			switch string(event.Event) {
			case "rs":
				submission := &Submission{}
				if err := json.Unmarshal(event.Data, &submission); err != nil {
					fmt.Println(err)
				}
				submissions <- submission
				break
			default:
				break
			}
		})
	}()
	return submissions
}

func (c *Client) StreamComments(q *FirehoseQuery) chan *Comment {
	comments := make(chan *Comment)
	client := sse.NewClient(sseEndpoint)
	go func() {
		client.Subscribe("messages", func(event *sse.Event) {
			switch string(event.Event) {
			case "rc":
				comment := &Comment{}
				if err := json.Unmarshal(event.Data, &comment); err != nil {
					fmt.Println(err)
				}
				comments <- comment
				break
			default:
				break
			}
		})
	}()
	return comments
}
