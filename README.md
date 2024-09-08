# Start.gg Seeding Uploader

Go Script for Seeding Phases to start.gg API

This script imports seed mappings from a Google Sheets CSV into a specific phase using start.gg's GraphQL API. Below is an overview of how it works and instructions on setting it up.

## Table of Contents
1. [Overview](#overview)
2. [Requirements](#requirements)
3. [How to Get Required IDs and Tokens](#how-to-get-required-ids-and-tokens)
4. [Setup and Usage](#setup-and-usage)
5. [Error Handling](#error-handling)
6. [Troubleshooting](#troubleshooting)

## Overview

This Go script reads seed mappings from a Google Sheets CSV and updates them in a tournament phase on start.gg via their GraphQL API.

- The script pulls data from a Google Sheet (in CSV format).
- It uses start.gg's API to update phase seeding.

## Requirements

- Go installed on your system
- A Google Sheets file containing the seed mappings
- start.gg API authentication token
- Phase ID (from start.gg)

## How to Get Required IDs and Tokens

### Phase ID
1. Go to the start.gg website and navigate to the tournament page.
2. Look for the phase of the tournament you want to update.
3. The Phase ID is visible in the URL as a part of the path, e.g., `https://start.gg/tournament/{tournament-name}/event/{event-name}/bracket/{phaseId}`.

### Google Sheets Key
1. Open the Google Sheets file with the seed data.
2. Copy the key from the URL, which is between `/d/` and `/edit` in the link:  
   Example:  
   `https://docs.google.com/spreadsheets/d/{sheetsKey}/edit`

### Auth Token
1. Navigate to [start.gg API](https://developer.start.gg/docs/authentication).
2. Create or log in to your start.gg developer account.
3. Generate a new API token.

## Setup and Usage

1. Clone the repository.
2. Set up the environment variables:
   - `phaseId` (replace with the actual Phase ID)
   - `sheetsKey` (replace with your Google Sheets key)
   - `authToken` (replace with your start.gg API token)
3. Run the script:

```bash
go run main.go
```

## Error Handling

If there are errors, the script will print the error messages returned by the API. Make sure:
- The Phase ID is valid.
- The Google Sheets CSV is properly formatted.
- The auth token has the correct permissions.

Example output for an error:

```json
{
  "errors": [
    {
      "message": "Invalid Phase ID",
      "locations": [{ "line": 2, "column": 3 }],
      "path": ["updatePhaseSeeding"],
      "extensions": {
        "code": "BAD_USER_INPUT"
      }
    }
  ]
}
```

If you see an error similar to the above, verify that the Phase ID and the input data are correct.

## Troubleshooting

- **Invalid Auth Token**: Check that your token is up-to-date and has the necessary permissions. You can regenerate the token via the start.gg API developer portal if needed.
  
- **Invalid Phase ID**: Ensure the Phase ID is correctly copied from the URL of the tournament's phase you want to update. If the Phase ID is incorrect, the API will return an error.

- **Incorrect CSV Format**: Ensure the CSV from Google Sheets has the expected structure. The script expects the third column to contain `SeedId` and the first column to contain `SeedNum`. Any mismatch in the columns could lead to errors.
  
- **Network Issues**: If the API or Google Sheets CSV cannot be accessed, check your internet connection or whether there is an issue with the respective services.

## Success Response

On success, the script will print:

```json
{
  "data": {
    "updatePhaseSeeding": {
      "id": "1696690"
    }
  }
}
```

This confirms that the seeds have been successfully imported to the phase.

