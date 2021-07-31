package peekalink_test

import (
	"os"
	"testing"

	"github.com/jozsefsallai/go-peekalink"
)

func TestNewClient(t *testing.T) {
	client := peekalink.NewClient("hello-world")

	t.Run("should have the correct key", func(t *testing.T) {
		expected := "hello-world"

		if client.Key != expected {
			t.Errorf("NewClient(\"hello-world\").Key got %v, want %v", client.Key, expected)
		}
	})

	t.Run("should have the correct base URL", func(t *testing.T) {
		expected := "https://api.peekalink.io"

		if client.BaseURL != expected {
			t.Errorf("NewClient(\"hello-world\").BaseURL got %v, want %v", client.BaseURL, expected)
		}
	})
}

func TestClient(t *testing.T) {
	client := peekalink.NewClient("hello-world")

	t.Run("SetKey", func(t *testing.T) {
		newKey := "foo-bar"
		client.SetKey(newKey)

		if client.Key != newKey {
			t.Errorf("Client#SetKey, client key is %v, want %v", client.Key, newKey)
		}
	})

	t.Run("SetBaseURL", func(t *testing.T) {
		newURL := "http://localhost:3000"
		client.SetBaseURL(newURL)

		if client.BaseURL != newURL {
			t.Errorf("Client#SetBaseURL, client base URL is %v, want %v", client.BaseURL, newURL)
		}
	})

	realKey := os.Getenv("PEEKALINK_API_KEY")
	if len(realKey) == 0 {
		return
	}

	client.SetBaseURL("https://api.peekalink.io")
	client.SetKey(realKey)

	t.Run("Client#Preview", func(t *testing.T) {
		t.Run("should return link preview data", func(t *testing.T) {
			preview, err := client.Preview("https://bit.ly/3frD2OP")
			if err != nil {
				t.Fatal(err)
			}

			expected := peekalink.LinkPreview{
				URL:              "https://bit.ly/3frD2OP",
				Domain:           "bit.ly",
				ContentType:      peekalink.ContentTypeHTML,
				MimeType:         "text/html",
				Redirected:       true,
				RedirectionURL:   "https://www.youtube.com/watch?feature=youtu.be&v=dQw4w9WgXcQ",
				RedirectionCount: 2,
				TrackersDetected: true,
				Details: &peekalink.LinkDetails{
					Type: "youtube",
					YouTubeLinkDetails: &peekalink.YouTubeLinkDetails{
						VideoID:  "dQw4w9WgXcQ",
						Duration: "213.0",
					},
				},
			}

			if preview.URL != expected.URL {
				t.Errorf("Client#Preview().URL want %v, got %v", expected.URL, preview.URL)
			}

			if preview.Domain != expected.Domain {
				t.Errorf("Client#Preview().Domain want %v, got %v", expected.Domain, preview.Domain)
			}

			if len(preview.LastUpdated) == 0 {
				t.Error("Client#Preview().LastUpdated want non-zero length, got 0")
			}

			if len(preview.NextUpdate) == 0 {
				t.Error("Client#Preview().NextUpdate want non-zero length, got 0")
			}

			if preview.ContentType != expected.ContentType {
				t.Errorf("Client#Preview().ContentType want %v, got %v", expected.ContentType, preview.ContentType)
			}

			if preview.MimeType != expected.MimeType {
				t.Errorf("Client#Preview().MimeType want %v, got %v", expected.MimeType, preview.MimeType)
			}

			if preview.Size == 0 {
				t.Error("Client#Preview().Size want non-zero value, got 0")
			}

			if preview.Redirected != expected.Redirected {
				t.Errorf("Client#Preview().Redirected want %v, got %v", expected.Redirected, preview.Redirected)
			}

			if preview.RedirectionURL != expected.RedirectionURL {
				t.Errorf("Client#Preview().RedirectionURL want %v, got %v", expected.RedirectionURL, preview.RedirectionURL)
			}

			if preview.RedirectionCount != expected.RedirectionCount {
				t.Errorf("Client#Preview().RedirectionCount want %v, got %v", expected.RedirectionCount, preview.RedirectionCount)
			}

			if len(preview.RedirectionTrail) != int(expected.RedirectionCount) {
				t.Errorf("Client#Preview().RedirectionTrail want length of %v, got length of %v", expected.RedirectionCount, len(preview.RedirectionTrail))
			}

			if len(preview.Title) == 0 {
				t.Error("Client#Preview().Title want non-zero length, got 0")
			}

			if len(preview.Description) == 0 {
				t.Error("Client#Preview().Description want non-zero length, got 0")
			}

			if len(preview.Name) == 0 {
				t.Error("Client#Preview().Name want non-zero length, got 0")
			}

			if preview.TrackersDetected != expected.TrackersDetected {
				t.Errorf("Client#Preview().TrackersDetected want %v, got %v", expected.TrackersDetected, preview.TrackersDetected)
			}

			if len(preview.Icon.URL) == 0 {
				t.Error("Client#Preview().Icon.URL want non-zero length, got 0")
			}

			if preview.Icon.Width == 0 {
				t.Error("Client#Preview().Icon.Width want non-zero value, got 0")
			}

			if preview.Icon.Height == 0 {
				t.Error("Client#Preview().Icon.Height want non-zero value, got 0")
			}

			if len(preview.Image.URL) == 0 {
				t.Error("Client#Preview().Image.URL want non-zero length, got 0")
			}

			if preview.Image.Width == 0 {
				t.Error("Client#Preview().Image.Width want non-zero value, got 0")
			}

			if preview.Image.Height == 0 {
				t.Error("Client#Preview().Image.Height want non-zero value, got 0")
			}

			if preview.Details == nil {
				t.Error("Client#Preview().Details want non-nil, got nil")
			}

			if preview.Details.Type != expected.Details.Type {
				t.Errorf("Client#Preview().Details.Type want %v, got %v", expected.Details.Type, preview.Details.Type)
			}

			if preview.Details.VideoID != expected.Details.VideoID {
				t.Errorf("Client#Preview().Details.VideoID want %v, got %v", expected.Details.VideoID, preview.Details.VideoID)
			}

			if preview.Details.Duration != expected.Details.Duration {
				t.Errorf("Client#Preview().Details.Duration want %v, got %v", expected.Details.Duration, preview.Details.Duration)
			}

			if preview.Details.YTViewCount == 0 {
				t.Error("Client#Preview().Details.YTViewCount want non-zero value, got 0")
			}

			if preview.Details.YTLikeCount == 0 {
				t.Error("Client#Preview().Details.YTLikeCount want non-zero value, got 0")
			}

			if preview.Details.YTDislikeCount == 0 {
				t.Error("Client#Preview().Details.YTDislikeCount want non-zero value, got 0")
			}

			if preview.Details.YTCommentCount == 0 {
				t.Error("Client#Preview().Details.YTCommentCount want non-zero value, got 0")
			}

			if len(preview.Details.PublishedAt) == 0 {
				t.Error("Client#Preview().Details.PublishedAt want non-zero length, got 0")
			}
		})

		t.Run("should return error on invalid link", func(t *testing.T) {
			_, err := client.Preview("https://this.domain.doesnot.exist")
			if err == nil {
				t.Fatalf("Client#Preview() expected to error, did not")
			}
		})
	})

	t.Run("Client#IsAvailable", func(t *testing.T) {
		t.Run("should return link availability", func(t *testing.T) {
			data, err := client.IsAvailable("https://peekalink.io")
			if err != nil {
				t.Fatal(err)
			}

			expectedResult := true

			if data.IsAvailable != expectedResult {
				t.Errorf("Client#IsAvailable() want %v, got %v", expectedResult, data.IsAvailable)
			}
		})
	})
}
