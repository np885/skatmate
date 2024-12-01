package main

import(
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Skatrunde struct {

}

func main() {
	filePath := "game_data.csv"
	content, err := ioutil.ReadFile(filePath)
	
	if err != nil {
		log.Fatalf("Faild to read file: %w", err)
	}

	strContent := string(content)
	fmt.Println(strContent)

	var players []string
	var abrechnungsForm string
	var date string

	for i, line := range strings.Split(strContent, "\n") {
		if i == 0 {
			//Get MetaData information
			metaData := strings.Split(line, ";")
			if len(metaData) == 2 {
				abrechnungsForm = strings.Split(metaData[0], "=")[1]
				date = strings.Split(metaData[1], "=")[1]
				fmt.Printf("[Parsed] AbrechnungsForm is %s and date is %s\n", abrechnungsForm, date)
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
				players = header[:3]
			} else if len(header) == 5 {
				//3 Player round
				players = header[:4]
			} else {
				log.Fatalf("[line %d] Error parsing HeaderData\n", i+1)
			}
			fmt.Printf("[Parsed] Players: %s\n", players)
		} else if i > 1 {
			//Get Gameround information
			gameRound := strings.Split(line, ";")
			if len(gameRound) == 4 || len(gameRound) == 5 {	
				for i,player := range players {
					fmt.Printf("[Parsed] Player: %s has %s Points in Gameround: %d\n",player, gameRound[i], i+1)
				}
			} else {
				log.Printf("[line %d] Error parsing gameround. Content: %s\n", i+1, line)
			}
		}
		
	}
}
