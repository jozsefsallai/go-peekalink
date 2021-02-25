package peekalink

// YouTubeLinkDetails contains additional information about a YouTube link.
type YouTubeLinkDetails struct {
	VideoID      string     `json:"videoId,omitempty"`
	Duration     string     `json:"duration,omitempty"`
	ViewCount    int64      `json:"viewCount,omitempty"`
	LikeCount    int64      `json:"likeCount,omitempty"`
	DislikeCount int64      `json:"dislikeCount,omitempty"`
	CommentCount int64      `json:"commentCount,omitempty"`
	PublishedAt  DateString `json:"publishedAt,omitempty"`
}

// TwitterLinkDetails contains additional information about a tweet.
type TwitterLinkDetails struct {
	StatusID     string     `json:"statusId,omitempty"`
	RetweetCount int64      `json:"retweetCount,omitempty"`
	LikesCount   int64      `json:"likesCount,omitempty"`
	PublishedAt  DateString `json:"publishedAt,omitempty"`
}

// LinkDetailType specifies the kind of the information contained in the link
// preview's `Details` field.
type LinkDetailType string

const (
	// LDTYouTube means the link preview contains details for a YouTube link.
	LDTYouTube LinkDetailType = "youtube"

	// LDTTwitter means the link preview contains details for a Twitter link.
	LDTTwitter LinkDetailType = "twitter"
)

// LinkDetails contains additional information about a link. The kind of
// information is determined by the `Type` field.
type LinkDetails struct {
	Type LinkDetailType `json:"type"`

	*YouTubeLinkDetails
	*TwitterLinkDetails
}
