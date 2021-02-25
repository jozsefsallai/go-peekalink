package peekalink

import (
	"fmt"
	"strings"
)

// LinkError represents an error that occurred when trying to fetch a link
// preview.
type LinkError int

const (
	// LinkUnexpectedError happens on unexpected errors (any non-200 HTTP status
	// code that isn't in the list of default errors).
	LinkUnexpectedError LinkError = -1

	// LinkMaxRedirectsError happens on HTTP 400 if a link has too many redirects.
	LinkMaxRedirectsError LinkError = 400

	// LinkIsPrivateError happens on HTTP 403 if the provided link is private.
	LinkIsPrivateError LinkError = 403

	// LinkDoesNotExist happens on HTTP 404 if a link cannot be found.
	LinkDoesNotExist LinkError = 404

	// LinkTimedOutHTTPError happens on HTTP 408 if the link preview request took
	// too long.
	LinkTimedOutHTTPError LinkError = 408

	// LinkEmpty happens on HTTP 409 if there is nothing that can be displayed
	// about the link.
	LinkEmpty LinkError = 409

	// LinkPreviewHTTPError happens on HTTP 409 if an HTTP error happened in the
	// link preview fetching request.
	LinkPreviewHTTPError LinkError = 409

	// LinkUnreachable happens on HTTP 409 if the link could not be reached at
	// that point.
	LinkUnreachable LinkError = 409

	// LinkContentIsTooLarge happens on HTTP 413 if the content returned by the
	// link is too large and cannot be returned.
	LinkContentIsTooLarge LinkError = 413
)

func (linkError LinkError) String() string {
	switch linkError {
	case LinkUnexpectedError:
		return "LinkUnexpectedError"
	case LinkMaxRedirectsError:
		return "LinkMaxRedirectsError"
	case LinkIsPrivateError:
		return "LinkIsPrivateError"
	case LinkDoesNotExist:
		return "LinkDoesNotExist"
	case LinkTimedOutHTTPError:
		return "LinkTimedOutHTTPError"
	case LinkEmpty:
		return "LinkEmpty | LinkPreviewHTTPError | LinkUnreachable"
	case LinkContentIsTooLarge:
		return "LinkContentIsTooLarge"
	default:
		return fmt.Sprintf("(%d)", linkError)
	}
}

// IsErrorKind is a wrapper function that checks if an error contains a given
// error kind.
func IsErrorKind(err error, errorKind LinkError) bool {
	return strings.Contains(err.Error(), errorKind.String())
}
