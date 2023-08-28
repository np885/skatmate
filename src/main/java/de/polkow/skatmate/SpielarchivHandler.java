package de.polkow.skatmate;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;
import de.polkow.skatmate.exception.ValidationException;
import de.polkow.skatmate.exception.model.Violation;
import de.polkow.skatmate.model.AbrechnungsFormEnum;
import de.polkow.skatmate.model.Skatrunde;
import de.polkow.skatmate.model.Spiel;
import lombok.SneakyThrows;
import org.apache.commons.lang3.math.NumberUtils;

import java.util.*;
import java.util.stream.Collectors;

public class SpielarchivHandler implements RequestHandler<Skatrunde, Skatrunde> {

    @SneakyThrows
    @Override
    public Skatrunde handleRequest(Skatrunde input, Context context) {
        //validate Input
        validateSkatrunde(input);

        Skatrunde output = input;
        //Berechne Platzierung
        output.setPlazierung(berechnePlatzierung(input.getSpielverlauf(), input.getAbrechnungsForm(), input.getSpielerReihenfolge()));

        return output;
    }

    private void validateSkatrunde(Skatrunde skatrunde) {
        List<Violation> violations = new ArrayList<>();

        //Spieleranzahl 3 oder 4
        if (skatrunde.getSpielerReihenfolge().size() < 3 || skatrunde.getSpielerReihenfolge().size() > 4) {
            violations.add(new Violation("spielerReihenfolge", "Die Anzahl der Spieler muss entweder 3 oder 4 sein."));
        }

        //Tageszeit required
        if (skatrunde.getTageszeit() == null) {
            violations.add(new Violation("tageszeit", "Die Tageszeit ist nicht vorhanden."));
        }

        //Abrechnungsform required
        if (skatrunde.getAbrechnungsForm() == null) {
            violations.add(new Violation("abrechnungsform", "Die Abrechnungsform ist nicht vorhanden."));
        }

        //Validate Spielverlauf
        violations.addAll(validateSpielverlauf(skatrunde.getSpielverlauf()));

        //Prüfe ob Violations vorliegen
        if ( !violations.isEmpty() ) {
            throw new ValidationException("Error on Validation Skatrunden input", violations);
        }
    }

    private List<Violation> validateSpielverlauf(List<Spiel> spielverlauf) {
        List<Violation> violations = new ArrayList<>();

        //Es muss mindestens 1 Spiel gespielt sein
        if (spielverlauf.isEmpty()) {
            violations.add(new Violation("spielverlauf", "Der Spielverlauf muss mindestens ein Spiel enthalten."));
        }

        for (Spiel spiel : spielverlauf) {
            violations.addAll(validateSpiel(spiel));
        }

        return violations;
    }

    private List<Violation> validateSpiel(Spiel spiel) {
        List<Violation> violations = new ArrayList<>();

        //Jedes Spiel benötigt eine eindeutige Nr
        if (spiel.getNr() == null) {
            violations.add(new Violation("spielverlauf.nr", "Mindestens ein Spiel besitzt keine Nummer"));
        }

        //Jedes Spiel benötigt eingetragene Punkte
        if (spiel.getPunkte() == null) {
            violations.add(new Violation("spielverlauf.punkte", "Mindestens ein Spiel besitzt keine eingetragenen Punkte"));
        }

        //Jedes Spiel benötigt Spielerwerte für 1-3 Spieler
        if (spiel.getSpieler1Wert() == null) {
            violations.add(new Violation("spielverlauf.spielerwert1", "Mindestens ein Spiel besitzt für Spieler 1 keinen Wert"));
        }

        //Jedes Spiel benötigt Spielerwerte für 1-3 Spieler
        if (spiel.getSpieler2Wert() == null) {
            violations.add(new Violation("spielverlauf.spielerwert2", "Mindestens ein Spiel besitzt für Spieler 3 keinen Wert"));
        }

        //Jedes Spiel benötigt Spielerwerte für 1-3 Spieler
        if (spiel.getSpieler3Wert() == null) {
            violations.add(new Violation("spielverlauf.spielerwert3", "Mindestens ein Spiel besitzt für Spieler 3 keinen Wert"));
        }

        return violations;
    }


    private List<String> berechnePlatzierung(List<Spiel> spielverlauf, AbrechnungsFormEnum abrechnungsForm, List<String> spieler) {
        //Suche die letzte Zahl für jeden Spielerwert im Spielverlauf
        List<Integer> spielwerteFuerSpieler = getSpielwerteFuerSpieler(spielverlauf);
        Map<String, Integer> spielerWerteMap = new HashMap<>(4);
        for (int i = 0; i < spieler.size(); i++) {
            spielerWerteMap.put(spieler.get(i), spielwerteFuerSpieler.get(i));
        }
        if (AbrechnungsFormEnum.BIERLACHS == abrechnungsForm) {
            return spielerWerteMap.entrySet().stream().sorted(Map.Entry.comparingByValue())
                .map(Map.Entry::getKey).collect(Collectors.toList());
        } else if (AbrechnungsFormEnum.KLASSISCH == abrechnungsForm) {
            return spielerWerteMap.entrySet().stream().sorted(Map.Entry.comparingByValue(Comparator.reverseOrder()))
                .map(Map.Entry::getKey).collect(Collectors.toList());
        } else {
            throw new RuntimeException("Abrechnungsform unbekannt. Platzierung kann nicht berechnet werden.");
        }
    }

    private List<Integer> getSpielwerteFuerSpieler(List<Spiel> spielverlauf) {
        //Initialisiere Spieler mit 0 Punkten
        Integer[] spielerPunkte = {0,0,0,0};
        //Durchlaufe den Spielverlauf von hinten nach vorne
        ListIterator<Spiel> it = spielverlauf.listIterator(spielverlauf.size());
        while(it.hasPrevious()) {
            Spiel spiel = it.previous();
            if (NumberUtils.isParsable(spiel.getSpieler1Wert()) && spielerPunkte[0] == 0) {
                spielerPunkte[0] = Integer.parseInt(spiel.getSpieler1Wert());
            }
            if (NumberUtils.isParsable(spiel.getSpieler2Wert()) && spielerPunkte[1] == 0) {
                spielerPunkte[1] = Integer.parseInt(spiel.getSpieler2Wert());
            }
            if (NumberUtils.isParsable(spiel.getSpieler3Wert()) && spielerPunkte[2] == 0) {
                spielerPunkte[2] = Integer.parseInt(spiel.getSpieler3Wert());
            }
            if (NumberUtils.isParsable(spiel.getSpieler4Wert()) && spielerPunkte[3] == 0) {
                spielerPunkte[3] = Integer.parseInt(spiel.getSpieler4Wert());
            }
        }

        return Arrays.asList(spielerPunkte);
    }
    
}
