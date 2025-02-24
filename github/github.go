package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	login := "mikebellcoder"

	n, nr, err := githubInfo(login)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("Name: %s, NumRepos: %d\n", n, nr)
}

// githubInfo returns name and number of public repos for login
func githubInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf(resp.Status)
	}

	defer resp.Body.Close()

	var r struct { // anonymous struct
		Name     string
		NumRepos int `json:"public_repos"`
	}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", 0, err
	}

	return r.Name, r.NumRepos, nil
}
