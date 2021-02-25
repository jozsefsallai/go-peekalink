package peekalink

// LinkAvailability contains the response to an is-available Peekalink reequst.
type LinkAvailability struct {
	IsAvailable bool `json:"isAvailable"`
}
