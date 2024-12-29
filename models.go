package main

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
