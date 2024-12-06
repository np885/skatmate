package skat

import (
	"fmt"
	"sort"
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
	Punkte        []int
	SpielerPunkte []string
}

type Platz struct {
	Nr        int
	Endpunkte int
}

func (p Platz) ToString() string {
	return fmt.Sprintf("%d,%d", p.Nr, p.Endpunkte)
}

type DocSkatrunde struct {
	Date            string
	Abrechnungsform Abrechnungsform
	Spieler         []string
	Spielverlauf    []DocSpiel
	Platzierung     []Platz
}

type Spiel struct {
	ReizGewinner int
	Gewonnen     bool
	Punkte       []int
}

type Skatrunde struct {
	Date         string
	Spieler      []string
	Spielverlauf []Spiel
}

// Kinda ugly there might be better ways to do it
func calculatePlatzierung(spielerPunkte []int, abrechnungsform Abrechnungsform) []Platz {
	spielerAnzahl := len(spielerPunkte)
	var platzierung = make([]Platz, spielerAnzahl)
	//copy and sort the copy to check who's won
	sortedSpielerPunkte := make([]int, spielerAnzahl)
	copy(sortedSpielerPunkte, spielerPunkte)
	sort.Ints(sortedSpielerPunkte)
	//set offset kinda ugly
	var offset []int
	if abrechnungsform == LeipzigerSkat {
		if spielerAnzahl == 4 {
			offset = []int{4, 2, 0, -2}
		} else {
			offset = []int{3, 1, -1}
		}
	} else if abrechnungsform == Bierlachs {
		offset = []int{1, 1, 1, 1}
	}

	for i, punkte := range spielerPunkte {
		for d, sortedPunkte := range sortedSpielerPunkte {
			if punkte == sortedPunkte {
				platzierung[i] = Platz{d + offset[d], punkte}
			}
		}
	}
	return platzierung
}

func (docSpiel DocSpiel) ToSpiel(abrechnungsform Abrechnungsform) Spiel {
	var spiel Spiel
	var countEmptyInput int

	ramsch := len(docSpiel.Punkte) > 1

	if abrechnungsform == Bierlachs {
		spiel.Punkte = docSpiel.Punkte
		//Read all blanks
		for _, value := range docSpiel.SpielerPunkte {
			if value == "" {
				countEmptyInput++
			}
		}
		//First check if its Ramsch
		if ramsch {
			spiel.Gewonnen = false
		} else {
			//decide if won or not
			if countEmptyInput == 1 {
				spiel.Gewonnen = true
			} else if countEmptyInput == 2 {
				spiel.Gewonnen = false
			}
		}

		//First check ramsch
		if ramsch {
			spiel.ReizGewinner = -1
		} else {
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
		}
	} else if abrechnungsform == LeipzigerSkat {
		for i, value := range docSpiel.SpielerPunkte {

			if ramsch {
				spiel.ReizGewinner = -1
				spiel.Gewonnen = false
				var spielPunkte []int
				for _, punkte := range docSpiel.Punkte {
					spielPunkte = append(spielPunkte, punkte*-1)
				}
				spiel.Punkte = spielPunkte
			} else {
				//check for number -> thats the player
				if _, err := strconv.Atoi(value); err == nil {
					spiel.ReizGewinner = i
					//check for Punkte pos -> won; neg -> lost
					if docSpiel.Punkte[0] > 0 {
						spiel.Gewonnen = true
						spiel.Punkte = docSpiel.Punkte
					} else {
						spiel.Gewonnen = false
						spiel.Punkte = []int{docSpiel.Punkte[0] * -1}
					}
				}
			}
		}
	}
	return spiel
}

func (docSpiel DocSpiel) ToString() string {
	var str string
	for _, spielerPunkte := range docSpiel.SpielerPunkte {
		str += fmt.Sprintf("%v;", spielerPunkte)
	}
	str += fmt.Sprintf("%v;", docSpiel.Punkte)
	return str
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
	ramsch := spiel.ReizGewinner == -1

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
					spielerPunkte[i] = strconv.Itoa(spiel.Punkte[0] + spielerPunkteZuvor[i])
				}
			}
		} else {
			//check if ramsch
			if ramsch {
				d := 0
				for i, _ := range spielerPunkte {
					if spielerPunkte[i] != "*" {
						spielerPunkte[i] = strconv.Itoa(spiel.Punkte[d] + spielerPunkteZuvor[i])
						d++
					}
				}
			} else {
				spielerPunkte[spiel.ReizGewinner] = strconv.Itoa(spiel.Punkte[0] + spielerPunkteZuvor[spiel.ReizGewinner])
			}
		}
	} else if abrechnungsform == LeipzigerSkat {
		if spiel.Gewonnen {
			docSpiel.Punkte = spiel.Punkte
		} else {
			var spielPunkte []int
			for _, value := range spiel.Punkte {
				spielPunkte = append(spielPunkte, value*-1)
			}
			docSpiel.Punkte = spielPunkte
		}

		//Set waiting person
		if isVierSpieler {
			spielerPunkte = make([]string, 4)
			spielerPunkte[spielNr%4] = "*"
		} else {
			spielerPunkte = make([]string, 3)
		}
		if ramsch {
			d := 0
			for i, _ := range spielerPunkte {
				if spielerPunkte[i] != "*" {
					spielerPunkte[i] = strconv.Itoa(spiel.Punkte[d] + spielerPunkteZuvor[i])
					d++
				}
			}
		} else {
			spielerPunkte[spiel.ReizGewinner] = strconv.Itoa(docSpiel.Punkte[0] + spielerPunkteZuvor[spiel.ReizGewinner])
		}
	}
	docSpiel.SpielerPunkte = spielerPunkte
	return docSpiel
}

func (spiel Spiel) ToString() string {
	var str string
	if spiel.ReizGewinner >= 0 {
		str = fmt.Sprintf("Spieler %d ist ReizGewinner;", spiel.ReizGewinner)
	} else {
		str = fmt.Sprintf("Kein Reizgewinner > RAMSCH;")
	}
	if spiel.Gewonnen {
		str += fmt.Sprintf("Spiel wurde gewonnen;")
	} else {
		str += fmt.Sprintf("Spiel wurde verloren;")
	}
	strArrOfPunkte := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(spiel.Punkte)), ","), "")
	str += fmt.Sprintf("Punkte %v ", strArrOfPunkte)
	return str
}

func (skatrunde Skatrunde) ToDocSkatrunde(abrechnungsform Abrechnungsform) DocSkatrunde {
	var docSkatrunde DocSkatrunde
	docSkatrunde.Date = skatrunde.Date
	docSkatrunde.Spieler = skatrunde.Spieler
	docSkatrunde.Abrechnungsform = abrechnungsform
	var isVierSpieler bool
	var spielerPunkteZuvor []int
	if len(skatrunde.Spieler) == 4 {
		isVierSpieler = true
		spielerPunkteZuvor = make([]int, 4)
	} else {
		isVierSpieler = false
		spielerPunkteZuvor = make([]int, 3)

	}

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
	//Platzierung berechnen
	docSkatrunde.Platzierung = calculatePlatzierung(spielerPunkteZuvor, abrechnungsform)
	return docSkatrunde
}

func (docSkatrunde DocSkatrunde) ToString() string {
	str := fmt.Sprintf("abrechnungsform=%v;date=%s\n", docSkatrunde.Abrechnungsform, docSkatrunde.Date)
	for _, spieler := range docSkatrunde.Spieler {
		str += fmt.Sprintf("%v;", spieler)
	}
	str += "Spiel;\n"
	for _, spiel := range docSkatrunde.Spielverlauf {
		str += fmt.Sprintf("%s\n", spiel.ToString())
	}
	str += "==========\n"
	for _, platz := range docSkatrunde.Platzierung {
		str += fmt.Sprintf("%s;", platz.ToString())
	}
	return str
}

func (skatrunde Skatrunde) ToString() string {
	str := fmt.Sprintf("Skatrunde vom %s\n", skatrunde.Date)
	str += fmt.Sprintf("Spieler %s\n", skatrunde.Spieler)
	for _, spiel := range skatrunde.Spielverlauf {
		str += fmt.Sprintf("%v\n", spiel.ToString())
	}
	return str
}
