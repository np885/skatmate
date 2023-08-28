package de.polkow.skatmate.service;

import de.polkow.skatmate.exception.ValidationException;
import de.polkow.skatmate.exception.model.Violation;
import de.polkow.skatmate.model.DocSkatrunde;
import de.polkow.skatmate.model.DocSpiel;

import java.util.ArrayList;
import java.util.List;

public class ValidationInputService {

    public static void validateSkatrunde(DocSkatrunde docSkatrunde) {
        List<Violation> violations = new ArrayList<>();

        //Spieleranzahl 3 oder 4
        if (docSkatrunde.getSpielerReihenfolge().size() < 3 || docSkatrunde.getSpielerReihenfolge().size() > 4) {
            violations.add(new Violation("spielerReihenfolge", "Die Anzahl der Spieler muss entweder 3 oder 4 sein."));
        }

        //Tageszeit required
        if (docSkatrunde.getTageszeit() == null) {
            violations.add(new Violation("tageszeit", "Die Tageszeit ist nicht vorhanden."));
        }

        //Abrechnungsform required
        if (docSkatrunde.getAbrechnungsForm() == null) {
            violations.add(new Violation("abrechnungsform", "Die Abrechnungsform ist nicht vorhanden."));
        }

        //Validate Spielverlauf
        violations.addAll(validateSpielverlauf(docSkatrunde.getSpielverlauf()));

        //Prüfe ob Violations vorliegen
        if ( !violations.isEmpty() ) {
            throw new ValidationException("Error on Validation Skatrunden input", violations);
        }
    }

    private static List<Violation> validateSpielverlauf(List<DocSpiel> spielverlauf) {
        List<Violation> violations = new ArrayList<>();

        //Es muss mindestens 1 Spiel gespielt sein
        if (spielverlauf.isEmpty()) {
            violations.add(new Violation("spielverlauf", "Der Spielverlauf muss mindestens ein Spiel enthalten."));
        }

        for (DocSpiel docSpiel : spielverlauf) {
            violations.addAll(validateSpiel(docSpiel));
        }

        return violations;
    }

    private static List<Violation> validateSpiel(DocSpiel docSpiel) {
        List<Violation> violations = new ArrayList<>();

        //Jedes Spiel benötigt eine eindeutige Nr
        if (docSpiel.getNr() == null) {
            violations.add(new Violation("spielverlauf.nr", "Mindestens ein Spiel besitzt keine Nummer"));
        }

        //Jedes Spiel benötigt eingetragene Punkte
        if (docSpiel.getPunkte() == null) {
            violations.add(new Violation("spielverlauf.punkte", "Mindestens ein Spiel besitzt keine eingetragenen Punkte"));
        }

        //Jedes Spiel benötigt mindestens 3 und höchstens 4 SpielerWerte
        if (docSpiel.getSpielerWerteMap().size() < 3 || docSpiel.getSpielerWerteMap().size() > 4) {
            violations.add(new Violation("spielverlauf.spielerWerteMap", "Es müssen mindestens 3 und höchstens 4 Spielerwerte für jede Runde vorliegen"));
        }

        //check null value bei Spielerwerten
        for(String spielerWert : docSpiel.getSpielerWerteMap().values()) {
            if (spielerWert == null) {
                violations.add(new Violation("spielverlauf.spielerWerteMap", "Mindestens ein Spiel besitzt einen Null Wert"));
            }
        }

        return violations;
    }
}
