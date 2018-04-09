// Utility functions to works with URLs
package urlutil

import (
	"strings"
	"net/http"
	"time"
)

// Check if a url starts with http:// or https://
func IsValidUrl(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

// Try to get the headers from the target url.
//
// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/HEAD
func IsReachable(url string) bool {
	if !(IsValidUrl(url)) {
		return false
	}

	cli := http.Client{Timeout: time.Duration(10 * time.Second)}
	r, err := cli.Head(url)

	return err == nil && r.StatusCode < 400
}
