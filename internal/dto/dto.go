package dto

type EnderecoViaCEP_DTO struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"localidade"`
	Estado      string `json:"uf"`
	Regiao      string `json:"regiao"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

type EnderecoBrasilAPI_DTO struct {
	CEP     string `json:"cep"`
	Estado  string `json:"state"`
	Cidade  string `json:"city"`
	Bairro  string `json:"neighborhood"`
	Rua     string `json:"street"`
	Service string `json:"service"`
}
