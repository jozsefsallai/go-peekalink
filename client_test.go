package peekalink_test

import (
	"os"
	"testing"

	"github.com/ankida/go-peekalink"
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
			preview, err := client.Preview("https://peekalink.io")
			if err != nil {
				t.Fatal(err)
			}

			expectedDomain := "peekalink.io"

			if preview.Domain != expectedDomain {
				t.Errorf("Client#Preview().Domain want %v, got %v", expectedDomain, preview.Domain)
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
