package main

import (
	"com/github/fastfoodfinance/scraper/internal"
	"com/github/fastfoodfinance/scraper/internal/ubereats"
	"encoding/json"
	"log"
)

func main() {
	menus := []*internal.Menu{}
	menus = append(menus, ubereats.Menus()...)

	json, err := json.Marshal(menus)
	if err != nil {
		log.Fatalln(err)
	}

	str := string(json)
	log.Println(str)
}
