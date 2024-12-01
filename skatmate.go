package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type Abrechnungsform string

const (
	Bierlachs     Abrechnungsform = "bierlachs"
	LeipzigerSkat Abrechnungsform = "leipzigerSkat"
	Unknown       Abrechnungsform = "Unkown"
)

type DocSpiel struct {
	Punkte        int
	SpielerPunkte []string
}

type DocSkatrunde struct {
	Date            string
	Abrechnungsform Abrechnungsform
	Spieler         []string
	Spielverlauf    []DocSpiel
}

type Spiel struct {
	ReizGewinner int
	Gewonnen     bool
	Punkte       int
}

type Skatrunde struct {
	Date         string
	Spieler      []string
	Spielverlauf []Spiel
}

func (docSpiel DocSpiel) ToSpiel(abrechnungsform Abrechnungsform) Spiel {
	var spiel Spiel
	spiel.Punkte = docSpiel.Punkte
	var countEmptyInput int
	if abrechnungsform == Bierlachs {
		//Read all blanks
		for _, value := range docSpiel.SpielerPunkte {
			if value == "" {
				countEmptyInput++
			}
		}
		//decide if won or not
		if countEmptyInput == 1 {
			spiel.Gewonnen = true
		} else if countEmptyInput == 2 {
			spiel.Gewonnen = false
		}
		//check who is Reizgewinner
		for i, value := range docSpiel.SpielerPunkte {
			if spiel.Gewonnen {
				if value == "" {
					spiel.ReizGewinner = i
				}
			} else if _, err := strconv.Atoi(value); err == nil {
				spiel.ReizGewinner = i
			}
		}
	} else if abrechnungsform == LeipzigerSkat {
		for i, value := range docSpiel.SpielerPunkte {
			//check for number -> thats the player
			if _, err := strconv.Atoi(value); err == nil {
				spiel.ReizGewinner = i
			}
			//check for Punkte pos -> won; neg -> lost
			if docSpiel.Punkte > 0 {
				spiel.Gewonnen = true
			} else {
				spiel.Gewonnen = false
			}
		}
	}
	return spiel
}

func (docSkatrunde DocSkatrunde) ToSkatrunde() Skatrunde {
	var skatrunde Skatrunde
	skatrunde.Date = docSkatrunde.Date
	skatrunde.Spieler = docSkatrunde.Spieler
	for _, spiel := range docSkatrunde.Spielverlauf {
		skatrunde.Spielverlauf = append(skatrunde.Spielverlauf, spiel.ToSpiel(docSkatrunde.Abrechnungsform))
	}
	return skatrunde
}

func (spiel Spiel) ToDocSpiel(abrechnungsform Abrechnungsform, isVierSpieler bool, spielNr int, spielerPunkteZuvor []int) DocSpiel {
	var docSpiel DocSpiel
	var spielerPunkte []string

	if abrechnungsform == Bierlachs {
		docSpiel.Punkte = spiel.Punkte
		//set waiting person
		if isVierSpieler {
			spielerPunkte = make([]string, 4)
			spielerPunkte[spielNr%4] = "*"
		} else {
			spielerPunkte = make([]string, 3)
		}
		if spiel.Gewonnen {
			for i, value := range spielerPunkte {
				if i != spiel.ReizGewinner && value == "" {
					spielerPunkte[i] = strconv.Itoa(spiel.Punkte + spielerPunkteZuvor[i])
				}
			}
		} else {
			spielerPunkte[spiel.ReizGewinner] = strconv.Itoa(spiel.Punkte + spielerPunkteZuvor[spiel.ReizGewinner])
		}

	} else if abrechnungsform == LeipzigerSkat {
		if spiel.Gewonnen {
			docSpiel.Punkte = spiel.Punkte
		} else {
			docSpiel.Punkte = spiel.Punkte * -1
		}

		//Set waiting person
		if isVierSpieler {
			spielerPunkte = make([]string, 4)
			spielerPunkte[spielNr%4] = "*"
		} else {
			spielerPunkte = make([]string, 3)
		}
		spielerPunkte[spiel.ReizGewinner] = strconv.Itoa(docSpiel.Punkte + spielerPunkteZuvor[spiel.ReizGewinner])
	}
	docSpiel.SpielerPunkte = spielerPunkte
	return docSpiel
}

func (skatrunde Skatrunde) ToDocSkatrunde(abrechnungsform Abrechnungsform) DocSkatrunde {
	var docSkatrunde DocSkatrunde
	docSkatrunde.Date = skatrunde.Date
	docSkatrunde.Spieler = skatrunde.Spieler
	docSkatrunde.Abrechnungsform = abrechnungsform
	var isVierSpieler bool
	if len(skatrunde.Spieler) == 4 {
		isVierSpieler = true
	} else {
		isVierSpieler = false
	}

	spielerPunkteZuvor := make([]int, 4)
	for i, spiel := range skatrunde.Spielverlauf {
		var docSpiel = spiel.ToDocSpiel(abrechnungsform, isVierSpieler, i, spielerPunkteZuvor)
		for i, spielerWert := range docSpiel.SpielerPunkte {
			intWert, _ := strconv.Atoi(spielerWert)
			if intWert != 0 {
				spielerPunkteZuvor[i] = intWert
			}
		}
		docSkatrunde.Spielverlauf = append(docSkatrunde.Spielverlauf, docSpiel)
	}

	return docSkatrunde
}

func ParseSpiel(input []string) (*DocSpiel, error) {
	var spiel DocSpiel
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

func ParseSkatCsvFile(input string) DocSkatrunde {
	//Import and parse DocSkatrunden Struktur
	var docSkatrunde DocSkatrunde
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
	filePath := "game_data.csv"
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
	var bierlachsSkatrunde = skatrunde.ToDocSkatrunde(Bierlachs)
	log.Printf("[Transformed] Skatrunde -> DocSkatrunde mit Bierlachs : %+v\n", bierlachsSkatrunde)

	//transfer Skatrunde -> DocSkatrunde LeipzigerSkat
	var leipzigSkatrunde = skatrunde.ToDocSkatrunde(LeipzigerSkat)
	log.Printf("[Transformed] Skatrunde -> DocSkatrunde mit LeipzigerSkat : %+v\n", leipzigSkatrunde)

}
