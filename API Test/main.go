package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	var cep string

	fmt.Println("Insira seu CEP")

	fmt.Scanln(&cep)

	res, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")

	if err != nil {
		fmt.Println("Erro na api!")
	}

	if res.StatusCode == 200 {

		var mapAPI map[string]any

		json.NewDecoder(res.Body).Decode(&mapAPI)

		if err != nil {
			fmt.Println(err)
			return
		}

		for key, value := range mapAPI {

			if value != nil && value != "" && value != " " {
				fmt.Println(key + ": " + value.(string))
			} else {
				fmt.Println(key + ": Vazio!")
			}

		}
	}

}
