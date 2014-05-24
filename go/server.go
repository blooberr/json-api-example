package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
)

type Response struct {
	Hello string `json:"hello"`
}

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		r := Response{Hello: "world"}
		encoded, err := json.Marshal(r)
		if err != nil {
			return "json.marshall error"
		} else {
			return string(encoded)
		}
	})
	m.Run()

}
