package main

import (
	"github.com/np885/skatmate/skat"
	"log"
	"os"
	"path/filepath"
)

func getCSVFiles(dir string) ([]string, error) {
	var csvFiles []string

	// Walk through the directory
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file has a .csv extension
		if !info.IsDir() && filepath.Ext(info.Name()) == ".csv" {
			csvFiles = append(csvFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return csvFiles, nil
}

func main() {
	//filePath := "../game_data.csv"
	//filePath := "../data/20241204_bierlachs.csv"
	//filePath := "../data/20241009_leipzigerskat.csv"
	csvFiles, err := getCSVFiles("../data")
	if err != nil {
		log.Fatalf("Error while reading directories for csv-Files, %v", err)
	}
	var skatrunden []skat.Skatrunde

	for _, filepath := range csvFiles {
		content, err := os.ReadFile(filepath)
		if err != nil {
			log.Fatalf("Faild to read file: %w", err)
		}
		strContent := string(content)

		docSkatrunde := ParseSkatCsvFile(strContent)

		skatrunden = append(skatrunden, docSkatrunde.ToSkatrunde())
	}

	log.Printf("Skatrunden geladen, Anzhal: %d\n", len(skatrunden))
	for _, skatrunde := range skatrunden {
		log.Printf("%v\n", skatrunde.ToString())
	}
}
