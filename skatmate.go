package main

import(
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"errors"
	"strconv"
)

type Abrechnungsform string

const (
	Bierlachs Abrechnungsform = "bierlachs"
	LeipzigerSkat Abrechnungsform = "leipzigerSkat"
	Unknown Abrechnungsform = "Unkown"
)

type Spiel struct {
	Punkte int
	SpielerPunkte []string
}

type Skatrunde struct {
	Date string
	Abrechnungsform Abrechnungsform
	Spieler []string
	Spielverlauf []Spiel
}

func ParseSpiel(input []string) (*Spiel, error) {
	var spiel Spiel
	if len(input) == 4 {
		spiel.SpielerPunkte = input[:3]
		punkteInt, err := strconv.Atoi(input[3])
		if err != nil {
			return nil, errors.New("Punkte in Spiel are not a number")
		}
		spiel.Punkte = punkteInt
	} else if len(input) == 5 {
		spiel.SpielerPunkte = input[:4]
		punkteInt, err := strconv.Atoi(input[4])
		if err != nil {
			return nil, errors.New("Punkte in Spiel are not a number")
		}
		spiel.Punkte = punkteInt
	} else {
		return nil, errors.New("Invalid Spiel in Dataset")
	}
	return &spiel, nil
}

func ParseAbrechnungsform(input string) (Abrechnungsform, error) {
	switch strings.ToLower(input) {
		case "bierlachs":
			return Bierlachs, nil
		case "leipzigerskat":
			return LeipzigerSkat, nil
		default:
			return Unknown, errors.New("Invalid Abrechnungsform as input")
	}
}

func main() {
	filePath := "game_data.csv"
	content, err := ioutil.ReadFile(filePath)
	
	if err != nil {
		log.Fatalf("Faild to read file: %w", err)
	}

	strContent := string(content)
	fmt.Println(strContent)

	var skatrunde Skatrunde

	for i, line := range strings.Split(strContent, "\n") {
		if i == 0 {
			//Get MetaData information
			metaData := strings.Split(line, ";")
			if len(metaData) == 2 {
				abrechnungsform, err := ParseAbrechnungsform(strings.Split(metaData[0], "=")[1])
				if err == nil {
					skatrunde.Abrechnungsform = abrechnungsform
				} else {
					log.Fatalf("[line %d] Error parsing Abrechnungsform.\n", i)
				}
				skatrunde.Date = strings.Split(metaData[1], "=")[1]
				fmt.Printf("[Parsed] AbrechnungsForm is %s and date is %s\n", skatrunde.Abrechnungsform, skatrunde.Date)
			} else {
				//error
				log.Fatalf("[line %d] Error parsing MetaData\n", i+1)
			}

		} else if i == 1 {
			//Get player information
			header := strings.Split(line, ";")
			game := header[len(header) - 1]
			if !strings.EqualFold("Spiel", game) {
				log.Fatalf("[line %d] Error parsing `Spiel`in Header", i+1)
			}
			if len(header) == 4 {
				//4 Player round
				skatrunde.Spieler = header[:3]
			} else if len(header) == 5 {
				//3 Player round
				skatrunde.Spieler = header[:4]
			} else {
				log.Fatalf("[line %d] Error parsing HeaderData\n", i+1)
			}
			fmt.Printf("[Parsed] Spieler: %s\n", skatrunde.Spieler)
		} else if i > 1 {
			//Get Gameround information
			spielRunde := strings.Split(line, ";")
			spiel, err := ParseSpiel(spielRunde)
			if err == nil {
				skatrunde.Spielverlauf = append(skatrunde.Spielverlauf, *spiel)
			} else {
				log.Printf("[line %d] Error parsing Spiel: %s\n", i+1, err)
			}
		}
	}

	//finished parsing
	log.Printf("[Paresed] Skatrunde: %+v\n", skatrunde)
}
