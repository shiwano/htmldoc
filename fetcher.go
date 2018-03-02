package htmldoc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
)

// Fetcher fetches a goquery Document.
type Fetcher interface {
	FetchDocument(url string) (*goquery.Document, error)
}

// NewFetcher creates a fetcher.
func NewFetcher(client *http.Client) Fetcher {
	return &fetcher{
		client:   client,
		detector: chardet.NewTextDetector(),
	}
}

type fetcher struct {
	client   *http.Client
	detector *chardet.Detector
}

func (f *fetcher) FetchDocument(url string) (*goquery.Document, error) {
	data, res, err := f.fetchData(url)
	if err != nil {
		return nil, err
	}
	decodedData, err := f.decode(data)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(decodedData)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	doc.Url = res.Request.URL
	return doc, nil
}

func (f *fetcher) fetchData(url string) ([]byte, *http.Response, error) {
	res, err := f.client.Get(url)
	if err != nil {
		return nil, nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, res, fmt.Errorf("Failed to fetch data: %v", res.StatusCode)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}
	return data, res, nil
}

func (f *fetcher) decode(data []byte) ([]byte, error) {
	enc, err := f.detectEncoding(data)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	decodeReader := transform.NewReader(reader, enc.NewDecoder())
	if _, err := io.Copy(writer, decodeReader); err != nil {
		return nil, err
	}
	if err := writer.Flush(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (f *fetcher) detectEncoding(data []byte) (encoding.Encoding, error) {
	detected, err := f.detector.DetectBest(data)
	if err != nil {
		return nil, err
	}
	enc, _ := charset.Lookup(detected.Charset)
	if enc == nil {
		return nil, fmt.Errorf("Unsupported charset: %v", detected.Charset)
	}
	return enc, nil
}
