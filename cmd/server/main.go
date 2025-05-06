package main

import (
	"fmt"
	"os"
	"time"

	"github.com/riccruzdev/desafiomulti/internal/models"
)

func main() {
	// Define default CEP
	cep := "01153000"

	// If CEP is provided as command-line argument, use it instead
	if len(os.Args) > 1 {
		cep = os.Args[1]
	}

	ch1 := make(chan models.Endereco)
	ch2 := make(chan models.Endereco)

	//API BrasilAPI
	go func() {
		//time.Sleep(time.Second * 1)
		enderecoBrasilAPI, err := models.GetEnderecoByCepBrasilAPI(cep)
		if err != nil {
			fmt.Println(err)
		}
		ch1 <- enderecoBrasilAPI
	}()

	//API ViaCEP
	go func() {
		//time.Sleep(time.Second * 1)
		enderecoViaCEP, err := models.GetEnderecoByCepViaCEP(cep)
		if err != nil {
			fmt.Println(err)
		}
		ch2 <- enderecoViaCEP
	}()

	select {
	case msg := <-ch1: //BrasilAPI
		fmt.Printf("Received from %s. CEP: %s. Estado: %s. Cidade: %s. Bairro: %s. Rua: %s\n", "BrasilAPI", msg.CEP, msg.Estado, msg.Cidade, msg.Bairro, msg.Rua)
	case msg := <-ch2: //ViaCEP
		fmt.Printf("Received from %s. CEP: %s. Estado: %s. Cidade: %s. Bairro: %s. Rua: %s\n", "ViaCEP", msg.CEP, msg.Estado, msg.Cidade, msg.Bairro, msg.Rua)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout.")
	}
}
