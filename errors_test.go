package peekalink_test

import (
	"fmt"
	"testing"

	"github.com/jozsefsallai/go-peekalink"
)

func TestLinkErrorString(t *testing.T) {
	t.Run("should return correct string value", func(t *testing.T) {
		tests := map[peekalink.LinkError]string{
			peekalink.LinkUnexpectedError:   "LinkUnexpectedError",
			peekalink.LinkMaxRedirectsError: "LinkMaxRedirectsError",
			peekalink.LinkIsPrivateError:    "LinkIsPrivateError",
			peekalink.LinkDoesNotExist:      "LinkDoesNotExist",
			peekalink.LinkTimedOutHTTPError: "LinkTimedOutHTTPError",
			peekalink.LinkEmpty:             "LinkEmpty | LinkPreviewHTTPError | LinkUnreachable",
			peekalink.LinkContentIsTooLarge: "LinkContentIsTooLarge",
		}

		for linkError, expected := range tests {
			actual := linkError.String()

			if actual != expected {
				t.Errorf("LinkError#String want %v, got %v", expected, actual)
			}
		}
	})

	t.Run("should return (error code) for unknown error", func(t *testing.T) {
		linkError := peekalink.LinkError(500)

		expected := "(500)"
		actual := linkError.String()

		if actual != expected {
			t.Errorf("LinkError#String want %v, got %v", expected, actual)
		}
	})
}

func TestIsErrorKind(t *testing.T) {
	message := fmt.Errorf("request failed with error: %s", peekalink.LinkPreviewHTTPError)

	t.Run("should return true if the error contains the provided kind", func(t *testing.T) {
		expected := true
		actual := peekalink.IsErrorKind(message, peekalink.LinkPreviewHTTPError)

		if actual != expected {
			t.Errorf("IsErrorKind(err, peekalink.LinkPreviewHTTPError) want %v, got %v", expected, actual)
		}
	})

	t.Run("should return false if the error does not contain the provided kind", func(t *testing.T) {
		expected := false
		actual := peekalink.IsErrorKind(message, peekalink.LinkMaxRedirectsError)

		if actual != expected {
			t.Errorf("IsErrorKind(err, peekalink.LinkMaxRedirectsError) want %v, got %v", expected, actual)
		}
	})
}
