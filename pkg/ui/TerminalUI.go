package ui

import (
	"errors"

	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
)

func PromptTVTitle() (bool, string) {
	validate := func(input string) error {
		if len(input) < 2 {
			return errors.New("title must have more than 2 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter a TV series title",
		Validate: validate,
	}

	searchQuery, err := prompt.Run()

	if err != nil {
		log.Errorf("Prompt failed %v\n", err)
		return false, ""
	}

	log.Tracef("You choose %q\n", searchQuery)
	return true, searchQuery
}

func SelectTVTitle(items []string) (bool, int, string) {
	prompt := promptui.Select{
		Label: "Select a TV series",
		Items: items,
	}
	index, result, err := prompt.Run()
	if err != nil {
		log.Errorf("Prompt failed %v\n", err)
		return false, 0, ""
	}

	log.Tracef("You choose %q [%d]\n", result, index)

	return true, index, result
}

func SelectTVSerieSeason(items []string) (bool, int, string) {
	prompt := promptui.Select{
		Label: "Select a TV serie season",
		Items: items,
	}
	index, result, err := prompt.Run()
	if err != nil {
		log.Errorf("Prompt failed %v\n", err)
		return false, 0, ""
	}

	log.Tracef("You choose %q [%d]\n", result, index)

	return true, index, result
}

func SelectTVSerieEpisode(items []string) (bool, int, string) {
	prompt := promptui.Select{
		Label: "Select a TV serie episode",
		Items: items,
	}
	index, result, err := prompt.Run()
	if err != nil {
		log.Errorf("Prompt failed %v\n", err)
		return false, 0, ""
	}

	log.Tracef("You choose %q [%d]\n", result, index)

	return true, index, result
}
