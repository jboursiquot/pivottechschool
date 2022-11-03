package main

import (
	"errors"
	"net/http"
)

var ErrNotFound = errors.New("Not found")

const baseURL = "https://gateway.marvel.com"

// makeRequest is a helper function that you can finish writing and use to make
// the actual API calls. You are welcome to change this function's signature if
// you so desire.
func makeRequest(url string) []byte {
	// Fill in the correct arguments here
	req, _ := http.NewRequest("fixme", "fixme", nil)

	// Leave this part alone to handle adding the necessary auth parameters
	q := addQueryAuth(req.URL.Query())
	req.URL.RawQuery = q.Encode()

	// here, you'd need to actually send the request, read the response, and return it
	return nil
}

type MarvelClient interface {
	// GetFirst25 returns the names of the first 25 characters in the catalog, ordered
	// by name.
	GetFirst25() []string

	// GetNameByID takes a character ID such as 1009150 and returns its name ("Agent Zero").
	GetNameByID(id int) (string, error)

	// GetCharacterDetail takes the name of a character and returns a populated CharacterDetail struct
	// if the character is found, otherwise an empty struct and ErrNotFound.
	GetCharacterDetail(name string) (CharacterDetail, error)

	// GetCharactersInComic takes an ID of a comic (such as 41112) and returns a slice of characters in the comic.
	GetCharactersInComic(comicID int) ([]string, error)
}

type CharacterDetail struct {
	ID           int
	Description  string
	ThumbnailURL string
	ComicCount   int
}

// Client is your client implementation.
type Client struct{}

func (c *Client) GetFirst25() []string {
	return nil
}

func (c *Client) GetNameByID(id int) (string, error) {
	return "", errors.New("not implemented")
}

func (c *Client) GetCharacterDetail(name string) (CharacterDetail, error) {
	return CharacterDetail{}, errors.New("not implemented")
}

func (c *Client) GetCharactersInComic(comicID int) ([]string, error) {
	return nil, errors.New("not implemented")
}
