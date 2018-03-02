package htmldoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	c := DefaultHTTPClient(UserAgentChrome)
	assert.NotNil(t, c)

	f := NewFetcher(c)
	assert.NotNil(t, f)

	d, err := f.FetchDocument("http://example.com/")
	assert.Nil(t, err)
	assert.Equal(t, "http://example.com/", d.Url.String())
	assert.Equal(t, "Example Domain", d.Find("title").Text())
}
