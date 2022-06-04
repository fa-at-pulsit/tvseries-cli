package themoviedb

type TVSeasonPrimary struct {
	ID           int    `json:"id"`
	Name         string `json:"name,omitempty"`
	Overview     string `json:"overview,omitempty"`
	PosterPath   string `json:"poster_path,omitempty"`
	EpisodeCount int    `json:"episode_count"`
	SeasonNumber int    `json:"season_number"`
}

type TVSeasonEpisodePrimary struct {
	ID            int    `json:"id"`
	Name          string `json:"name,omitempty"`
	Overview      string `json:"overview,omitempty"`
	SeasonNumber  int    `json:"season_number"`
	EpisodeNumber int    `json:"episode_number"`
}

type TVSerieMin struct {
	ID      int               `json:"id"`
	Name    string            `json:"name"`
	Seasons []TVSeasonPrimary `json:"seasons"`
}

type TVSerieSeasonMin struct {
	ID           string                   `json:"_id"`
	Name         string                   `json:"name"`
	Overview     string                   `json:"overview,omitempty"`
	Episodes     []TVSeasonEpisodePrimary `json:"episodes"`
	SeasonNumber int                      `json:"season_number"`
}
