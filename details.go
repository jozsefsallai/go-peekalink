package peekalink

// YouTubeLinkDetails contains additional information about a YouTube link.
type YouTubeLinkDetails struct {
	VideoID        string `json:"videoId,omitempty"`
	Duration       string `json:"duration,omitempty"`
	YTViewCount    int64  `json:"viewCount,omitempty"`
	YTLikeCount    int64  `json:"likeCount,omitempty"`
	YTDislikeCount int64  `json:"dislikeCount,omitempty"`
	YTCommentCount int64  `json:"commentCount,omitempty"`
}

// TwitterLinkDetails contains additional information about a tweet.
type TwitterLinkDetails struct {
	StatusID          string `json:"statusId,omitempty"`
	RetweetCount      int64  `json:"retweetCount,omitempty"`
	TwitterLikesCount int64  `json:"likesCount,omitempty"`
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

	PublishedAt DateString `json:"publishedAt,omitempty"`
}
