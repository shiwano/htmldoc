package htmldoc

import (
	"net/http"
)

// UserAgent constants.
const (
	UserAgentIE      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Trident/7.0; rv:11.0) like Gecko"
	UserAgentEdge    = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.79 Safari/537.36 Edge/14.14393"
	UserAgentFirefox = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0"
	UserAgentChrome  = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36"
)

// DefaultHTTPClient creates a simple HTTP client.
func DefaultHTTPClient(userAgent string) *http.Client {
	return &http.Client{
		Transport: &roundTripper{
			underlying: http.DefaultTransport,
			userAgent:  userAgent,
		},
	}
}

type roundTripper struct {
	underlying http.RoundTripper
	userAgent  string
}

func (t *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.userAgent)
	return t.underlying.RoundTrip(req)
}
