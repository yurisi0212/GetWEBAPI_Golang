package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type PokeAPI struct {
	Height int    `json:"height"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Types  []struct {
		Type struct {
			Name string `json:"name"`
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

	fmt.Printf("--%s--\n", apires.Name)
	fmt.Printf("ID %s\n", strconv.Itoa(apires.ID))
	fmt.Printf("大きさ %scm, 重さ %skg\n", strconv.Itoa(apires.Weight), strconv.Itoa(apires.Height))
	fmt.Printf("タイプ %s", apires.Types[0].Type.Name)

}
