package themoviedb

type SearchTVResults struct {
	PosterPath       string    `json:"poster_path,omitempty"`
	Popularity       float64   `json:"popularity,omitempty"`
	ID               int       `json:"id,omitempty"`
	BackdropPath     string    `json:"backdrop_path,omitempty"`
	VoteAverage      float64   `json:"vote_average,omitempty"`
	Overview         string    `json:"overview,omitempty"`
	FirstAirDate     string    `json:"first_air_date,omitempty"`
	OriginCountry    []*string `json:"origin_country,omitempty"`
	GenreIDs         []*int    `json:"genre_ids,omitempty"`
	OriginalLanguage string    `json:"original_language,omitempty"`
	VoteCount        int       `json:"vote_count,omitempty"`
	Name             string    `json:"name,omitempty"`
	OriginalName     string    `json:"original_name,omitempty"`
}

type SearchTVResponse struct {
	Page         int               `json:"page"`
	Results      []SearchTVResults `json:"results"`
	TotalResults int               `json:"total_results"`
	TotalPages   int               `json:"total_pages"`
}
