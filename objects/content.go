package object

import "github.com/lib/pq"

type Content struct {
	ContentID    int
	Title        string
	Description  string
	ReleaseYear  int
	Duration     int
	ContentType  string
	TotalSeasons int
	ThumbnailURL string
	VideoURL     string
	Rating       float64
	Director     string
	Languages    pq.StringArray
	Subtitles    pq.StringArray
	Categories   pq.StringArray
}
