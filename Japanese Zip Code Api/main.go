package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Digite o seu Zip Code (CEP)")

	var zipCode string

	fmt.Scanln(&zipCode)

	resp, err := http.Get("https://zipcloud.ibsnet.co.jp/api/search?zipcode=" + zipCode)

	if err != nil {
		fmt.Println(err)
	}

	var respApi map[string]any

	json.NewDecoder(resp.Body).Decode(&respApi)

	var result []any

	for key, value := range respApi {

		if key == "results" {

			valueResult := value.([]any)

			for _, valueMap := range valueResult {

				resultValueMap := valueMap.(map[string]any)

				for keyFinal, valueFinal := range resultValueMap {

					switch keyFinal {

					case "address1":
						result = append(result, keyFinal+": "+valueFinal.(string))

					case "address2":
						result = append(result, keyFinal+": "+valueFinal.(string))

					case "address3":
						result = append(result, keyFinal+": "+valueFinal.(string))

					case "kana1":
						result = append(result, keyFinal+": "+valueFinal.(string))

					case "kana2":
						result = append(result, keyFinal+": "+valueFinal.(string))

					case "kana3":
						result = append(result, keyFinal+": "+valueFinal.(string))

					case "prefcode":
						result = append(result, keyFinal+": "+valueFinal.(string))

					case "zipcode":
						result = append(result, keyFinal+": "+valueFinal.(string))

					}

				}

			}

		}

	}

	fmt.Println(result)

}
