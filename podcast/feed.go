package podcast

import (
	"time"

	"github.com/eduncan911/podcast"
)

type Feed struct {
	Title         string
	Link          string
	Description   string
	PubDate       *time.Time
	LastBuildDate *time.Time
}

func NewFeed() *Feed {
	return nil
}

func (f *Feed) buildFeed() {
	podcast.New(f.Title, f.Link, f.Description, f.PubDate, f.LastBuildDate)
}
