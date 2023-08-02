package main

import (
	"fmt"
	"log"
	"net/http"
)

func createClientWithTracing() (*http.Client, error) {
	return nil, nil
}

func createDefaultClient() (*http.Client, error) {
	return nil, nil
}

func main() {
	var client *http.Client
	var tracing bool
	if tracing {
		fmt.Println("11111")
		client, err := createClientWithTracing()
		if err != nil {
			fmt.Println("error")
		}
		log.Println(client)
	} else {
		fmt.Println("22222")
		client, err := createDefaultClient()
		if err != nil {
			fmt.Println("error")
		}
		log.Println(client)
	}
	fmt.Println(client)
}
