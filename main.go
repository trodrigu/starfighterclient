package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Order struct {
	Account   string `json:"account"`
	Venue     string `json:"venue"`
	Stock     string `json:"stock"`
	Price     int    `json:"price"`
	Quantity  int    `json:"qty"`
	Direction string `json:"direction"`
	OrderType string `json:"orderType"`
}

func main() {

	apiKey := "6097504d8f517b6edfad40627b3d2feab02f9b80"
	baseUrl := "https://api.stockfighter.io/ob/api"
	orderFile := "/Users/thomas/tmp/order.json"
	orderConfig, err := ioutil.ReadFile(orderFile)

	if err != nil {
		log.Fatal(err)
	}

	var orderJson Order

	err = json.Unmarshal(orderConfig, &orderJson)

	fullUrl := baseUrl + "/venues/" + orderJson.Venue + "/stocks/" + orderJson.Stock + "/orders"
	req, err := http.NewRequest("POST", fullUrl, bytes.NewReader(orderConfig))
	req.Header.Add("X-Starfighter-Authorization", apiKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	resp, _ := client.Do(req)
	fmt.Println(resp)
}
