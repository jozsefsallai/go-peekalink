package peekalink_test

import (
	"testing"

	"github.com/ankida/go-peekalink"
)

func TestLinkPreview(t *testing.T) {
	youtubeLinkData := peekalink.LinkPreview{
		Details: &peekalink.LinkDetails{
			Type:               peekalink.LDTYouTube,
			YouTubeLinkDetails: &peekalink.YouTubeLinkDetails{},
		},
	}

	twitterLinkData := peekalink.LinkPreview{
		Details: &peekalink.LinkDetails{
			Type:               peekalink.LDTTwitter,
			TwitterLinkDetails: &peekalink.TwitterLinkDetails{},
		},
	}

	emptyData := peekalink.LinkPreview{}

	t.Run("IsYouTubeLink", func(t *testing.T) {
		t.Run("should return true for YouTube links", func(t *testing.T) {
			if youtubeLinkData.IsYouTubeLink() != true {
				t.Errorf("LinkPreview#IsYouTubeLink want %v, got %v", true, false)
			}
		})

		t.Run("should return false for non-YouTube links", func(t *testing.T) {
			if twitterLinkData.IsYouTubeLink() != false {
				t.Errorf("LinkPreview#IsYouTubeLink want %v, got %v", false, true)
			}

			if emptyData.IsYouTubeLink() != false {
				t.Errorf("LinkPreview#IsYouTubeLink want %v, got %v", false, true)
			}
		})
	})

	t.Run("IsTwitterLink", func(t *testing.T) {
		t.Run("should return true for Twitter links", func(t *testing.T) {
			if twitterLinkData.IsTwitterLink() != true {
				t.Errorf("LinkPreview#IsTwitterLink want %v, got %v", true, false)
			}
		})

		t.Run("should return false for non-Twitter links", func(t *testing.T) {
			if youtubeLinkData.IsTwitterLink() != false {
				t.Errorf("LinkPreview#IsTwitterLink want %v, got %v", false, true)
			}

			if emptyData.IsTwitterLink() != false {
				t.Errorf("LinkPreview#IsTwitterLink want %v, got %v", false, true)
			}
		})
	})

	t.Run("GetYouTubeDetails", func(t *testing.T) {
		t.Run("should return non-nil for YouTube links", func(t *testing.T) {
			result, err := youtubeLinkData.GetYouTubeDetails()

			if err != nil {
				t.Error(err)
			}

			if result == nil {
				t.Error("LinkPreview#GetYouTubeDetails want non-nil, got nil")
			}
		})

		t.Run("should return nil for non-YouTube links", func(t *testing.T) {
			result, err := emptyData.GetYouTubeDetails()

			if err == nil {
				t.Error("LinkPreview#GetYouTubeDetails wanted error, got nil")
			}

			if result != nil {
				t.Errorf("LinkPreview#GetYouTubeDetails want %v, got %v", nil, result)
			}
		})
	})

	t.Run("GetTwitterDetails", func(t *testing.T) {
		t.Run("should return non-nil for Twitter links", func(t *testing.T) {
			result, err := twitterLinkData.GetTwitterDetails()

			if err != nil {
				t.Error(err)
			}

			if result == nil {
				t.Error("LinkPreview#GetTwitterDetails want non-nil, got nil")
			}
		})

		t.Run("should return nil for non-Twitter links", func(t *testing.T) {
			result, err := emptyData.GetTwitterDetails()

			if err == nil {
				t.Error("LinkPreview#GetTwitterDetails wanted error, got nil")
			}

			if result != nil {
				t.Errorf("LinkPreview#GetTwitterDetails want %v, got %v", nil, result)
			}
		})
	})
}
