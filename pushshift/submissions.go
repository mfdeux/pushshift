package pushshift

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// SubmissionQuery is a query
type SubmissionQuery struct {
	Sort                string `url:"sort,omitempty" json:"sort,omitempty"`
	SortType            string `url:"sort_type,omitempty" json:"sort_type,omitempty"`
	After               int    `url:"after,omitempty" json:"after,omitempty"`
	Before              int    `url:"before,omitempty" json:"before,omitempty"`
	AfterID             int    `url:"after_id,omitempty" json:"after_id,omitempty"`
	BeforeID            int    `url:"before_id,omitempty" json:"before_id,omitempty"`
	CreatedUTC          int    `url:"created_utc,omitempty" json:"created_utc,omitempty"`
	Score               int    `url:"score,omitempty" json:"score,omitempty"`
	Gilded              int    `url:"gilded,omitempty" json:"gilded,omitempty"`
	Edited              bool   `url:"edited,omitempty" json:"edited,omitempty"`
	Author              string `url:"author,omitempty" json:"author,omitempty"`
	Subreddit           string `url:"subreddit,omitempty" json:"subreddit,omitempty"`
	Distinguished       string `url:"distinguished,omitempty" json:"distinguished,omitempty"`
	RetrievedOn         int    `url:"retrieved_on,omitempty" json:"retrieved_on,omitempty"`
	LastUpdated         int    `url:"last_updated,omitempty" json:"last_updated,omitempty"`
	Query               string `url:"q,omitempty" json:"q,omitempty"`
	ID                  int    `url:"id,omitempty" json:"id,omitempty"`
	Metadata            bool   `url:"metadata,omitempty" json:"metadata,omitempty"`
	Unique              string `url:"unique,omitempty" json:"unique,omitempty"`
	Pretty              bool   `url:"pretty,omitempty" json:"pretty,omitempty"`
	HTMLDecode          bool   `url:"html_decode,omitempty" json:"html_decode,omitempty"`
	Permalink           string `url:"permalink,omitempty" json:"permalink,omitempty"`
	IsUserRemoved       bool   `url:"user_removed,omitempty" json:"user_removed,omitempty"`
	IsModRemoved        bool   `url:"mod_removed,omitempty" json:"mod_removed,omitempty"`
	SubredditType       string `url:"subreddit_type,omitempty" json:"subreddit_type,omitempty"`
	AuthorFlairCSSClass string `url:"author_flair_css_class,omitempty" json:"author_flair_css_class,omitempty"`
	AuthorFlairText     string `url:"author_flair_text,omitempty" json:"author_flair_text,omitempty"`

	IsOver18          bool `url:"over_18,omitempty" json:"over_18,omitempty"`
	IsLocked          bool `url:"author_flair_text,omitempty" json:"author_flair_text,omitempty"`
	IsSpoiler         bool `url:"author_flair_text,omitempty" json:"author_flair_text,omitempty"`
	IsVideo           bool `url:"author_flair_text,omitempty" json:"author_flair_text,omitempty"`
	IsSelf            bool `url:"author_flair_text,omitempty" json:"author_flair_text,omitempty"`
	IsOriginalContent bool `url:"author_flair_text,omitempty" json:"author_flair_text,omitempty"`
	isRedditMedia     bool `url:"author_flair_text,omitempty" json:"author_flair_text,omitempty"`
}

// Submission is a holder for submission data

type SubmissionGildings struct {
}
type Resolutions struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}
type Source struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}
type Variants struct {
}
type Images struct {
	ID          string        `json:"id"`
	Resolutions []Resolutions `json:"resolutions"`
	Source      Source        `json:"source"`
	Variants    Variants      `json:"variants"`
}
type Preview struct {
	Enabled bool     `json:"enabled"`
	Images  []Images `json:"images"`
}
type Submission struct {
	AllAwardings               []interface{}      `json:"all_awardings"`
	AllowLiveComments          bool               `json:"allow_live_comments"`
	Author                     string             `json:"author"`
	AuthorFlairCSSClass        interface{}        `json:"author_flair_css_class"`
	AuthorFlairRichtext        []interface{}      `json:"author_flair_richtext"`
	AuthorFlairText            interface{}        `json:"author_flair_text"`
	AuthorFlairType            string             `json:"author_flair_type"`
	AuthorFullname             string             `json:"author_fullname"`
	AuthorPatreonFlair         bool               `json:"author_patreon_flair"`
	CanModPost                 bool               `json:"can_mod_post"`
	ContestMode                bool               `json:"contest_mode"`
	CreatedUtc                 int                `json:"created_utc"`
	Domain                     string             `json:"domain"`
	FullLink                   string             `json:"full_link"`
	Gildings                   SubmissionGildings `json:"gildings"`
	ID                         string             `json:"id"`
	IsCrosspostable            bool               `json:"is_crosspostable"`
	IsMeta                     bool               `json:"is_meta"`
	IsOriginalContent          bool               `json:"is_original_content"`
	IsRedditMediaDomain        bool               `json:"is_reddit_media_domain"`
	IsRobotIndexable           bool               `json:"is_robot_indexable"`
	IsSelf                     bool               `json:"is_self"`
	IsVideo                    bool               `json:"is_video"`
	LinkFlairBackgroundColor   string             `json:"link_flair_background_color"`
	LinkFlairRichtext          []interface{}      `json:"link_flair_richtext"`
	LinkFlairTextColor         string             `json:"link_flair_text_color"`
	LinkFlairType              string             `json:"link_flair_type"`
	Locked                     bool               `json:"locked"`
	MediaOnly                  bool               `json:"media_only"`
	NoFollow                   bool               `json:"no_follow"`
	NumComments                int                `json:"num_comments"`
	NumCrossposts              int                `json:"num_crossposts"`
	Over18                     bool               `json:"over_18"`
	Permalink                  string             `json:"permalink"`
	Pinned                     bool               `json:"pinned"`
	RetrievedOn                int                `json:"retrieved_on"`
	Score                      int                `json:"score"`
	Selftext                   string             `json:"selftext"`
	SendReplies                bool               `json:"send_replies"`
	Spoiler                    bool               `json:"spoiler"`
	Stickied                   bool               `json:"stickied"`
	Subreddit                  string             `json:"subreddit"`
	SubredditID                string             `json:"subreddit_id"`
	SubredditSubscribers       int                `json:"subreddit_subscribers"`
	SubredditType              string             `json:"subreddit_type"`
	Thumbnail                  string             `json:"thumbnail"`
	Title                      string             `json:"title"`
	TotalAwardsReceived        int                `json:"total_awards_received"`
	URL                        string             `json:"url"`
	ParentWhitelistStatus      string             `json:"parent_whitelist_status,omitempty"`
	Pwls                       int                `json:"pwls,omitempty"`
	WhitelistStatus            string             `json:"whitelist_status,omitempty"`
	Wls                        int                `json:"wls,omitempty"`
	PostHint                   string             `json:"post_hint,omitempty"`
	Preview                    Preview            `json:"preview,omitempty"`
	SuggestedSort              string             `json:"suggested_sort,omitempty"`
	ThumbnailHeight            int                `json:"thumbnail_height,omitempty"`
	ThumbnailWidth             int                `json:"thumbnail_width,omitempty"`
	LinkFlairCSSClass          string             `json:"link_flair_css_class,omitempty"`
	LinkFlairTemplateID        string             `json:"link_flair_template_id,omitempty"`
	LinkFlairText              string             `json:"link_flair_text,omitempty"`
	AuthorFlairBackgroundColor string             `json:"author_flair_background_color,omitempty"`
	AuthorFlairTextColor       string             `json:"author_flair_text_color,omitempty"`
	ContentCategories          []string           `json:"content_categories,omitempty"`
}

// GetSubmissions queries the API for submissions
func (c *Client) GetSubmissions(q *SubmissionQuery) ([]Submission, error) {
	v, err := query.Values(q)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?%s", submissionEndpoint, v.Encode())
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

	var result submissionListing
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	var submissions []Submission
	for _, submission := range result.Data {
		submissions = append(submissions, submission)
	}

	return submissions, nil
}

type submissionListing struct {
	Data []Submission `json:"data"`
}
