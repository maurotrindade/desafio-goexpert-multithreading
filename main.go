package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

const (
	cep              = "35675-000"
	brasilapiAddress = "https://brasilapi.com.br/api/cep/v1/" + cep
	viacepAdress     = "http://viacep.com.br/ws/" + cep + "/json/"
)

func call(url string, ch chan string) {
	var client http.Client
	var body []byte

	res, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		body, err = io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("Erro: status %d", res.StatusCode)
		return
	}
	ch <- string(body)
}

func printResult(bodyStr, url string) {
	log.Println("API utilizada: " + url)
	log.Println("Resposta:")
	log.Println(bodyStr)
}

func main() {
	channelBrasilApi := make(chan string)
	channelViacep := make(chan string)

	go call(brasilapiAddress, channelBrasilApi)
	go call(viacepAdress, channelViacep)

	select {
	case msg := <-channelBrasilApi:
		printResult(msg, brasilapiAddress)
	case msg := <-channelViacep:
		printResult(msg, viacepAdress)
	case <-time.After(1 * time.Second):
		println("Error: timeout")
	}
}
