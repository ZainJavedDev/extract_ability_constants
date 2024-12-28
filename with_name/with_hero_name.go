package with_name

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type HeroAbilities struct {
	Data Data `json:"data"`
}

type Data struct {
	Constants Constants `json:"constants"`
}

type Constants struct {
	Heroes []Hero `json:"heroes"`
}

type Hero struct {
	Id          int       `json:"id"`
	DisplayName string    `json:"displayName"`
	Abilities   []Ability `json:"abilities"`
}

type Ability struct {
	Ability AbilityDetail `json:"ability"`
}

type AbilityDetail struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type FinalAbilities struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

type MapAbility struct {
	Dname string `json:"dname"`
}

func WithName() {

	// Define the GraphQL query
	query := `{
		"query": "query { constants { heroes { id displayName abilities { ability { id name } } } } }"
	}`

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiToken := os.Getenv("STRATZ_API_TOKEN")

	// Create a request to the GraphQL API
	req, err := http.NewRequest("POST", "https://api.stratz.com/graphql", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "STRATZ_API")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Decode the response into heroAbilitiesStruct
	var heroAbilitiesStruct HeroAbilities
	err = json.Unmarshal(bodyBytes, &heroAbilitiesStruct)
	if err != nil {
		log.Fatal(err)
	}

	var abilitiesMap map[string]MapAbility
	resp, err = http.Get("https://api.opendota.com/api/constants/abilities")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&abilitiesMap)
	if err != nil {
		log.Fatal(err)
	}

	var finalAbilities []FinalAbilities

	for _, hero := range heroAbilitiesStruct.Data.Constants.Heroes {
		for _, ability := range hero.Abilities {
			if ability.Ability.Name == "generic_hidden" {
				continue
			}
			if mapAbility, ok := abilitiesMap[ability.Ability.Name]; ok {

				finalAbilities = append(finalAbilities, FinalAbilities{
					Id:          ability.Ability.Id,
					Name:        ability.Ability.Name,
					DisplayName: hero.DisplayName + " " + mapAbility.Dname,
				})
			}
		}
	}
	finalAbilitiesJson, err := json.Marshal(finalAbilities)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("final_hero_abilities.json", finalAbilitiesJson, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
