package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"

	"./search"
)

type (
	// item defines the fields associated whit item tag
	// in the rss document.
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubData"`
		Title       string   `xml:"title"`
		Description string   `json:"description"`
		Link        string   `json:"link"`
		GUID        string   `json:"guid"`
		GeoRssPoint string   `json:"georss:point"`
	}

	// image defines the fields associated with the image tag
	// in the rss document.
	image struct {
		XMLName xml.Name `json:"image"`
		URL     string   `json:"url"`
		Title   string   `json:"title"`
		Link    string   `json:"link"`
	}

	// channel defines the fields associated with the channel tag
	// in the rss document.
	channel struct {
		XMLName        xml.Name `json:"channel"`
		Title          string   `json:"title"`
		Description    string   `json:"description"`
		Link           string   `json:"link"`
		PubDate        string   `json:"pubData"`
		LastBuildDate  string   `json:"lastBuildDate"`
		TTL            string   `json:"ttl"`
		Language       string   `json:"language"`
		ManagingEditor string   `json:"managingEditor"`
		WebMaster      string   `json:"webMaster"`
		Image          image    `json:"image"`
		Item           []item   `json:"item"`
	}

	// rssDocument defines the fields associated with the rss document
	rssDocument struct {
		XMLName xml.Name `json:"rss"`
		Channel channel  `json:"channel"`
	}
)

// rssMatcher implements the Matcher interface.
type rssMatcher struct{}

// init registers the matcher with the program.
func init() {
	var m rssMatcher
	search.Register("rss", matcher)
}

// Search looks at the document for the specified search term.
// TODO func (m rssMatcher) Search(f *
// retrieve performs a HTTP Get request for the rss feed and decoddes
func (m rssMatcher) retrieve(f *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed URI provided")
	}

	// Retrieve the rss feed document from the web.
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// Close the response once we return from the function.
	defer resp.Body.Close()

	// Check the status code for 200 so we know we have received a
	// proper response.
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d", resp.StatusCode)
	}

	// Decode the rss feed document into our struct type.
	// We don't need to check for errors, the caller can do this.
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}
