package main

import (
	"io"
	"log"
	"net/http"
	"sync"
)

const (
	cep              = "35675-000"
	brasilapiAddress = "https://brasilapi.com.br/api/cep/v1/" + cep
	viacepAdress     = "http://viacep.com.br/ws/" + cep + "/json/"
)

func call(url string, wg *sync.WaitGroup) {
	var client http.Client
	res, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(body)
		log.Print(bodyString)
	} else {
		log.Printf("Erro: status %d", res.StatusCode)
	}
	wg.Done()
}

func main() {
	waitgroup := sync.WaitGroup{}
	waitgroup.Add(1)
	defer waitgroup.Wait()

	go call(brasilapiAddress, &waitgroup)
	go call(viacepAdress, &waitgroup)
}
