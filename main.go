package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type People struct {
	Name   string   `json:"name"`
	Height string   `json:"height"`
	Film   []string `json:"films"`
}

func main() {
	args := os.Args

	if len(args) < 1 {
		fmt.Println("Usage: ./main.go <url>")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("URL is in invalid format: %s\n", err)
		os.Exit(1)
	}

	response, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP Code %d): %s\n", response.StatusCode, body)
		os.Exit(1)
	}

	var people People

	err = json.Unmarshal(body, &people)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("JSON Parsed\nName: %s\nHeight: %s\nFilms: %v", people.Name, people.Height, strings.Join(people.Film, ", "))
}
