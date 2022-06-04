package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fa-at-pulsit/tvseries-cli/pkg/themoviedb"
	"github.com/fa-at-pulsit/tvseries-cli/pkg/ui"
	log "github.com/sirupsen/logrus"
)

func main() {
	if traceLevel, err := log.ParseLevel(getEnv("LOG_TRACE_LEVEL", "INFO")); err == nil {
		log.SetLevel(traceLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	clientTheMovieDB := *themoviedb.NewClient()
	errs := make(chan error, 1)

	go func() {
		adultContent := getEnv("ADULT_CONTENT", "false") == "true"
		fmt.Print("\n")
		fmt.Println(strings.Repeat("-", 120))
		for {
			if ok, searchQuery := ui.PromptTVTitle(); ok {
				res, err := clientTheMovieDB.SarchTV(context.Background(), searchQuery, 1, adultContent)
				fmt.Print(".")
				if err != nil {
					log.Error(err.Error())
				}
				if res != nil {
					items := []string{}
					if res.TotalPages > 1 {
						for page := 2; page < res.TotalPages; page++ {
							res1, err := clientTheMovieDB.SarchTV(context.Background(), searchQuery, page, adultContent)
							fmt.Print(".")
							if err != nil {
								log.Error(err.Error())
							}
							if res1 != nil {
								res.Results = append(res.Results, res1.Results...)
							}
						}
					}

					for i := range res.Results {
						items = append(items, res.Results[i].Name)
					}

					if ok, idx, _ := ui.SelectTVTitle(items); ok {
						tvSeason := res.Results[idx]
						serieSessions, err := clientTheMovieDB.GetTVSerieSeasons(context.Background(), tvSeason.ID)
						if err != nil {
							log.Error(err.Error())
						}
						if res != nil {
							items := []string{}
							for i := range serieSessions.Seasons {
								items = append(items, serieSessions.Seasons[i].Name)
							}
							if ok, idx, _ := ui.SelectTVSerieSeason(items); ok {
								serieEpisodes, err := clientTheMovieDB.GetTVSerieEpisodes(context.Background(), tvSeason.ID, serieSessions.Seasons[idx].SeasonNumber)
								if err != nil {
									log.Error(err.Error())
								}
								if res != nil {
									items := []string{}
									for i := range serieEpisodes.Episodes {
										items = append(items, fmt.Sprintf("S%02dE%02d - %s", serieEpisodes.Episodes[i].SeasonNumber, serieEpisodes.Episodes[i].EpisodeNumber, serieEpisodes.Episodes[i].Name))
									}
									if ok, idx, _ := ui.SelectTVSerieEpisode(items); ok {
										tvEpisode := serieEpisodes.Episodes[idx]
										fmt.Println(strings.Repeat("-", 120))
										fmt.Printf("%s :: Season %02d :: Episode %02d - %s\n", tvSeason.Name, tvEpisode.SeasonNumber, tvEpisode.EpisodeNumber, tvEpisode.Name)
										fmt.Println(strings.Repeat("-", 3))
										fmt.Println(tvEpisode.Overview)
										fmt.Println(strings.Repeat("-", 3))
										fmt.Print("\n\n\n")
									}
								}
							}
						}
					}
				}
			} else {
				break
			}
		}
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	log.Printf("Terminated %s\n", <-errs)
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
