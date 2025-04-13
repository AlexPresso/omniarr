package media

import "omniarr/internal/core/torrent"

type Media struct {
	ID            string  `json:"id"`
	Title         string  `json:"title"`
	OriginalTitle string  `json:"originalTitle"`
	Description   string  `json:"description"`
	Cover         string  `json:"cover"`
	Popularity    float32 `json:"popularity"`
	ReleaseDate   string  `json:"releaseDate"`
	Type          string  `json:"type"`
}

type Details struct {
	Torrents []torrent.Torrent `json:"torrents"`
	*Media
}
