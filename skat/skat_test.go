package skat

import (
	"testing"
)

func TestDocSpiel_ToSpielBierlachsLost(t *testing.T) {
	//Given
	docSpiel := DocSpiel{[]int{40}, []string{"40", "", ""}}

	//When
	spiel := docSpiel.ToSpiel(Bierlachs)

	//Then
	if spiel.Gewonnen {
		t.Errorf("spiel should be lost but its won")
	}
	if spiel.Punkte[0] != 40 {
		t.Errorf("spiel points are not 20")
	}
	if spiel.ReizGewinner != 0 {
		t.Errorf("reizGewinner should be the first player")
	}
}

func TestDocSpiel_ToSpielBierlachsWon(t *testing.T) {
	//Given
	docSpiel := DocSpiel{[]int{20}, []string{"", "20", "20"}}

	//When
	spiel := docSpiel.ToSpiel(Bierlachs)

	//Then
	if !spiel.Gewonnen {
		t.Errorf("spiel should be won but its lost")
	}
	if spiel.Punkte[0] != 20 {
		t.Errorf("spiel points are not 20")
	}
	if spiel.ReizGewinner != 0 {
		t.Errorf("reizGewinner should be the first player")
	}
}

func TestDocSpiel_ToSpielRamsch4SpielerBierachs(t *testing.T) {
	//Given
	docSpiel := DocSpiel{[]int{40, 60, 20}, []string{"40", "60", "20", "*"}}

	//When
	spiel := docSpiel.ToSpiel(Bierlachs)

	//Then
	if spiel.Gewonnen {
		t.Errorf("spiel should be lost but its won")
	}
	if len(spiel.Punkte) != 3 {
		t.Errorf("spiel Punkte should be 3 integer Fields but was %d", len(spiel.Punkte))
	}
	if spiel.Punkte[0] != 40 {
		t.Errorf("spiel points are not 40, but was %d", spiel.Punkte[0])
	}
	if spiel.Punkte[1] != 60 {
		t.Errorf("spiel points are not 60, but was %d", spiel.Punkte[1])
	}
	if spiel.Punkte[2] != 20 {
		t.Errorf("spiel points are not 20, but was %d", spiel.Punkte[2])
	}
	if spiel.ReizGewinner != -1 {
		t.Errorf("Reizgewinner should be -1 but was %d", spiel.ReizGewinner)
	}
}

func TestDocSpiel_ToSpielRamsch3SpielerBierachs(t *testing.T) {
	//Given
	docSpiel := DocSpiel{[]int{40, 60, 20}, []string{"40", "60", "20"}}

	//When
	spiel := docSpiel.ToSpiel(Bierlachs)

	//Then
	if spiel.Gewonnen {
		t.Errorf("spiel should be lost but its won")
	}
	if len(spiel.Punkte) != 3 {
		t.Errorf("spiel Punkte should be 3 integer Fields but was %d", len(spiel.Punkte))
	}
	if spiel.Punkte[0] != 40 {
		t.Errorf("spiel points are not 40, but was %d", spiel.Punkte[0])
	}
	if spiel.Punkte[1] != 60 {
		t.Errorf("spiel points are not 60, but was %d", spiel.Punkte[1])
	}
	if spiel.Punkte[2] != 20 {
		t.Errorf("spiel points are not 20, but was %d", spiel.Punkte[2])
	}
	if spiel.ReizGewinner != -1 {
		t.Errorf("Reizgewinner should be -1 but was %d", spiel.ReizGewinner)
	}
}

func TestDocSpiel_ToSpielLeipzigerSkatWon(t *testing.T) {
	//Given
	docSpiel := DocSpiel{[]int{20}, []string{"", "20", ""}}

	//When
	spiel := docSpiel.ToSpiel(LeipzigerSkat)

	//Then
	if !spiel.Gewonnen {
		t.Errorf("spiel should be won but its lost")
	}
	if spiel.Punkte[0] != 20 {
		t.Errorf("spiel points are not 20")
	}
	if spiel.ReizGewinner != 1 {
		t.Errorf("reizGewinner should be the secound player")
	}
}

func TestDocSpiel_ToSpielLeipzigerSkatLost(t *testing.T) {
	//Given
	docSpiel := DocSpiel{[]int{-20}, []string{"", "-20", ""}}

	//When
	spiel := docSpiel.ToSpiel(LeipzigerSkat)

	//Then
	if spiel.Gewonnen {
		t.Errorf("spiel should be lost but its won")
	}
	if spiel.Punkte[0] != 20 {
		t.Errorf("spiel points are not 20, but was %d", spiel.Punkte[0])
	}
	if spiel.ReizGewinner != 1 {
		t.Errorf("reizGewinner should be the secound player, but was %d", spiel.ReizGewinner)
	}
}

func TestDocSpiel_ToSpielRamsch4SpielerLeipzigerSkat(t *testing.T) {
	//Given
	docSpiel := DocSpiel{[]int{-40, -60, -20}, []string{"-40", "-60", "-20", "*"}}

	//When
	spiel := docSpiel.ToSpiel(LeipzigerSkat)

	//Then
	if spiel.Gewonnen {
		t.Errorf("spiel should be lost but its won")
	}
	if len(spiel.Punkte) != 3 {
		t.Errorf("spiel Punkte should be 3 integer Fields but was %d", len(spiel.Punkte))
	}
	if spiel.Punkte[0] != 40 {
		t.Errorf("spiel points are not 40, but was %d", spiel.Punkte[0])
	}
	if spiel.Punkte[1] != 60 {
		t.Errorf("spiel points are not 60, but was %d", spiel.Punkte[1])
	}
	if spiel.Punkte[2] != 20 {
		t.Errorf("spiel points are not 20, but was %d", spiel.Punkte[2])
	}
	if spiel.ReizGewinner != -1 {
		t.Errorf("Reizgewinner should be -1 but was %d", spiel.ReizGewinner)
	}
}

func TestDocSpiel_ToSpielRamsch3SpielerLeipzigerSkat(t *testing.T) {
	//Given
	docSpiel := DocSpiel{[]int{-40, -60, -20}, []string{"40", "60", "20"}}

	//When
	spiel := docSpiel.ToSpiel(LeipzigerSkat)

	//Then
	if spiel.Gewonnen {
		t.Errorf("spiel should be lost but its won")
	}
	if len(spiel.Punkte) != 3 {
		t.Errorf("spiel Punkte should be 3 integer Fields but was %d", len(spiel.Punkte))
	}
	if spiel.Punkte[0] != 40 {
		t.Errorf("spiel points are not 40, but was %d", spiel.Punkte[0])
	}
	if spiel.Punkte[1] != 60 {
		t.Errorf("spiel points are not 60, but was %d", spiel.Punkte[1])
	}
	if spiel.Punkte[2] != 20 {
		t.Errorf("spiel points are not 20, but was %d", spiel.Punkte[2])
	}
	if spiel.ReizGewinner != -1 {
		t.Errorf("Reizgewinner should be -1 but was %d", spiel.ReizGewinner)
	}
}

func TestSpiel_ToDocSpielBierlachs4Player(t *testing.T) {
	//Given
	spiel := Spiel{1, true, []int{22}}

	//When
	docSpiel := spiel.ToDocSpiel(Bierlachs, true, 0, []int{0, 0, 0, 0})

	//Then
	if docSpiel.Punkte[0] != 22 {
		t.Errorf("docSpiel points are not 22")
	}
	if len(docSpiel.SpielerPunkte) != 4 {
		t.Errorf("docSpiel should have 4 players")
	}
	if docSpiel.SpielerPunkte[0] != "*" {
		t.Errorf("docSpiel player 1 should sit out")
	}
	if docSpiel.SpielerPunkte[1] != "" {
		t.Errorf("docSpiel player 2 won should have 0 points but has %s points", docSpiel.SpielerPunkte[1])
	}
	if docSpiel.SpielerPunkte[2] != "22" {
		t.Errorf("docSpiel player 3 lost should have 22 points but has %s points", docSpiel.SpielerPunkte[2])
	}
	if docSpiel.SpielerPunkte[3] != "22" {
		t.Errorf("docSpiel player 3 lost should have 22 points but has %s points", docSpiel.SpielerPunkte[3])
	}
}

func TestSpiel_ToDocSpielBierlachs3Player(t *testing.T) {
	//Given
	spiel := Spiel{2, false, []int{48}}

	//When
	docSpiel := spiel.ToDocSpiel(Bierlachs, false, 1, []int{0, 0, 40})

	//Then
	if docSpiel.Punkte[0] != 48 {
		t.Errorf("docSpiel points are not 48 but was %d", docSpiel.Punkte)
	}
	if len(docSpiel.SpielerPunkte) != 3 {
		t.Errorf("docSpiel should have 3 players but was %d", len(docSpiel.SpielerPunkte))
	}
	if docSpiel.SpielerPunkte[0] != "" {
		t.Errorf("docSpiel player 1 should has `` but has %s points", docSpiel.SpielerPunkte[0])
	}
	if docSpiel.SpielerPunkte[1] != "" {
		t.Errorf("docSpiel player 2 should has `` but has %s points", docSpiel.SpielerPunkte[1])
	}
	if docSpiel.SpielerPunkte[2] != "88" {
		t.Errorf("docSpiel player 3 lost should have 88 points but has %s points", docSpiel.SpielerPunkte[2])
	}
}

func TestSpiel_ToDocSpielLeipzigerSkat4Player(t *testing.T) {
	//Given
	spiel := Spiel{1, true, []int{22}}

	//When
	docSpiel := spiel.ToDocSpiel(LeipzigerSkat, true, 0, []int{0, 0, 0, 0})

	//Then
	if docSpiel.Punkte[0] != 22 {
		t.Errorf("docSpiel points are not 22")
	}
	if len(docSpiel.SpielerPunkte) != 4 {
		t.Errorf("docSpiel should have 4 players")
	}
	if docSpiel.SpielerPunkte[0] != "*" {
		t.Errorf("docSpiel player 1 should sit out")
	}
	if docSpiel.SpielerPunkte[1] != "22" {
		t.Errorf("docSpiel player 2 won should have 22 points but has %s points", docSpiel.SpielerPunkte[1])
	}
	if docSpiel.SpielerPunkte[2] != "" {
		t.Errorf("docSpiel player 3 lost `` points but has %s points", docSpiel.SpielerPunkte[2])
	}
	if docSpiel.SpielerPunkte[3] != "" {
		t.Errorf("docSpiel player 3 lost `` points but has %s points", docSpiel.SpielerPunkte[3])
	}
}

func TestSpiel_ToDocSpielLeipzigerSkat3Player(t *testing.T) {
	//Given
	spiel := Spiel{2, false, []int{48}}

	//When
	docSpiel := spiel.ToDocSpiel(LeipzigerSkat, false, 1, []int{0, 0, 40})

	//Then
	if docSpiel.Punkte[0] != -48 {
		t.Errorf("docSpiel points are not -48 but was %d", docSpiel.Punkte)
	}
	if len(docSpiel.SpielerPunkte) != 3 {
		t.Errorf("docSpiel should have 3 players but was %d", len(docSpiel.SpielerPunkte))
	}
	if docSpiel.SpielerPunkte[0] != "" {
		t.Errorf("docSpiel player 1 should has `` but has %s points", docSpiel.SpielerPunkte[0])
	}
	if docSpiel.SpielerPunkte[1] != "" {
		t.Errorf("docSpiel player 2 should has `` but has %s points", docSpiel.SpielerPunkte[1])
	}
	if docSpiel.SpielerPunkte[2] != "-8" {
		t.Errorf("docSpiel player 3 lost should have -8 points but has %s points", docSpiel.SpielerPunkte[2])
	}
}

func TestDocSkatrunde_ToSkatrundeBierlachs(t *testing.T) {
	//Given
	date := "01.01.2020"
	spielerList := []string{"Thomas", "Niclas", "Dennis"}
	docSpiel1 := DocSpiel{[]int{40}, []string{"40", "", ""}}
	docSpiel2 := DocSpiel{[]int{48}, []string{"88", "48", ""}}
	spielverlauf := []DocSpiel{docSpiel1, docSpiel2}

	var docSkatrunde DocSkatrunde
	docSkatrunde.Date = date
	docSkatrunde.Abrechnungsform = Bierlachs
	docSkatrunde.Spieler = spielerList
	docSkatrunde.Spielverlauf = spielverlauf

	//When
	skatrunde := docSkatrunde.ToSkatrunde()

	//Then
	for i, _ := range spielerList {
		if skatrunde.Spieler[i] != spielerList[i] {
			t.Errorf("player are different wanted %s but got %s", spielerList[i], skatrunde.Spieler[i])
		}
	}
	if skatrunde.Date != date {
		t.Errorf("date is not the same wanted %s but got %s", date, skatrunde.Date)
	}
	if len(skatrunde.Spielverlauf) != 2 {
		t.Errorf("Spielverlauf wrong wanted 2 but got %d", len(skatrunde.Spielverlauf))
	}

	if skatrunde.Spielverlauf[0].Gewonnen {
		t.Errorf("Spiel 1 got won but should be lost")
	}
	if skatrunde.Spielverlauf[0].ReizGewinner != 0 {
		t.Errorf("ReizGewinner wrong wanted 0 but got %d", skatrunde.Spielverlauf[0].ReizGewinner)
	}
	if skatrunde.Spielverlauf[0].Punkte[0] != 40 {
		t.Errorf("Spiel Punkte wrong wanted 40 but got %d", skatrunde.Spielverlauf[0].Punkte)
	}

	if !skatrunde.Spielverlauf[1].Gewonnen {
		t.Errorf("Spiel 2 got lost but should be won")
	}
	if skatrunde.Spielverlauf[1].ReizGewinner != 2 {
		t.Errorf("ReizGewinner wrong wanted 2 but got %d", skatrunde.Spielverlauf[0].ReizGewinner)
	}
	if skatrunde.Spielverlauf[1].Punkte[0] != 48 {
		t.Errorf("Spiel Punkte wrong wanted 48 but got %d", skatrunde.Spielverlauf[0].Punkte)
	}
}

func TestDocSkatrunde_ToSkatrundeLeipzigerSkat(t *testing.T) {
	//Given
	date := "01.01.2020"
	spielerList := []string{"Thomas", "Niclas", "Dennis", "Stefan"}
	docSpiel1 := DocSpiel{[]int{20}, []string{"*", "", "20", ""}}
	docSpiel2 := DocSpiel{[]int{-96}, []string{"-96", "*", "", ""}}
	spielverlauf := []DocSpiel{docSpiel1, docSpiel2}

	var docSkatrunde DocSkatrunde
	docSkatrunde.Date = date
	docSkatrunde.Abrechnungsform = LeipzigerSkat
	docSkatrunde.Spieler = spielerList
	docSkatrunde.Spielverlauf = spielverlauf

	//When
	skatrunde := docSkatrunde.ToSkatrunde()

	//Then
	for i, _ := range spielerList {
		if skatrunde.Spieler[i] != spielerList[i] {
			t.Errorf("player are different wanted %s but got %s", spielerList[i], skatrunde.Spieler[i])
		}
	}
	if skatrunde.Date != date {
		t.Errorf("date is not the same wanted %s but got %s", date, skatrunde.Date)
	}
	if len(skatrunde.Spielverlauf) != 2 {
		t.Errorf("Spielverlauf wrong wanted 2 but got %d", len(skatrunde.Spielverlauf))
	}

	if !skatrunde.Spielverlauf[0].Gewonnen {
		t.Errorf("Spiel 1 got lost but should be won")
	}
	if skatrunde.Spielverlauf[0].ReizGewinner != 2 {
		t.Errorf("ReizGewinner wrong wanted 2 but got %d", skatrunde.Spielverlauf[0].ReizGewinner)
	}
	if skatrunde.Spielverlauf[0].Punkte[0] != 20 {
		t.Errorf("Spiel Punkte wrong wanted 40 but got %d", skatrunde.Spielverlauf[0].Punkte)
	}

	if skatrunde.Spielverlauf[1].Gewonnen {
		t.Errorf("Spiel 2 got won but should be lost")
	}
	if skatrunde.Spielverlauf[1].ReizGewinner != 0 {
		t.Errorf("ReizGewinner wrong wanted 2 but got %d", skatrunde.Spielverlauf[0].ReizGewinner)
	}
	if skatrunde.Spielverlauf[1].Punkte[0] != 96 {
		t.Errorf("Spiel Punkte wrong wanted 48 but got %d", skatrunde.Spielverlauf[0].Punkte)
	}
}

func TestSkatrunde_ToDocSkatrundeBierlachs(t *testing.T) {
	//Given
	date := "01.01.2020"
	spielerList := []string{"Thomas", "Niclas", "Dennis"}
	spiel1 := Spiel{2, false, []int{48}}
	spiel2 := Spiel{2, true, []int{22}}
	spielverlauf := []Spiel{spiel1, spiel2}
	skatrunde := Skatrunde{date, spielerList, spielverlauf}

	//When
	docSkatrunde := skatrunde.ToDocSkatrunde(Bierlachs)

	//Then

	for i, _ := range spielerList {
		if docSkatrunde.Spieler[i] != spielerList[i] {
			t.Errorf("player are different wanted %s but got %s", spielerList[i], docSkatrunde.Spieler[i])
		}
	}
	if docSkatrunde.Date != date {
		t.Errorf("date is not the same wanted %s but got %s", date, docSkatrunde.Date)
	}
	if len(docSkatrunde.Spielverlauf) != 2 {
		t.Errorf("Spielverlauf wrong wanted 2 but got %d", len(skatrunde.Spielverlauf))
	}

	if docSkatrunde.Spielverlauf[0].Punkte[0] != 48 {
		t.Errorf("Spiel 1 got won but should be lost")
	}
	if docSkatrunde.Spielverlauf[0].SpielerPunkte[0] != "" {
		t.Errorf("Spieler 1 should has `` but got %s", docSkatrunde.Spielverlauf[0].SpielerPunkte[0])
	}
	if docSkatrunde.Spielverlauf[0].SpielerPunkte[1] != "" {
		t.Errorf("Spieler 2 should has `` but got %s", docSkatrunde.Spielverlauf[0].SpielerPunkte[0])
	}
	if docSkatrunde.Spielverlauf[0].SpielerPunkte[2] != "48" {
		t.Errorf("Spieler 3 should has 48 points but got %s", docSkatrunde.Spielverlauf[0].SpielerPunkte[0])
	}

	if docSkatrunde.Spielverlauf[1].Punkte[0] != 22 {
		t.Errorf("Spiel 1 got won but should be lost")
	}
	if docSkatrunde.Spielverlauf[1].SpielerPunkte[0] != "22" {
		t.Errorf("Spieler 1 should has 22 points but got %s", docSkatrunde.Spielverlauf[0].SpielerPunkte[0])
	}
	if docSkatrunde.Spielverlauf[1].SpielerPunkte[1] != "22" {
		t.Errorf("Spieler 2 should has 22 points but got %s", docSkatrunde.Spielverlauf[0].SpielerPunkte[0])
	}
	if docSkatrunde.Spielverlauf[1].SpielerPunkte[2] != "" {
		t.Errorf("Spieler 3 should has `` but got %s", docSkatrunde.Spielverlauf[0].SpielerPunkte[0])
	}
}

func TestSkatrunde_ToDocSkatrundeLeipzigerSkat(t *testing.T) {
	//Given
	date := "01.01.2020"
	spielerList := []string{"Thomas", "Niclas", "Dennis", "Stefan"}
	spiel1 := Spiel{2, false, []int{48}}
	spiel2 := Spiel{2, true, []int{22}}
	spielverlauf := []Spiel{spiel1, spiel2}
	skatrunde := Skatrunde{date, spielerList, spielverlauf}

	//When
	docSkatrunde := skatrunde.ToDocSkatrunde(LeipzigerSkat)

	//Then

	for i, _ := range spielerList {
		if docSkatrunde.Spieler[i] != spielerList[i] {
			t.Errorf("player are different wanted %s but got %s", spielerList[i], docSkatrunde.Spieler[i])
		}
	}
	if docSkatrunde.Date != date {
		t.Errorf("date is not the same wanted %s but got %s", date, docSkatrunde.Date)
	}
	if len(docSkatrunde.Spielverlauf) != 2 {
		t.Errorf("Spielverlauf wrong wanted 2 but got %d", len(skatrunde.Spielverlauf))
	}

	if docSkatrunde.Spielverlauf[0].Punkte[0] != -48 {
		t.Errorf("Spiel 1 got won but should be lost")
	}
	if docSkatrunde.Spielverlauf[0].SpielerPunkte[0] != "*" {
		t.Errorf("Spieler 1 should has `` but got %s", docSkatrunde.Spielverlauf[0].SpielerPunkte[0])
	}
	if docSkatrunde.Spielverlauf[0].SpielerPunkte[1] != "" {
		t.Errorf("Spieler 2 should has `` but got %s", docSkatrunde.Spielverlauf[0].SpielerPunkte[0])
	}
	if docSkatrunde.Spielverlauf[0].SpielerPunkte[2] != "-48" {
		t.Errorf("Spieler 3 should has 48 points but got %s", docSkatrunde.Spielverlauf[0].SpielerPunkte[0])
	}
	if docSkatrunde.Spielverlauf[0].SpielerPunkte[3] != "" {
		t.Errorf("Spieler 3 should has 48 points but got %s", docSkatrunde.Spielverlauf[0].SpielerPunkte[0])
	}

	if docSkatrunde.Spielverlauf[1].Punkte[0] != 22 {
		t.Errorf("Spiel 1 got won but should be lost")
	}
	if docSkatrunde.Spielverlauf[1].SpielerPunkte[0] != "" {
		t.Errorf("Spieler 1 should has 22 points but got %s", docSkatrunde.Spielverlauf[0].SpielerPunkte[0])
	}
	if docSkatrunde.Spielverlauf[1].SpielerPunkte[1] != "*" {
		t.Errorf("Spieler 2 should has 22 points but got %s", docSkatrunde.Spielverlauf[0].SpielerPunkte[0])
	}
	if docSkatrunde.Spielverlauf[1].SpielerPunkte[2] != "-26" {
		t.Errorf("Spieler 3 should has `` but got %s", docSkatrunde.Spielverlauf[0].SpielerPunkte[0])
	}
	if docSkatrunde.Spielverlauf[1].SpielerPunkte[3] != "" {
		t.Errorf("Spieler 3 should has `` but got %s", docSkatrunde.Spielverlauf[0].SpielerPunkte[0])
	}
}

func TestCalculatePlatzierung4SpielerBierlachs(t *testing.T) {
	//Given
	spielerPunkte := []int{22, 33, 44, 11}

	//When
	platzierung := calculatePlatzierung(spielerPunkte, Bierlachs)

	//Then
	if platzierung[0].Nr != 2 || platzierung[0].Endpunkte != 22 {
		t.Errorf("Spieler 1 should be places 2 with 22 points but was %d with points %d", platzierung[0].Nr, platzierung[0].Endpunkte)
	}
	if platzierung[1].Nr != 3 || platzierung[1].Endpunkte != 33 {
		t.Errorf("Spieler 2 should be places 3 with 33 points but was %d with points %d", platzierung[1].Nr, platzierung[1].Endpunkte)
	}
	if platzierung[2].Nr != 4 || platzierung[2].Endpunkte != 44 {
		t.Errorf("Spieler 3 should be places 4 with 44 points but was %d with points %d", platzierung[2].Nr, platzierung[2].Endpunkte)
	}
	if platzierung[3].Nr != 1 || platzierung[3].Endpunkte != 11 {
		t.Errorf("Spieler 4 should be places 1 with 11 points but was %d with points %d", platzierung[3].Nr, platzierung[3].Endpunkte)
	}
}

func TestCalculatePlatzierung3SpielerBierlachs(t *testing.T) {
	//Given
	spielerPunkte := []int{22, 33, 11}

	//When
	platzierung := calculatePlatzierung(spielerPunkte, Bierlachs)

	//Then
	if platzierung[0].Nr != 2 || platzierung[0].Endpunkte != 22 {
		t.Errorf("Spieler 1 should be places 2 with 22 points but was %d with points %d", platzierung[0].Nr, platzierung[0].Endpunkte)
	}
	if platzierung[1].Nr != 3 || platzierung[1].Endpunkte != 33 {
		t.Errorf("Spieler 2 should be places 3 with 33 points but was %d with points %d", platzierung[1].Nr, platzierung[1].Endpunkte)
	}
	if platzierung[2].Nr != 1 || platzierung[2].Endpunkte != 11 {
		t.Errorf("Spieler 4 should be places 1 with 11 points but was %d with points %d", platzierung[2].Nr, platzierung[2].Endpunkte)
	}
}

func TestCalculatePlatzierung4SpielerLeipzigerSkat(t *testing.T) {
	//Given
	spielerPunkte := []int{22, 33, 44, 11}

	//When
	platzierung := calculatePlatzierung(spielerPunkte, LeipzigerSkat)

	//Then
	if platzierung[0].Nr != 3 || platzierung[0].Endpunkte != 22 {
		t.Errorf("Spieler 1 should be places 3 with 22 points but was %d with points %d", platzierung[0].Nr, platzierung[0].Endpunkte)
	}
	if platzierung[1].Nr != 2 || platzierung[1].Endpunkte != 33 {
		t.Errorf("Spieler 2 should be places 2 with 33 points but was %d with points %d", platzierung[1].Nr, platzierung[1].Endpunkte)
	}
	if platzierung[2].Nr != 1 || platzierung[2].Endpunkte != 44 {
		t.Errorf("Spieler 3 should be places 1 with 44 points but was %d with points %d", platzierung[2].Nr, platzierung[2].Endpunkte)
	}
	if platzierung[3].Nr != 4 || platzierung[3].Endpunkte != 11 {
		t.Errorf("Spieler 4 should be places 4 with 11 points but was %d with points %d", platzierung[3].Nr, platzierung[3].Endpunkte)
	}
}

func TestCalculatePlatzierung3SpielerLeipzigerSkat(t *testing.T) {
	//Given
	spielerPunkte := []int{22, 33, 11}

	//When
	platzierung := calculatePlatzierung(spielerPunkte, LeipzigerSkat)

	//Then
	if platzierung[0].Nr != 2 || platzierung[0].Endpunkte != 22 {
		t.Errorf("Spieler 1 should be places 2 with 22 points but was %d with points %d", platzierung[0].Nr, platzierung[0].Endpunkte)
	}
	if platzierung[1].Nr != 1 || platzierung[1].Endpunkte != 33 {
		t.Errorf("Spieler 2 should be places 1 with 33 points but was %d with points %d", platzierung[1].Nr, platzierung[1].Endpunkte)
	}
	if platzierung[2].Nr != 3 || platzierung[2].Endpunkte != 11 {
		t.Errorf("Spieler 3 should be places 3 with 11 points but was %d with points %d", platzierung[2].Nr, platzierung[2].Endpunkte)
	}
}
