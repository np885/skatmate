package skat

import "strconv"

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

	var countEmptyInput int
	if abrechnungsform == Bierlachs {
		spiel.Punkte = docSpiel.Punkte
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
				spiel.Punkte = docSpiel.Punkte
			} else {
				spiel.Gewonnen = false
				spiel.Punkte = docSpiel.Punkte * -1
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
