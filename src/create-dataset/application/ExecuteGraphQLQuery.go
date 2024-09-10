package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func ExecuteGraphQLQuery(query string) (*GitHubGraphQLRepositoriesResponseDto, error) {
	accessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	if accessToken == "" {
		return nil, fmt.Errorf("personal accces token not found.")
	}

	jsonMapInstance := map[string]string{
		"query": query,
	}

	jsonResult, err := json.Marshal(jsonMapInstance)
	if err != nil {
		return nil, fmt.Errorf("there was an error marshaling the JSON instance %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonResult))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	if err != nil {
		return nil, fmt.Errorf("Error on create http request: %v", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error on send http request: %v", err)
	}
	defer resp.Body.Close()

	jsonBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error on read http response: %v", err)
	}

	var GitHubGraphQLRepositoriesResponseDto GitHubGraphQLRepositoriesResponseDto
	GitHubGraphQLRepositoriesResponseDtoErr := json.Unmarshal(jsonBody, &GitHubGraphQLRepositoriesResponseDto)
	if err != nil {
		log.Fatalf("json.Unmarshal fail: %v", GitHubGraphQLRepositoriesResponseDtoErr)
	}

	return &GitHubGraphQLRepositoriesResponseDto, nil
}
