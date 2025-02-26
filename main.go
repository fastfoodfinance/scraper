package main

import (
	"encoding/json"
	"log"

	"github.com/fastfoodfinance/scraper/internal"
	"github.com/fastfoodfinance/scraper/internal/ubereats"
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
