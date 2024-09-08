package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SeedMapping struct {
	SeedId  string `json:"seedId"`
	SeedNum string `json:"seedNum"`
}

type GraphQLRequest struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"`
}

func main() {
	// Replace with your Phase ID
	const phaseId = 1696690
	// Replace with your Google Sheets key
	const sheetsKey = "1upZq9QCvNK9jPlx6LB7tNt9TecUR_QU89P0mbVBZ9nw"
	// Replace with your actual API token
	const authToken = ""
	const apiVersion = "alpha"

	// URL to download the CSV from Google Sheets
	url := "https://docs.google.com/spreadsheets/d/" + sheetsKey + "/export?format=csv"
	resp, err := http.Get(url)
	if err != nil {
		// Error handling for failed GET request
		panic(err)
	}
	defer resp.Body.Close()

	// Parse CSV
	r := csv.NewReader(resp.Body)
	records, err := r.ReadAll()
	if err != nil {
		// Error handling for CSV parsing
		panic(err)
	}

	var seedMappings []SeedMapping
	for i, record := range records {
		// Skip header row
		if i == 0 {
			continue
		}
		// Populate SeedMapping struct
		seedMappings = append(seedMappings, SeedMapping{
			SeedId:  record[2], // Make sure column 2 has seed ID, , or change the number 0 to the column with your seed IDs from your sheet (Column A is 0, B is 1, etc.)
			SeedNum: record[0], // Make sure column 0 has seed number, or change the number 0 to the column with your seed number from your sheet (Column A is 0, B is 1, etc.)
		})
	}

	numSeeds := len(seedMappings)
	fmt.Printf("Importing %d seeds to phase %d...\n", numSeeds, phaseId)

	// Prepare GraphQL variables for API request
	variables := map[string]interface{}{
		"phaseId":     phaseId,
		"seedMapping": seedMappings,
	}

	// GraphQL mutation for updating phase seeding
	gqlRequest := GraphQLRequest{
		Query: `
		mutation UpdatePhaseSeeding ($phaseId: ID!, $seedMapping: [UpdatePhaseSeedInfo]!) {
			updatePhaseSeeding (phaseId: $phaseId, seedMapping: $seedMapping) {
				id
			}
		}
		`,
		Variables: variables,
	}

	// Serialize GraphQL request to JSON
	reqBodyBytes, err := json.Marshal(gqlRequest)
	if err != nil {
		// Error handling for JSON marshaling
		panic(err)
	}

	// API endpoint for the request
	endpoint := "https://api.start.gg/gql/" + apiVersion
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(reqBodyBytes))
	if err != nil {
		// Error handling for creating the HTTP request
		panic(err)
	}
	// Set headers for the API request
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		// Error handling for request execution
		panic(err)
	}
	defer resp.Body.Close()

	// Read and process the response
	bodyBytes, _ := io.ReadAll(resp.Body) 
	var result map[string]interface{}
	json.Unmarshal(bodyBytes, &result)

	// Check for errors in the response
	if _, ok := result["errors"]; ok {
		fmt.Println("Error:")
		fmt.Println(result["errors"]) // Output errors for debugging
	} else {
		fmt.Println("Success!") // Success message
	}
}
