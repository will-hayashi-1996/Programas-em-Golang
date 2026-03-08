package main

import (
		"fmt"
		"net/http"
		"encoding/json"
	)

func main() {

	var cep string

	fmt.Println("Insira seu CEP")

	fmt.Scanln(&cep)

	res,err := http.Get("https://viacep.com.br/ws/"+cep+"/json/")

	if(err != nil){
		fmt.Println("Erro na api!")
	}

	if(res.StatusCode == 200){

		var mapAPI map[string]any

		json.NewDecoder(res.Body).Decode(&mapAPI)

		if(err != nil){
			fmt.Println(err)
			return
		}

		fmt.Println(mapAPI)
	}

}
