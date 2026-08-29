package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	var option int

	var optionChosen string

	fmt.Println("Opções:  films(1) , people (2), planets(3), species(4), starships(5), vehicles(6)")

	fmt.Scanln(&option)

	fmt.Println("Opções:  id(1) , texto(2), mostrar todos da opção escolhida anteriormente (3)")

	var typeSearch int

	var idSearch int

	var textSearch string

	fmt.Scanln(&typeSearch)

	switch typeSearch {

	case 1:
		fmt.Println("Digite o que deseja pesquisar")
		fmt.Scanln(&idSearch)
	case 2:
		fmt.Println("Digite o que deseja pesquisar")
		fmt.Scanln(&textSearch)
	case 3:
		fmt.Println("Exibir todos do tipo escolhido")
	}

	switch option {

	case 1:
		optionChosen = "https://swapi.dev/api/films/"
		if idSearch > 0 {
			optionChosen = optionChosen + strconv.Itoa(idSearch)
		}
	case 2:
		optionChosen = "https://swapi.dev/api/people/"
		if idSearch > 0 {
			optionChosen = optionChosen + strconv.Itoa(idSearch)
		}
	case 3:
		optionChosen = "https://swapi.dev/api/planets/"
		if idSearch > 0 {
			optionChosen = optionChosen + strconv.Itoa(idSearch)
		}
	case 4:
		optionChosen = "https://swapi.dev/api/species/"
		if idSearch > 0 {
			optionChosen = optionChosen + strconv.Itoa(idSearch)
		}
	case 5:
		optionChosen = "https://swapi.dev/api/starships/"
		if idSearch > 0 {
			optionChosen = optionChosen + strconv.Itoa(idSearch)
		}
	case 6:
		optionChosen = "https://swapi.dev/api/vehicles/"
		if idSearch > 0 {
			optionChosen = optionChosen + strconv.Itoa(idSearch)
		}
	default:
		fmt.Println("Selecione uma das 6 opções!")
		return

	}

	resp, err := http.Get(optionChosen)

	if err != nil {
		fmt.Println(err.Error())
	}

	var respApi map[string]any

	json.NewDecoder(resp.Body).Decode(&respApi)

	var nextPage string

	nextPage = ""

	var result []any

	for index, value := range respApi {

		if index == "results" {

			result = append(result, map[string]any{
				index: value,
			})
		}

		if index == "next" && value != nil {

			nextPage = value.(string)

			for {

				if nextPage == "" {
					break
				}

				respNew, err := http.Get(nextPage)

				if err != nil {
					fmt.Println(err.Error())
				}

				var respApiNewPage map[string]any

				json.NewDecoder(respNew.Body).Decode(&respApiNewPage)

				for index, value := range respApiNewPage {

					if index == "results" {

						result = append(result, map[string]any{
							index: value,
						})
					}

					if index == "next" && value != nil {

						nextPage = value.(string)

					}

					if index == "next" && value == nil {

						nextPage = ""

					}

				}

			}

		}

		if idSearch > 0 {

			result = append(result, map[string]any{
				index: value,
			})

		}

	}

	if textSearch != "" {

		for _, valueResult := range result {

			var mapResult = valueResult.(map[string]any)

			for indexMapResult, valueMapResult := range mapResult {

				if indexMapResult == "results" {

					var mapResultValue = valueMapResult.([]any)

					for _, valueResultValue := range mapResultValue {

						var mapResultValueInner = valueResultValue.(map[string]any)

						for _, valueResultValueInner := range mapResultValueInner {

							switch v := valueResultValueInner.(type) {
							case string:

								if strings.Contains(v, textSearch) {

									fmt.Println(valueResultValue)

								}

							}

						}

					}

				}

			}

		}

	}

	if idSearch > 0 {

		fmt.Println(result)

	}

}
