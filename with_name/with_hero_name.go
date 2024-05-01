package with_name

import (
	"encoding/json"
	"log"
	"os"
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

	heroAbilitiesFile, err := os.ReadFile("hero_abilities.json")
	if err != nil {
		log.Fatal(err)
	}

	var heroAbilitiesStruct HeroAbilities
	err = json.Unmarshal(heroAbilitiesFile, &heroAbilitiesStruct)
	if err != nil {
		log.Fatal(err)
	}

	var abilitiesMap map[string]MapAbility
	abilitiesFile, err := os.ReadFile("abilities.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(abilitiesFile, &abilitiesMap)
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
