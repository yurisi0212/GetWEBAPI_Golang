package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PokeAPI struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices            []interface{} `json:"game_indices"`
	Height                 int           `json:"height"`
	HeldItems              []interface{} `json:"held_items"`
	ID                     int           `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name      string        `json:"name"`
	Order     int           `json:"order"`
	PastTypes []interface{} `json:"past_types"`
	Species   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string      `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        string      `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault interface{} `json:"front_default"`
				FrontFemale  interface{} `json:"front_female"`
			} `json:"dream_world"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
			} `json:"official-artwork"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault  interface{} `json:"back_default"`
					BackGray     interface{} `json:"back_gray"`
					FrontDefault interface{} `json:"front_default"`
					FrontGray    interface{} `json:"front_gray"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault  interface{} `json:"back_default"`
					BackGray     interface{} `json:"back_gray"`
					FrontDefault interface{} `json:"front_default"`
					FrontGray    interface{} `json:"front_gray"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault  interface{} `json:"back_default"`
					BackShiny    interface{} `json:"back_shiny"`
					FrontDefault interface{} `json:"front_default"`
					FrontShiny   interface{} `json:"front_shiny"`
				} `json:"crystal"`
				Gold struct {
					BackDefault  interface{} `json:"back_default"`
					BackShiny    interface{} `json:"back_shiny"`
					FrontDefault interface{} `json:"front_default"`
					FrontShiny   interface{} `json:"front_shiny"`
				} `json:"gold"`
				Silver struct {
					BackDefault  interface{} `json:"back_default"`
					BackShiny    interface{} `json:"back_shiny"`
					FrontDefault interface{} `json:"front_default"`
					FrontShiny   interface{} `json:"front_shiny"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault interface{} `json:"front_default"`
					FrontShiny   interface{} `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  interface{} `json:"back_default"`
					BackShiny    interface{} `json:"back_shiny"`
					FrontDefault interface{} `json:"front_default"`
					FrontShiny   interface{} `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  interface{} `json:"back_default"`
					BackShiny    interface{} `json:"back_shiny"`
					FrontDefault interface{} `json:"front_default"`
					FrontShiny   interface{} `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      interface{} `json:"back_default"`
					BackFemale       interface{} `json:"back_female"`
					BackShiny        interface{} `json:"back_shiny"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontDefault     interface{} `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       interface{} `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      interface{} `json:"back_default"`
					BackFemale       interface{} `json:"back_female"`
					BackShiny        interface{} `json:"back_shiny"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontDefault     interface{} `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       interface{} `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      interface{} `json:"back_default"`
					BackFemale       interface{} `json:"back_female"`
					BackShiny        interface{} `json:"back_shiny"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontDefault     interface{} `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       interface{} `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      interface{} `json:"back_default"`
						BackFemale       interface{} `json:"back_female"`
						BackShiny        interface{} `json:"back_shiny"`
						BackShinyFemale  interface{} `json:"back_shiny_female"`
						FrontDefault     interface{} `json:"front_default"`
						FrontFemale      interface{} `json:"front_female"`
						FrontShiny       interface{} `json:"front_shiny"`
						FrontShinyFemale interface{} `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string      `json:"back_default"`
					BackFemale       interface{} `json:"back_female"`
					BackShiny        string      `json:"back_shiny"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontDefault     string      `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       string      `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     interface{} `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       interface{} `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     interface{} `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       interface{} `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string      `json:"front_default"`
					FrontFemale  interface{} `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     string      `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       string      `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string      `json:"front_default"`
					FrontFemale  interface{} `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func main() {
	uri := "https://pokeapi.co/api/v2/pokemon/magikarp"
	req, err := http.Get(uri)
	if err != nil {
		log.Fatal("Get Http Error:", err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal("IO Read Error:", err)
	}

	defer req.Body.Close()
	var apires PokeAPI
	if err := json.Unmarshal(body, &apires); err != nil {
		panic(err)
	}

	fmt.Println(apires.Species.Name)
}
