package download

import "omniarr/internal/core/media"

type SearchQuery struct {
	ExternalId    string `json:"externalId"`
	Title         string `json:"title"`
	OriginalTitle string `json:"originalTitle"`
	Year          string `json:"year"`
	Type          string `json:"type"`
}

type Download struct {
	Title       string `json:"title"`
	Size        uint   `json:"size"`
	Seeders     uint   `json:"seeders"`
	Leechers    uint   `json:"leechers"`
	Indexer     string `json:"indexer"`
	PublishDate string `json:"publishDate"`
	Link        string `json:"link"`
	MagnetUri   string `json:"magnetURI"`
}

type QueueDownloadRequest struct {
	Url  string     `json:"url"`
	Type media.Type `json:"type"`
}
