package media

import (
	"github.com/Open-pi/gol"
	tmdb "github.com/cyruzin/golang-tmdb"
	"strconv"
	"strings"
)

type TMDBMediaLight struct {
	ID            int64
	Title         string
	Name          string
	OriginalTitle string
	OriginalName  string
	Overview      string
	Popularity    float32
	ReleaseDate   string
	FirstAirDate  string
	PosterPath    string
	Type          string
}

type OpenLibraryMediaLight struct {
	Key   string
	Title string
	Cover string `mapstructure:"cover_i"`
}

func (m *TMDBMediaLight) GetID() string { return strconv.FormatInt(m.ID, 10) }
func (m *TMDBMediaLight) GetTitle() string {
	if m.Title != "" {
		return m.Title
	}
	return m.Name
}
func (m *TMDBMediaLight) GetOriginalTitle() string {
	if m.OriginalTitle != "" {
		return m.OriginalTitle
	}
	return m.OriginalName
}
func (m *TMDBMediaLight) GetOverview() string    { return m.Overview }
func (m *TMDBMediaLight) GetPopularity() float32 { return m.Popularity }
func (m *TMDBMediaLight) GetReleaseDate() string {
	if m.ReleaseDate != "" {
		return m.ReleaseDate
	}
	return m.FirstAirDate
}
func (m *TMDBMediaLight) GetCover() string { return m.PosterPath }
func (m *TMDBMediaLight) GetType() Type    { return Type(m.Type) }

type MovieWrapper struct {
	*tmdb.MovieDetails
}

func (m *MovieWrapper) GetID() string            { return strconv.FormatInt(m.ID, 10) }
func (m *MovieWrapper) GetTitle() string         { return m.Title }
func (m *MovieWrapper) GetOriginalTitle() string { return m.OriginalTitle }
func (m *MovieWrapper) GetOverview() string      { return m.Overview }
func (m *MovieWrapper) GetPopularity() float32   { return m.Popularity }
func (m *MovieWrapper) GetReleaseDate() string   { return m.ReleaseDate }
func (m *MovieWrapper) GetCover() string         { return m.PosterPath }
func (m *MovieWrapper) GetType() Type            { return "movie" }

type TVWrapper struct {
	*tmdb.TVDetails
}

func (t *TVWrapper) GetID() string            { return strconv.FormatInt(t.ID, 10) }
func (t *TVWrapper) GetTitle() string         { return t.Name }
func (t *TVWrapper) GetOriginalTitle() string { return t.OriginalName }
func (t *TVWrapper) GetOverview() string      { return t.Overview }
func (t *TVWrapper) GetPopularity() float32   { return t.Popularity }
func (t *TVWrapper) GetReleaseDate() string   { return t.FirstAirDate }
func (t *TVWrapper) GetCover() string         { return t.PosterPath }
func (t *TVWrapper) GetType() Type            { return "tv" }

type BookWrapper struct {
	*gol.Work
}

func (b *OpenLibraryMediaLight) GetID() string            { return strings.ReplaceAll(b.Key, "/works/", "") }
func (b *OpenLibraryMediaLight) GetTitle() string         { return b.Title }
func (b *OpenLibraryMediaLight) GetOriginalTitle() string { return "" }
func (b *OpenLibraryMediaLight) GetOverview() string      { return "" }
func (b *OpenLibraryMediaLight) GetPopularity() float32   { return 0 }
func (b *OpenLibraryMediaLight) GetReleaseDate() string   { return "" }
func (b *OpenLibraryMediaLight) GetCover() string         { return b.Cover }
func (b *OpenLibraryMediaLight) GetType() Type            { return "book" }

func (b *BookWrapper) GetID() string {
	if key, err := b.Key(); err == nil {
		return key
	}
	return ""
}
func (b *BookWrapper) GetTitle() string {
	if title, err := b.Title(); err == nil {
		return title
	}
	return ""
}
func (b *BookWrapper) GetOriginalTitle() string { return "" }
func (b *BookWrapper) GetOverview() string {
	if desc, err := b.Desc(); err == nil {
		return desc
	}
	return ""
}
func (b *BookWrapper) GetPopularity() float32 { return 0 }
func (b *BookWrapper) GetReleaseDate() string { return "" }
func (b *BookWrapper) GetCover() string       { return b.FirstCoverKey() }
func (b *BookWrapper) GetType() Type          { return "book" }
