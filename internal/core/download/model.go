package download

import "omniarr/internal/core/media"

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
