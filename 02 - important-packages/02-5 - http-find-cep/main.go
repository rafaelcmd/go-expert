package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
		}
		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Printf("Error parsing response: %v\n", err)
		}
		fmt.Println(data)
		file, err := os.Create("cep.json")
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s\nLogradouro: %s\nComplemento: %s\nUnidade: %s\nBairro: "+
			"%s\nLocalidade: %s\nUF: %s\nEstado: %s\nRegião: %s\nIBGE: %s\nGIA: %s\nDDD: %s\nSIAFI: %s\n",
			data.Cep, data.Logradouro, data.Complemento, data.Unidade, data.Bairro, data.Localidade, data.Uf,
			data.Estado, data.Regiao, data.Ibge, data.Gia, data.Ddd, data.Siafi))
	}
}
