package de.polkow.skatmate.service;

import de.polkow.skatmate.exception.SkatmateBusinessException;
import de.polkow.skatmate.model.AbrechnungsFormEnum;
import de.polkow.skatmate.model.DocSkatrunde;
import de.polkow.skatmate.model.DocSpiel;
import de.polkow.skatmate.persistence.entity.Skatrunde;
import de.polkow.skatmate.persistence.entity.Spielrunde;
import org.apache.commons.lang3.math.NumberUtils;
import org.apache.commons.lang3.tuple.Pair;

import java.time.LocalDateTime;
import java.time.OffsetDateTime;
import java.time.ZoneId;
import java.time.ZoneOffset;
import java.util.*;
import java.util.stream.Collectors;

import static de.polkow.skatmate.model.AbrechnungsFormEnum.BIERLACHS;
import static de.polkow.skatmate.model.AbrechnungsFormEnum.KLASSISCH;

public class SkatrundenTranslateService {

    public static DocSkatrunde translateSkatrunde(Skatrunde skatrunde, AbrechnungsFormEnum abrechnungsForm) {
        DocSkatrunde docSkatrunde = new DocSkatrunde();
        docSkatrunde.setId(skatrunde.getId());
        docSkatrunde.setTageszeit(toOffsetDateTime(skatrunde.getTageszeit()));
        docSkatrunde.setAbrechnungsForm(abrechnungsForm);
        docSkatrunde.setSpielerReihenfolge(skatrunde.getSpieler().values().stream().toList());
        //Stelle Spielverlauf her
        docSkatrunde.setSpielverlauf(erstelleSpielverlauf(skatrunde, abrechnungsForm));

        //auf Basis von Spielverlauf ermittel Platzierung
        docSkatrunde.setPlazierung(berechnePlatzierung(docSkatrunde.getSpielverlauf(), abrechnungsForm));
        return docSkatrunde;
    }

    private static List<DocSpiel> erstelleSpielverlauf(Skatrunde skatrunde, AbrechnungsFormEnum abrechnungsForm) {
        List<DocSpiel> spielverlauf = new ArrayList<>(skatrunde.getSpielrunden().size());
        //Merke den letzten Spielerwert für eine einfach addition
        Map<String, Integer> letzterWertProSpieler = new HashMap<>(skatrunde.getSpieler().size());
        //Jede Spielrunde wird analysiert
        for (Spielrunde runde : skatrunde.getSpielrunden()) {
            Integer punkteNew = runde.getPunkte();
            if ((BIERLACHS == abrechnungsForm && runde.getPunkte() < 0) || (KLASSISCH == abrechnungsForm && !runde.isGewonnen())) {
                punkteNew = punkteNew * -1;
            }
            DocSpiel docSpiel = new DocSpiel(runde.getRundeNr(), punkteNew, runde.getSpielArt());
            Map<String, String> spielerWerteMap = new HashMap<>(skatrunde.getSpieler().size());
            //Pro position und Spieler wird der Wert berechnet {x, zahl o. ""}
            for (Integer position : skatrunde.getSpieler().keySet()) {
                String spieler = skatrunde.getSpieler().get(position);
                //initialisiere Map für jeden Spieler einmalig
                letzterWertProSpieler.putIfAbsent(spieler, 0);

                //berechne Wertepaar; einmal als String value (left) {x, zahl o. ""} und als Integer Wert (right)
                Pair<String, Integer> wertePaar = berechneWert(spieler, position, letzterWertProSpieler.get(spieler),
                        runde, punkteNew, skatrunde.getSpieler().size(), abrechnungsForm);

                //Aktualisiere aktuellen Wert
                if (wertePaar.getRight() != 0) {
                    letzterWertProSpieler.put(spieler, wertePaar.getRight());
                }
                //schreibe String Wert in die SpielerMap
                spielerWerteMap.put(spieler, wertePaar.getLeft());
            }

            docSpiel.setSpielerWerteMap(spielerWerteMap);
            spielverlauf.add(docSpiel);
        }

        return spielverlauf;

    }

    private static Pair<String, Integer> berechneWert(String spieler, Integer sitzPosition, Integer letzterSpielerWert,
                                                      Spielrunde runde, Integer punkteNew, int spielerAnzahl,
                                                      AbrechnungsFormEnum abrechnungsForm) {
        //Setzt Spieler aus?
        if (spielerAnzahl == 4 && (((runde.getRundeNr() - 1) + spielerAnzahl) % spielerAnzahl == sitzPosition)) {
            return Pair.of("x", 0);
        }

        //berechne neuen Spielwert
        Integer neuerSpielerWert = letzterSpielerWert != null? punkteNew + letzterSpielerWert : punkteNew;

        if (AbrechnungsFormEnum.BIERLACHS == abrechnungsForm) {
            //Ist Spieler Reizgewinner
            if (runde.getReizGewinner().equalsIgnoreCase(spieler)) {
                if (!runde.isGewonnen()) {
                    return Pair.of(neuerSpielerWert.toString(), neuerSpielerWert);
                }
            } else {
                if (runde.isGewonnen()) {
                    return Pair.of(neuerSpielerWert.toString(), neuerSpielerWert);
                }
            }
            return Pair.of("", 0);
        } else if (AbrechnungsFormEnum.KLASSISCH == abrechnungsForm) {
            if (runde.getReizGewinner().equalsIgnoreCase(spieler)) {
                return Pair.of(neuerSpielerWert.toString(), neuerSpielerWert);
            }
            return Pair.of("", 0);
        }

        throw new SkatmateBusinessException("Unsupported Abrechnungsform: " + abrechnungsForm);
    }

    private static List<String> berechnePlatzierung(List<DocSpiel> spielverlauf, AbrechnungsFormEnum abrechnungsForm) {
        //Suche die letzte Zahl für jeden Spielerwert im Spielverlauf
        Map<String, Integer> spielerMaxWerte = getSpielerMaxwerte(spielverlauf);

        if (AbrechnungsFormEnum.BIERLACHS == abrechnungsForm) {
            return spielerMaxWerte.entrySet().stream().sorted(Map.Entry.comparingByValue())
                    .map(Map.Entry::getKey).collect(Collectors.toList());
        } else if (AbrechnungsFormEnum.KLASSISCH == abrechnungsForm) {
            return spielerMaxWerte.entrySet().stream().sorted(Map.Entry.comparingByValue(Comparator.reverseOrder()))
                    .map(Map.Entry::getKey).collect(Collectors.toList());
        } else {
            throw new RuntimeException("Abrechnungsform unbekannt. Platzierung kann nicht berechnet werden.");
        }
    }

    private static Map<String, Integer> getSpielerMaxwerte(List<DocSpiel> spielverlauf) {
        Map<String, Integer> spielerMaxWerte = new HashMap<>();
        //Durchlaufe den Spielverlauf von hinten nach vorne
        ListIterator<DocSpiel> it = spielverlauf.listIterator(spielverlauf.size());
        while(it.hasPrevious()) {
            DocSpiel docSpiel = it.previous();
            for (String spieler : docSpiel.getSpielerWerteMap().keySet()) {
                String spielerWert = docSpiel.getSpielerWerteMap().get(spieler);
                if (!spielerMaxWerte.containsKey(spieler)) {
                    spielerMaxWerte.put(spieler, 0);
                }
                if (NumberUtils.isParsable(spielerWert) && spielerMaxWerte.get(spieler) == 0) {
                    spielerMaxWerte.put(spieler, Integer.parseInt(spielerWert));
                }
            }
        }

        return spielerMaxWerte;
    }

    private static OffsetDateTime toOffsetDateTime(LocalDateTime localDateTime) {
        final ZoneId zone = ZoneId.of("Europe/Berlin");
        ZoneOffset zoneOffSet = zone.getRules().getOffset(localDateTime);
        return localDateTime.atOffset(zoneOffSet);
    }
}
