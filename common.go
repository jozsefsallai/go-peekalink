package peekalink

import "time"

// DateString contains an RFC3339-compliant datetime string.
type DateString string

// ToDate will attempt to turn the date string into a `Time` object.
func (ds DateString) ToDate() (time.Time, error) {
	return time.Parse(time.RFC3339, string(ds))
}

// ImageAssetDetails contains information about an image resource (i.e. image
// preview or favicon).
type ImageAssetDetails struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// ContentType specifies the type of the content that the link points to.
type ContentType string

const (
	// ContentTypeHTML represents an HTML content type.
	ContentTypeHTML ContentType = "html"

	// ContentTypeImage represents an image content type.
	ContentTypeImage ContentType = "image"

	// ContentTypeVideo represents a video content type.
	ContentTypeVideo ContentType = "video"

	// ContentTypeGIF represents a GIF content type.
	ContentTypeGIF ContentType = "gif"
)
