package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemon struct {
	Name            string           `json:"name"`
	ID              int              `json:"id"`
	Height          float64          `json:"height"`
	Base_experience float64          `json:"base_experience"`
	Weight          float64          `json:"weight"`
	Abilities       []PokemonAbility `json:"abilities"`
	Cries           []PokemonCry     `json:"cries"`
	Forms           []PokemonForm    `json:"forms"`
	Moves           []PokemonMove    `json:"moves"`
	Stats           []PokemonStat    `json:"stats"`
	Types           []PokemonType    `json:"types"`
}

type PokemonAbility struct {
	Ability   AbilityDetail `json:"ability"`
	Is_hidden bool          `json:"is_hidden"`
	Slot      int           `json:"slot"`
}

type AbilityDetail struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonCry struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type PokemonForm struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonMove struct {
	Move PokemonMoveDetail `json:"move"`
}

type PokemonMoveDetail struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonStat struct {
	Base_stat int        `json:"base_stat"`
	Stat      StatDetail `json:"stat"`
}

type StatDetail struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonType struct {
	PokeType PokemonTypeDetail `json:"type"`
}

type PokemonTypeDetail struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {

	var idPokemon string

	fmt.Println("Digite o id do Pokémon:")

	fmt.Scanln(&idPokemon)

	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + idPokemon)

	if err != nil {
		panic(err)
	}

	var newPokemon Pokemon

	json.NewDecoder(resp.Body).Decode(&newPokemon)

	fmt.Println("Nome:", newPokemon.Name)
	fmt.Println("ID:", newPokemon.ID)
	fmt.Println("Experiência base:", newPokemon.Base_experience)

	fmt.Printf("Habilidades:\n")

	for index, value := range newPokemon.Abilities {
		fmt.Printf("Habilidade %d:\n", index+1)
		fmt.Printf("  name: %s\n", value.Ability.Name)
		fmt.Printf("  url: %s\n", value.Ability.Url)
		fmt.Printf("  is_hidden: %t\n", value.Is_hidden)
		fmt.Printf("  slot: %d\n", value.Slot)
	}

	fmt.Printf("Golpes:\n")

	for index, value := range newPokemon.Moves {
		fmt.Printf("Habilidade %d:\n", index+1)
		fmt.Printf("name: %s\n", value.Move.Name)
		fmt.Printf("url: %s\n", value.Move.Url)
	}

	fmt.Printf("Formas:\n")

	for index, value := range newPokemon.Forms {
		fmt.Printf("Forma %d:\n", index+1)
		fmt.Printf("name: %s\n", value.Name)
		fmt.Printf("url: %s\n", value.Url)
	}

	fmt.Printf("Stats:\n")

	for _, value := range newPokemon.Stats {
		fmt.Printf("%s\n", value.Stat.Name)
		fmt.Printf("Valor: %d\n", value.Base_stat)
	}

	fmt.Printf("Tipos:\n")

	for index, value := range newPokemon.Types {
		fmt.Printf("Tipo %d:\n", index+1)
		fmt.Printf("%s\n", value.PokeType.Name)
		fmt.Printf("Url: %s\n", value.PokeType.Url)
	}

	fmt.Printf("Altura:\n")

	fmt.Println(newPokemon.Height)

	fmt.Printf("Peso:\n")

	fmt.Println(newPokemon.Weight)

}
