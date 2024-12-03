package main

import (
	"errors"
	"github.com/np885/skatmate/skat"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseSpiel(input []string) (*skat.DocSpiel, error) {
	var spiel skat.DocSpiel
	if len(input) == 4 {
		spiel.SpielerPunkte = input[:3]
		punkteInt, err := strconv.Atoi(input[3])
		if err != nil {
			return nil, errors.New("Punkte in DocSpiel are not a number")
		}
		spiel.Punkte = punkteInt
	} else if len(input) == 5 {
		spiel.SpielerPunkte = input[:4]
		punkteInt, err := strconv.Atoi(input[4])
		if err != nil {
			return nil, errors.New("Punkte in DocSpiel are not a number")
		}
		spiel.Punkte = punkteInt
	} else {
		return nil, errors.New("Invalid DocSpiel in Dataset")
	}
	return &spiel, nil
}

func ParseAbrechnungsform(input string) (skat.Abrechnungsform, error) {
	switch strings.ToLower(input) {
	case "bierlachs":
		return skat.Bierlachs, nil
	case "leipzigerskat":
		return skat.LeipzigerSkat, nil
	default:
		return skat.Unknown, errors.New("Invalid Abrechnungsform as input")
	}
}

func ParseSkatCsvFile(input string) skat.DocSkatrunde {
	//Import and parse DocSkatrunden Struktur
	var docSkatrunde skat.DocSkatrunde
	for i, line := range strings.Split(input, "\n") {
		if i == 0 {
			//Get MetaData information
			metaData := strings.Split(line, ";")
			if len(metaData) == 2 {
				abrechnungsform, err := ParseAbrechnungsform(strings.Split(metaData[0], "=")[1])
				if err == nil {
					docSkatrunde.Abrechnungsform = abrechnungsform
				} else {
					log.Fatalf("[line %d] Error parsing Abrechnungsform.\n", i)
				}
				docSkatrunde.Date = strings.Split(metaData[1], "=")[1]
			} else {
				//error
				log.Fatalf("[line %d] Error parsing MetaData\n", i+1)
			}

		} else if i == 1 {
			//Get player information
			header := strings.Split(line, ";")
			game := header[len(header)-1]
			if !strings.EqualFold("Spiel", game) {
				log.Fatalf("[line %d] Error parsing `DocSpiel`in Header", i+1)
			}

			if len(header) == 4 {
				//4 Player round
				docSkatrunde.Spieler = header[:3]
			} else if len(header) == 5 {
				//3 Player round
				docSkatrunde.Spieler = header[:4]
			} else {
				log.Fatalf("[line %d] Error parsing HeaderData\n", i+1)
			}
		} else if i > 1 {
			//Get Gameround information
			spielRunde := strings.Split(line, ";")
			spiel, err := ParseSpiel(spielRunde)
			if err == nil {
				docSkatrunde.Spielverlauf = append(docSkatrunde.Spielverlauf, *spiel)
			} else {
				log.Printf("[line %d] Error parsing DocSpiel: %s\n", i+1, err)
			}
		}
	}
	return docSkatrunde
}

func main() {
	filePath := "../game_data.csv"
	content, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Faild to read file: %w", err)
	}
	strContent := string(content)

	docSkatrunde := ParseSkatCsvFile(strContent)

	//finished parsing
	log.Printf("[Paresed] DocSkatrunde: %+v\n", docSkatrunde)

	//transfer DocSkatrunde -> Skatrunde
	var skatrunde = docSkatrunde.ToSkatrunde()
	log.Printf("[Transformed] DocSkatrunde -> Skatrunde: %+v\n", skatrunde)

	//transfer Skatrunde -> DocSkatrunde Bierlachs
	var bierlachsSkatrunde = skatrunde.ToDocSkatrunde(skat.Bierlachs)
	log.Printf("[Transformed] Skatrunde -> DocSkatrunde mit Bierlachs : %+v\n", bierlachsSkatrunde)

	//transfer Skatrunde -> DocSkatrunde LeipzigerSkat
	var leipzigSkatrunde = skatrunde.ToDocSkatrunde(skat.LeipzigerSkat)
	log.Printf("[Transformed] Skatrunde -> DocSkatrunde mit LeipzigerSkat : %+v\n", leipzigSkatrunde)

}
