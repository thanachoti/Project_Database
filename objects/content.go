package object

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
	Languages    []string // List of available languages
	Subtitles    []string // List of available subtitles
	Categories   []string // List of categories
}
