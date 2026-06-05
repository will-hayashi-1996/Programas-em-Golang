package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	var idPokemon string

	fmt.Println("Digite o id do pokémon")

	fmt.Scanln(&idPokemon)

	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + idPokemon)

	if err != nil {

		fmt.Println("Erro:", err)

	}

	var resp map[string]any

	var result []any

	json.NewDecoder(res.Body).Decode(&resp)

	for key := range resp {

		if key == "name" {
			//fmt.Println("Nome do Pokémon:")
			//fmt.Println(resp["name"])
			result = append(result, "Nome do Pokémon:"+resp["name"].(string))
		} else if key == "moves" {
			moves := resp["moves"].([]any)

			//fmt.Println("Nome dos  Golpes desse Pokémon:")

			for _, valueMove := range moves {

				valueMoveMap := valueMove.(map[string]any)

				move := valueMoveMap["move"].(map[string]any)

				for keyName, _ := range move {

					//valueNameMap := valueName.(string)

					//name := valueNameMap["name"].(string)

					if keyName == "name" {

						//fmt.Println("Golpe:")
						//fmt.Println(move["name"])
						result = append(result, "Golpe:"+move["name"].(string))

					}

				}

			}
		} else if key == "stats" {

			stats := resp["stats"].([]any)

			for _, valueStats := range stats {
				statMap := valueStats.(map[string]any)

				statName := statMap["stat"].(map[string]any)
				name := statName["name"].(string)

				baseStat := statMap["base_stat"].(float64)
				strConvert := strconv.FormatFloat(baseStat, 'f', 2, 64)

				result = append(result, "Status:"+name+"\n"+strConvert)
			}

		} else if key == "abilities" {

			abilities := resp["abilities"].([]any)

			for keyAbilities, valueAbilities := range abilities {

				abilityMap := valueAbilities.(map[string]any)

				abilityName := abilityMap["ability"].(map[string]any)

				ability := abilityName["name"].(string)

				abilityIndex := strconv.Itoa(keyAbilities + 1)

				result = append(result, "Habilidade "+abilityIndex+": "+ability)

			}

		} else if key == "types" {

			types := resp["types"].([]any)

			for KeyType, valueType := range types {

				typeMap := valueType.(map[string]any)

				//typeIndex := typeMap["slot"].(float64)

				//strConvert := strconv.FormatFloat(typeIndex, 'f', 2, 64)

				indexIndex := strconv.Itoa(KeyType + 1)

				typeName := typeMap["type"].(map[string]any)

				typeResult := typeName["name"].(string)

				result = append(result, "Tipo "+indexIndex+" Do Pokémon: "+typeResult)
			}
		} else if key == "weight" {

			weight := resp["weight"].(float64)

			strConvert := strconv.FormatFloat(weight, 'f', 2, 64)

			result = append(result, "Peso do Pokémon: "+strConvert)

		}

	}

	for _, valueResult := range result {

		//fmt.Println(keyResult)
		fmt.Println(valueResult)

	}

}
