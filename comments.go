package pushshift

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// CommentQuery is the query for comments
type CommentQuery struct {
	Sort                string `url:"sort"`
	SortType            string `url:"sort_type"`
	After               int    `url:"after"`
	Before              int    `url:"before"`
	AfterID             int    `url:"after_id"`
	BeforeID            int    `url:"before_id"`
	CreatedUTC          int    `url:"created_utc"`
	Score               int    `url:"score"`
	Gilded              int    `url:"gilded"`
	Edited              bool   `url:"edited"`
	Author              string `url:"author"`
	Subreddit           string `url:"subreddit"`
	Distinguished       string `url:"distinguished"`
	RetrievedOn         int    `url:"retrieved_on"`
	LastUpdated         int    `url:"last_updated"`
	Query               string `url:"q"`
	ID                  int    `url:"id"`
	Metadata            bool   `url:"metadata"`
	Unique              string `url:"unique"`
	Pretty              bool   `url:"pretty"`
	HTMLDecode          bool   `url:"html_decode"`
	Permalink           string `url:"permalink"`
	IsUserRemoved       bool   `url:"user_removed"`
	IsModRemoved        bool   `url:"mod_removed"`
	SubredditType       string `url:"subreddit_type"`
	AuthorFlairCSSClass string `url:"author_flair_css_class"`
	AuthorFlairText     string `url:"author_flair_text"`

	ReplyDelay    int `url:"reply_delay"`
	NestLevel     int `url:"nest_level"`
	SubReplyDelay int `url:"sub_reply_delay"`
	UTCHourofWeek int `url:"utc_hour_of_week"`
	LinkID        int `url:"link_id"`
	ParentID      int `url:"parent_id"`
}

type CommentGildings struct{}

// Comment is a holder for comment data
type Comment struct {
	AllAwardings               []interface{}   `json:"all_awardings"`
	ApprovedAtUtc              interface{}     `json:"approved_at_utc"`
	Author                     string          `json:"author"`
	AuthorFlairBackgroundColor interface{}     `json:"author_flair_background_color"`
	AuthorFlairCSSClass        interface{}     `json:"author_flair_css_class"`
	AuthorFlairRichtext        []interface{}   `json:"author_flair_richtext"`
	AuthorFlairTemplateID      interface{}     `json:"author_flair_template_id"`
	AuthorFlairText            interface{}     `json:"author_flair_text"`
	AuthorFlairTextColor       interface{}     `json:"author_flair_text_color"`
	AuthorFlairType            string          `json:"author_flair_type"`
	AuthorFullname             string          `json:"author_fullname"`
	AuthorPatreonFlair         bool            `json:"author_patreon_flair"`
	BannedAtUtc                interface{}     `json:"banned_at_utc"`
	Body                       string          `json:"body"`
	CanModPost                 bool            `json:"can_mod_post"`
	Collapsed                  bool            `json:"collapsed"`
	CollapsedReason            interface{}     `json:"collapsed_reason"`
	CreatedUtc                 int             `json:"created_utc"`
	Distinguished              interface{}     `json:"distinguished"`
	Edited                     bool            `json:"edited"`
	Gildings                   CommentGildings `json:"gildings"`
	ID                         string          `json:"id"`
	IsSubmitter                bool            `json:"is_submitter"`
	LinkID                     string          `json:"link_id"`
	Locked                     bool            `json:"locked"`
	NoFollow                   bool            `json:"no_follow"`
	ParentID                   string          `json:"parent_id"`
	Permalink                  string          `json:"permalink"`
	RetrievedOn                int             `json:"retrieved_on"`
	Score                      int             `json:"score"`
	SendReplies                bool            `json:"send_replies"`
	Stickied                   bool            `json:"stickied"`
	Subreddit                  string          `json:"subreddit"`
	SubredditID                string          `json:"subreddit_id"`
	TotalAwardsReceived        int             `json:"total_awards_received"`
}

// GetComments queries the API for comments
func (c *Client) GetComments(q *CommentQuery) ([]Comment, error) {
	v, err := query.Values(q)
	if err != nil {
		return nil, err
	}
	v.Encode()

	url := fmt.Sprintf("%s?%s", commentsEndpoint, v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.userAgent)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result commentsListing
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	var comments []Comment
	for _, comment := range result.Data {
		comments = append(comments, comment)
	}

	return comments, nil
}

type commentsListing struct {
	Data []Comment `json:"data"`
}
