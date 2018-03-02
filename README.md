# htmldoc [![Build Status](https://secure.travis-ci.org/shiwano/htmldoc.png?branch=master)](http://travis-ci.org/shiwano/htmldoc)

> :closed_book: Fetch a HTML document with goquery.

## Installation

```bash
$ go get -u github.com/shiwano/htmldoc
```

## Usage

```go
import (
    "fmt"
    "github.com/shiwano/htmldoc"
)

func main() {
    c := htmldoc.DefaultHTTPClient(htmldoc.UserAgentChrome)
    f := htmldoc.NewFetcher(c)
    d := f.FetchDocument("http://example.com/") // HTTP GET only.

    fmt.Println(d.Url.String())                 // You will get: http://example.com/
    fmt.Println(d.Find("title").Text())         // You will get: Example Domain
}
```

This is the only thing it can do. See also [goquery](https://github.com/PuerkitoBio/goquery) documents.
