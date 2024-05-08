# Go Application for Extracting Dota 2 Ability Data

This Go application is designed to extract and format Dota 2 ability data into a JSON file. The resulting JSON file will have the following format:

```json
{
  "ability_id": "Ability ID",
  "ability_name": "Ability Name",
  "hero_ability_display_name": "Hero and Ability Display Name"
}
```

The application uses existing constants as input data. This was necessary because the required data format was not readily available from the Stratz or OpenDota APIs.

## Purpose
The primary purpose of this application is to generate data for the Ability Draft Explorer Application. By providing a structured and consistent data format, it aids in the exploration and understanding of Dota 2 abilities.

## Usage
To use this application, follow these steps:

Ensure you have Go installed on your machine.
Clone this repository.
Run the application with the command go run main.go.
Please note that this application requires access to the Dota 2 constants file. Ensure this file is available and correctly formatted for the application to run successfully.

