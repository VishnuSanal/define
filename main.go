package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/fatih/color"
)

type Word struct {
	WordString string `json:"word"`
	Meanings   []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string `json:"definition"`
			Example    string `json:"example"`
		} `json:"definitions"`
	} `json:"meanings"`
}

func main() {

	if len(os.Args) < 2 {
		panic("No word provided!")
	}

	argWord := os.Args[1]

	queryURL := "https://api.dictionaryapi.dev/api/v2/entries/en/" + argWord

	resp, err := http.Get(queryURL)

	if err != nil {
		fmt.Println("HTTP request failed: please make sure that you are conencted to the internet!")
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("API not available: please make sure that you are conencted to the internet!")
	}

	var word []Word

	err = json.NewDecoder(resp.Body).Decode(&word)

	if err != nil {
		panic(err)
	}

	color.HiBlue(argWord + "\n\n")

	for _, meaning := range word[0].Meanings {

		fmt.Printf("PartOfSpeech: %s\n", meaning.PartOfSpeech)

		for _, definition := range meaning.Definitions {

			fmt.Printf(": %s\n", definition.Definition)

		}

		fmt.Println()

	}

}
