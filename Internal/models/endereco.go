package models

import (
	"encoding/json"
	"net/http"

	"desafiomulti/internal/dto"
)

type Endereco struct {
	CEP    string `json:"cep"`
	Estado string `json:"estado"`
	Cidade string `json:"cidade"`
	Bairro string `json:"bairro"`
	Rua    string `json:"rua"`
}

func GetEnderecoByCepViaCEP(cep string) (Endereco, error) {
	enderecoViaCEP, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return Endereco{}, err
	}
	defer enderecoViaCEP.Body.Close()

	var enderecoViaCEP_DTO dto.EnderecoViaCEP_DTO
	if err := json.NewDecoder(enderecoViaCEP.Body).Decode(&enderecoViaCEP_DTO); err != nil {
		return Endereco{}, err
	}

	return Endereco{
		CEP:    cep,
		Estado: enderecoViaCEP_DTO.Estado,
		Cidade: enderecoViaCEP_DTO.Cidade,
		Bairro: enderecoViaCEP_DTO.Bairro,
		Rua:    enderecoViaCEP_DTO.Logradouro,
	}, nil
}

func GetEnderecoByCepBrasilAPI(cep string) (Endereco, error) {
	enderecoBrasilAPI, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		return Endereco{}, err
	}
	defer enderecoBrasilAPI.Body.Close()

	var enderecoBrasilAPI_DTO dto.EnderecoBrasilAPI_DTO
	if err := json.NewDecoder(enderecoBrasilAPI.Body).Decode(&enderecoBrasilAPI_DTO); err != nil {
		return Endereco{}, err
	}

	return Endereco{
		CEP:    cep,
		Estado: enderecoBrasilAPI_DTO.Estado,
		Cidade: enderecoBrasilAPI_DTO.Cidade,
		Bairro: enderecoBrasilAPI_DTO.Bairro,
		Rua:    enderecoBrasilAPI_DTO.Rua,
	}, nil
}
