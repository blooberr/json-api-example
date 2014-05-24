package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Prices struct {
	Amount   float64 `json:"amount,string"`
	Currency string  `json:"currency"`
}

func main() {
	url := "https://coinbase.com/api/v1/prices/spot_rate"
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	p := Prices{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v \n", p)
}
