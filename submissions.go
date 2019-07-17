package pushshift

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// SubmissionQuery is a query
type SubmissionQuery struct {
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

	IsOver18          bool `url:"over_18"`
	IsLocked          bool `url:"author_flair_text"`
	IsSpoiler         bool `url:"author_flair_text"`
	IsVideo           bool `url:"author_flair_text"`
	IsSelf            bool `url:"author_flair_text"`
	IsOriginalContent bool `url:"author_flair_text"`
	isRedditMedia     bool `url:"author_flair_text"`
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
