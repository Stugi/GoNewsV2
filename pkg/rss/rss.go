package rss

import (
	"encoding/xml"
	"fmt"
	"gonews/v2/pkg/storage"
	"io"
	"net/http"
	"strings"
	"time"
)

type RSSDocument struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Link        string `xml:"link"`
		Items       []Item `xml:"item"`
	} `xml:"channel"`
}

type Item struct {
	GUID        string `xml:"guid"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
}

func Parse(url string) ([]storage.Post, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var doc RSSDocument
	err = xml.Unmarshal(body, &doc)
	if err != nil {
		return nil, err
	}

	source := &storage.Source{
		Link:        doc.Channel.Link,
		Name:        doc.Channel.Title,
		Description: doc.Channel.Description,
	}

	var posts []storage.Post
	for _, item := range doc.Channel.Items {
		posts = append(posts, NewPostFromItem(item, source))
	}

	return posts, err
}

func NewPostFromItem(item Item, source *storage.Source) storage.Post {
	return storage.Post{
		Title:      item.Title,
		Content:    item.Description,
		PubTime:    parseDate(item.PubDate),
		Link:       item.Link,
		Source:     source,
		ExternalID: item.GUID,
	}
}

// const maskRegexp = `([a-zA-Z]{3}), ([0-9]{1,2}) ([A-Z]{3}) ([0-9]{4}) ([0-9]{2}):([0-9]{2}):([0-9]{2}) ([+-][0-9]{4})`

func parseDate(date string) int {
	modifiedDate := strings.Replace(date, "GMT", "+0000", 1)

	layouts := []string{
		"Mon, 2 Jan 2006 15:04:05 -0700",
	}
	for _, layout := range layouts {
		t, err := time.Parse(layout, modifiedDate)
		if err == nil {
			return int(t.Unix())
		}
	}

	fmt.Printf("wrong date format: %s\n", modifiedDate)

	return 0
}
