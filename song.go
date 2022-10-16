package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)


// song struct
type Song struct {
	Id               string  `json:"Id"`
	Name             string  `json:"Name"`
	Artists          string  `json:"Artists"`
	Danceability     float64 `json:"Danceability"`
	Energy           float64 `json:"Energy"`
	Key              int	 `json:"Key"`
	Loudness         float64 `json:"Loudness"`
	Mode             int	 `json:"Mode"`
	Speechiness      float64 `json:"Speechiness"`
	Acousticness     float64 `json:"Acousticness"`
	Instrumentalness float64 `json:"Instrumentalness"`
	Liveness         float64 `json:"Liveness"`
	Valence          float64 `json:"Valence"`
	Tempo            float64 `json:"Tempo"`
	Duration_ms      int	 `json:"Duration_ms"`
	Time_signature   int	 `json:"Time_signature"`
}

func stringToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

func stringToInt(s string) int {
	s = strings.TrimSuffix(s, ".0")
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// parseCsv parses a csv file and returns a 2D slice of strings for the data and a slice of strings for the headers
func parseCsv(filepath string) ([][]string, []string) {

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	var records [][]string
	for {

		record, err := reader.Read()
		// if we've reached the end of the file, break
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
		}

		records = append(records, record)
	}
	// remove the first line
	return records[1:], records[0]

}

// parseSong takes a slice of strings and returns a slice of song struct
func parseSong(records [][]string) []Song {
	var songs []Song
	for _, record := range records {
		song := Song{
			Id:               record[0],
			Name:             record[1],
			Artists:          record[2],
			Danceability:     stringToFloat(record[3]),
			Energy:           stringToFloat(record[4]),
			Key:              stringToInt(record[5]),
			Loudness:         stringToFloat(record[6]),
			Mode:             stringToInt(record[7]),
			Speechiness:      stringToFloat(record[8]),
			Acousticness:     stringToFloat(record[9]),
			Instrumentalness: stringToFloat(record[10]),
			Liveness:         stringToFloat(record[11]),
			Valence:          stringToFloat(record[12]),
			Tempo:            stringToFloat(record[13]),
			Duration_ms:      stringToInt(record[14]),
			Time_signature:   stringToInt(record[15]),
		}
		songs = append(songs, song)
	}
	return songs
}
