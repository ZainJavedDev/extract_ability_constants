package main

import (
	"Zain/get_abilities/with_name"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Ability struct {
	Dname string `json:"dname"`
}

type FinalAbilities struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

func main() {

	with_name.WithName()
	os.Exit(1)

	Extract()
}

func Extract() {
	abilityIdFile, err := os.ReadFile("ability_ids.json")
	if err != nil {
		log.Fatal(err)
	}
	var abilityIdMap map[string]string
	err = json.Unmarshal(abilityIdFile, &abilityIdMap)
	if err != nil {
		log.Fatal(err)
	}

	var abilitiesMap map[string]Ability
	abilitiesFile, err := os.ReadFile("abilities.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(abilitiesFile, &abilitiesMap)
	if err != nil {
		log.Fatal(err)
	}

	var finalAbilities []FinalAbilities

	for id, name := range abilityIdMap {
		if ability, ok := abilitiesMap[name]; ok {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal(err)
			}

			finalAbilities = append(finalAbilities, FinalAbilities{
				Id:          idInt,
				Name:        name,
				DisplayName: ability.Dname,
			})
		}
	}

	finalAbilitiesJson, err := json.Marshal(finalAbilities)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("final_abilities.json", finalAbilitiesJson, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
