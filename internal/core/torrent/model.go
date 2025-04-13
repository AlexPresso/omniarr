package torrent

type Torrent struct {
	Guid        string `json:"guid"`
	Title       string `json:"title"`
	Size        uint   `json:"size"`
	Seeders     uint   `json:"seeders"`
	Leechers    uint   `json:"leechers"`
	Category    string `json:"category"`
	Indexer     string `json:"indexer"`
	PublishDate string `json:"publishDate"`
}
