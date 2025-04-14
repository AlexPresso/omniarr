package media

type Type string

type MediaInterface interface {
	GetID() string
	GetTitle() string
	GetOriginalTitle() string
	GetOverview() string
	GetPopularity() float32
	GetReleaseDate() string
	GetCover() string
	GetType() Type
}

type Media struct {
	ID            string  `json:"id"`
	Title         string  `json:"title"`
	OriginalTitle string  `json:"originalTitle"`
	Description   string  `json:"description"`
	Cover         string  `json:"cover"`
	Popularity    float32 `json:"popularity"`
	ReleaseDate   string  `json:"releaseDate"`
	Type          Type    `json:"type"`
}

func ToMedia(m MediaInterface) *Media {
	return &Media{
		ID:            m.GetID(),
		Title:         m.GetTitle(),
		OriginalTitle: m.GetOriginalTitle(),
		Description:   m.GetOverview(),
		Cover:         m.GetCover(),
		Popularity:    m.GetPopularity(),
		ReleaseDate:   m.GetReleaseDate(),
		Type:          m.GetType(),
	}
}
