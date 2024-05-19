package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

type Match struct {
	URL       string `json:"URL"`
	Team1     string `json:"Team1"`
	Team2     string `json:"Team2"`
	Score     []int  `json:"Score"`
	Rounds    []int  `json:"Rounds"`
	StartTime string `json:"StartTime"`
	Tag       string `json:"Tag"`
	Status    string `json:"Status"`
	Region    string `json:"Region"`
}

type resJSON struct {
	Matches []Match `json:"matches"`
	Count   int     `json:"count"`
}

func ParseResponse(body io.Reader, live bool, count int, region string) {

	var res resJSON
	err := json.NewDecoder(body).Decode(&res)
	if err != nil {
		fmt.Println("Error parsing response.", err)
	}

	res.Matches = filterMatches(res.Matches, live, count, region)

	// Clear the previous output
	fmt.Print("\033[1A\033[K")
	for _, match := range res.Matches {
		printMatchData(match)
	}
}

func filterMatches(matches []Match, live bool, count int, region string) []Match {
	var filteredMatches []Match
	for _, match := range matches {
		if live && match.Status != "live" {
			continue
		}
		if region != "" && strings.ToLower(match.Region) != region {
			continue
		}
		filteredMatches = append(filteredMatches, match)
		if len(filteredMatches) == count {
			break
		}
	}
	return filteredMatches
}

func printMatchData(match Match) {
	fmt.Println("+----------------------------------------------------------+")
	printVariableLengthString(match.Tag, 60, color.FgHiBlue)
	fmt.Println("\n|----------------------------------------------------------|")
	printVariableLengthString(match.Team1, 30, color.FgWhite)
	if match.Status == "live" {
		printVariableLengthString(fmt.Sprintf("%d  (%d)", match.Score[0], match.Rounds[0]), 15, color.FgHiGreen)
		printVariableLengthString(match.StartTime, 15, color.FgHiRed)
	} else if match.Status == "upcoming" {
		printVariableLengthString("-", 15, color.FgHiGreen)
		printVariableLengthString(match.StartTime, 15, color.FgHiCyan)
	} else {
		printVariableLengthString(fmt.Sprintf("%d", match.Score[0]), 15, color.FgHiGreen)
		printVariableLengthString(match.StartTime, 15, color.FgHiRed)
	}

	fmt.Println()
	printVariableLengthString(match.Team2, 30, color.FgWhite)
	if match.Status == "live" {
		printVariableLengthString(fmt.Sprintf("%d  (%d)", match.Score[1], match.Rounds[1]), 15, color.FgHiYellow)
	} else if match.Status == "upcoming" {
		printVariableLengthString("-", 15, color.FgHiYellow)
	} else {
		printVariableLengthString(fmt.Sprintf("%d", match.Score[1]), 15, color.FgHiYellow)
	}
	printVariableLengthString("", 15, color.FgRed)
	fmt.Println("\n+----------------------------------------------------------+")
	fmt.Println()
}

func printVariableLengthString(s any, size int, colorStr color.Attribute) {
	fmt.Print("|")
	color.Set(colorStr)
	str := fmt.Sprintf("%v", s)
	spacesRequired := size - 2 - utf8.RuneCountInString(str)
	spaces := strings.Repeat(" ", spacesRequired/2)
	finalString := fmt.Sprintf("%s\033[1m%s\033[0m%s", spaces, str, spaces)
	if spacesRequired%2 != 0 {
		finalString += " "
	}
	fmt.Print(finalString)
	color.Unset()
	fmt.Print("|")
}
