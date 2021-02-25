package peekalink

// LinkPreviewRequest contains the JSON body that should be sent when requesting
// a link preview (POST /)
type LinkPreviewRequest struct {
	Link string `json:"link"`
}

// LinkAvailabilityRequest contains the JSON body that should be sent when
// requesting a link availability check (POST /is-available)
type LinkAvailabilityRequest struct {
	Link string `json:"link"`
}
