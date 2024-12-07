package main

import (
	"github.com/np885/skatmate/skat"
	"log"
	"os"
	"path/filepath"
	"time"
)

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func getCSVFiles(args []string) ([]string, error) {
	var csvFiles []string

	//check if its file or dir
	for _, arg := range args {
		isDir, err := isDirectory(arg)
		if err != nil {
			return nil, err
		}
		if isDir {
			csvFilesFromDir, err := getCSVFilesFromDir(arg)
			if err != nil {
				return nil, err
			}
			csvFiles = append(csvFiles, csvFilesFromDir...)
		} else {
			if filepath.Ext(arg) == ".csv" {
				csvFiles = append(csvFiles, arg)
			}
		}
	}
	return csvFiles, nil
}

func getCSVFilesFromDir(dir string) ([]string, error) {
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

func getDocSkatrundeFromFile(filepath string) skat.DocSkatrunde {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Faild to read file: %w", err)
	}
	strContent := string(content)

	docSkatrunde := ParseSkatCsvFile(strContent)
	return docSkatrunde
}

func main() {
	//An programmargument is either a file or a directory
	args := os.Args
	log.Printf("Programm executed with args: %v", args[1:])
	csvFiles, err := getCSVFiles(args)
	if err != nil {
		log.Fatalf("Error while reading directories for csv-Files, %v", err)
	}
	var skatrunden []skat.Skatrunde

	//Paralleles Einlesen der Files
	startImportCSVFiles := time.Now()
	c := make(chan skat.DocSkatrunde)
	for _, csvFilepath := range csvFiles {
		go func() { c <- getDocSkatrundeFromFile(csvFilepath) }()
	}

	for range csvFiles {
		docSkatrunde := <-c
		skatrunden = append(skatrunden, docSkatrunde.ToSkatrunde())
	}
	elapsedImportCSVFiles := time.Since(startImportCSVFiles)
	log.Printf("Parallel: Import CSV-Files took %s", elapsedImportCSVFiles)

	log.Printf("Skatrunden geladen, Anzhal: %d\n", len(skatrunden))
	for _, skatrunde := range skatrunden {
		log.Printf("%v\n", skatrunde.ToString())
	}
}
