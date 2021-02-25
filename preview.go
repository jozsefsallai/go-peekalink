package peekalink

import "fmt"

// LinkPreview is the response of a Peekalink link preview request.
type LinkPreview struct {
	URL              string            `json:"url"`
	Domain           string            `json:"domain"`
	LastUpdated      DateString        `json:"lastUpdated"`
	NextUpdate       DateString        `json:"nextUpdate"`
	ContentType      ContentType       `json:"contentType"`
	MimeType         string            `json:"mimeType"`
	Size             int64             `json:"size"`
	Redirected       bool              `json:"redirected"`
	RedirectionURL   string            `json:"redirectionUrl,omitempty"`
	RedirectionCount int64             `json:"redirectionCount,omitempty"`
	RedirectionTrail []string          `json:"redirectionTrail,omitempty"`
	Title            string            `json:"title,omitempty"`
	Description      string            `json:"description,omitempty"`
	Name             string            `json:"name"`
	TrackersDetected bool              `json:"trackersDetected"`
	Icon             ImageAssetDetails `json:"icon,omitempty"`
	Image            ImageAssetDetails `json:"image,omitempty"`
	Details          *LinkDetails      `json:"details,omitempty"`
}

// IsYouTubeLink checks whether the link is a YouTube link that contains
// YouTube-specific additional details.
func (lp *LinkPreview) IsYouTubeLink() bool {
	return lp.Details != nil && lp.Details.Type == LDTYouTube
}

// IsTwitterLink checks whether the link is a tweet link that contains
// Twitter-specific additional details.
func (lp *LinkPreview) IsTwitterLink() bool {
	return lp.Details != nil && lp.Details.Type == LDTTwitter
}

// GetYouTubeDetails will attempt to retrieve the additional details of a
// YouTube link.
func (lp *LinkPreview) GetYouTubeDetails() (*YouTubeLinkDetails, error) {
	if lp.Details == nil || !lp.IsYouTubeLink() {
		return nil, fmt.Errorf("the provided link is not a YouTube link")
	}

	return lp.Details.YouTubeLinkDetails, nil
}

// GetTwitterDetails will attempt to retrieve the additional details of a
// Twitter link.
func (lp *LinkPreview) GetTwitterDetails() (*TwitterLinkDetails, error) {
	if lp.Details == nil || !lp.IsTwitterLink() {
		return nil, fmt.Errorf("the provided link is not a Twitter link")
	}

	return lp.Details.TwitterLinkDetails, nil
}
