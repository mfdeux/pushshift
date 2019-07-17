package pushshift

import (
	"encoding/json"
	"fmt"

	"github.com/r3labs/sse"
)

// url = "stream.pushshift.io"
// from: https://github.com/pushshift/reddit_sse_stream

type FirehoseQuery struct {
	Type               string `url:"type,omitempty"`
	Author             string `url:"author,omitempty"`
	Domain             string `url:"domain,omitempty"`
	Subreddit          string `url:"subreddit,omitempty"`
	SubmissionBackfill int    `url:"submission_backfill,omitempty"`
	CommentBackfill    int    `url:"comment_backfill,omitempty"`
	SubmissionStartID  int    `url:"submission_start_id,omitempty"`
	CommentStartID     int    `url:"comment_start_id,omitempty"`
	IsOver18           bool   `url:"over_18,omitempty"`
	IsSelf             bool   `url:"is_self,omitempty"`
	Filter             string `url:"filter,omitempty"`
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
