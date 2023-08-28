package de.polkow.skatmate.service;

import de.polkow.skatmate.model.DocSkatrunde;
import de.polkow.skatmate.model.DocSpiel;
import de.polkow.skatmate.persistence.entity.Skatrunde;
import de.polkow.skatmate.persistence.entity.Spielrunde;
import org.apache.commons.lang3.math.NumberUtils;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class SkatrundenImportService {

    public static Skatrunde importDocSkatrunde(DocSkatrunde docSkatrunde) {
        Skatrunde skatrunde = new Skatrunde();
        //TODO: auto genarte from DB
        skatrunde.setId(docSkatrunde.getId());
        skatrunde.setTageszeit(docSkatrunde.getTageszeit().toLocalDateTime());
        skatrunde.setAbrechnungsForm(docSkatrunde.getAbrechnungsForm());

        Map<Integer, String> spielerMap = getPositionAndSpieler(docSkatrunde);
        skatrunde.setSpieler(spielerMap);
        List<Spielrunde> spielrunden = new ArrayList<>(docSkatrunde.getSpielverlauf().size());
        for (DocSpiel docSpiel : docSkatrunde.getSpielverlauf()) {
            Spielrunde spielrunde = new Spielrunde();
            spielrunde.setRunde(docSpiel.getNr());
            spielrunde.setPunkte(docSpiel.getPunkte());
            spielrunde.setSpielArt(docSpiel.getSpielArt());

            Map<String, String> spielerWerte = docSpiel.getSpielerWerteMap();

            long countBlanks = spielerWerte.values().stream().filter(""::equalsIgnoreCase).count();
            for(String spieler : spielerWerte.keySet()) {
                String wert = spielerWerte.get(spieler);
                if (countBlanks == 1 && "".equalsIgnoreCase(wert)) {
                    //Spiel gewonnen
                    spielrunde.setGewonnen(true);
                    spielrunde.setReizGewinner(spieler);
                    break;
                } else if  (countBlanks == 2 && NumberUtils.isParsable(wert)) {
                    //Spiel verloren
                    spielrunde.setGewonnen(false);
                    spielrunde.setReizGewinner(spieler);
                    break;

                }
            }

            spielrunden.add(spielrunde);
        }
        skatrunde.setSpielrunden(spielrunden);

        return skatrunde;
    }

    private static Map<Integer, String> getPositionAndSpieler(DocSkatrunde docSkatrunde) {
        Map<Integer, String> spielerMap = new HashMap<>();
        spielerMap.put(0, docSkatrunde.getSpielerReihenfolge().get(0));
        spielerMap.put(1, docSkatrunde.getSpielerReihenfolge().get(1));
        spielerMap.put(2, docSkatrunde.getSpielerReihenfolge().get(2));
        //check if 4 Spieler
        if (docSkatrunde.getSpielerReihenfolge().size() > 3) {
            spielerMap.put(3, docSkatrunde.getSpielerReihenfolge().get(3));
        }
        return spielerMap;
    }
}
