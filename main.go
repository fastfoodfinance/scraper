package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/fastfoodfinance/scraper/internal"
	"github.com/fastfoodfinance/scraper/internal/ubereats"
)

func writeFile(menus []*internal.Menu) {
	json, err := json.Marshal(menus)
	if err != nil {
		log.Fatalln("failed to marshal json", err)
	}

	outputDir := "output"
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalln("failed to create directory", err)
	}

	outputFile := filepath.Join(outputDir, "menus.json")
	err = os.WriteFile(outputFile, json, os.ModePerm)
	if err != nil {
		log.Fatalln("failed to write to file", err)
	}
}

func main() {
	menus := []*internal.Menu{}
	menus = append(menus, ubereats.Menus()...)

	writeFile(menus)
}
