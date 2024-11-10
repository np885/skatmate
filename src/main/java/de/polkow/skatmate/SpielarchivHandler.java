package de.polkow.skatmate;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;
import de.polkow.skatmate.model.AbrechnungsFormEnum;
import de.polkow.skatmate.model.DocSkatrunde;
import de.polkow.skatmate.model.DocSpiel;
import de.polkow.skatmate.service.ValidationInputService;
import org.apache.commons.lang3.math.NumberUtils;

import java.util.*;
import java.util.stream.Collectors;

public class SpielarchivHandler implements RequestHandler<String, String> {

    @Override
    public String handleRequest(String inputString, Context context) {
        DocSkatrunde input = new DocSkatrunde();
        //validate Input
        ValidationInputService.validateSkatrunde(input);

        DocSkatrunde output = input;
        //Berechne Platzierung
        output.setPlazierung(berechnePlatzierung(input.getSpielverlauf(), input.getAbrechnungsForm()));

        return "Hello World: " + inputString;
    }

    private List<String> berechnePlatzierung(List<DocSpiel> spielverlauf, AbrechnungsFormEnum abrechnungsForm) {
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

    private Map<String, Integer> getSpielerMaxwerte(List<DocSpiel> spielverlauf) {
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
    
}
