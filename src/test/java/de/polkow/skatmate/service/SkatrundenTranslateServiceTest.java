package de.polkow.skatmate.service;

import de.polkow.skatmate.model.AbrechnungsFormEnum;
import de.polkow.skatmate.model.DocSkatrunde;
import de.polkow.skatmate.model.DocSpiel;
import de.polkow.skatmate.model.SpielArtEnum;
import de.polkow.skatmate.persistence.entity.Skatrunde;
import de.polkow.skatmate.persistence.entity.Spielrunde;
import org.junit.jupiter.api.Test;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;

import static org.junit.jupiter.api.Assertions.assertEquals;

class SkatrundenTranslateServiceTest {

    @Test
    void translateSkatrundeToBierlachs() {
        //Given skatrunde
        LocalDateTime now = LocalDateTime.now();
        Skatrunde skatrunde = new Skatrunde();
        skatrunde.setId((long) 1);
        skatrunde.setSpieler(new TreeMap<>(Map.of(0, "Dennis", 1, "Thomas", 2, "Niclas", 3, "Stefan")));
        skatrunde.setAbrechnungsForm(AbrechnungsFormEnum.BIERLACHS);
        skatrunde.setTageszeit(now);
        skatrunde.setSpielrunden(List.of(
                new Spielrunde(1, "Thomas", false, 40, SpielArtEnum.HERZ),
                new Spielrunde(2, "Stefan", true, 33, SpielArtEnum.PIEK),
                new Spielrunde(3, "Thomas", true, 27, SpielArtEnum.KARO),
                new Spielrunde(4, "Niclas", true, 24, SpielArtEnum.KREUZ)));

        //When translate it to DocSkatrunde in Bierlachs
        DocSkatrunde docSkatrunde = SkatrundenTranslateService.translateSkatrunde(skatrunde, AbrechnungsFormEnum.BIERLACHS);

        //Then DocSkatrunde should look like this
        assertEquals(1, docSkatrunde.getId());
        assertEquals(AbrechnungsFormEnum.BIERLACHS, docSkatrunde.getAbrechnungsForm());

        assertEquals("Dennis", docSkatrunde.getSpielerReihenfolge().get(0));
        assertEquals("Thomas", docSkatrunde.getSpielerReihenfolge().get(1));
        assertEquals("Niclas", docSkatrunde.getSpielerReihenfolge().get(2));
        assertEquals("Stefan", docSkatrunde.getSpielerReihenfolge().get(3));

        assertEquals("Stefan", docSkatrunde.getPlazierung().get(0));
        assertEquals("Niclas", docSkatrunde.getPlazierung().get(1));
        assertEquals("Thomas", docSkatrunde.getPlazierung().get(2));
        assertEquals("Dennis", docSkatrunde.getPlazierung().get(3));

        List<DocSpiel> spielverlauf = docSkatrunde.getSpielverlauf();
        assertEquals(new DocSpiel(1, Map.of("Dennis", "x", "Thomas", "40", "Niclas", "", "Stefan", ""), 40, SpielArtEnum.HERZ),
                spielverlauf.get(0));
        assertEquals(new DocSpiel(2, Map.of("Dennis", "33", "Thomas", "x", "Niclas", "33", "Stefan", ""), 33, SpielArtEnum.PIEK),
                spielverlauf.get(1));
        assertEquals(new DocSpiel(3, Map.of("Dennis", "60", "Thomas","", "Niclas", "x", "Stefan", "27"), 27, SpielArtEnum.KARO),
                spielverlauf.get(2));
        assertEquals(new DocSpiel(4, Map.of("Dennis", "84", "Thomas", "64", "Niclas", "", "Stefan", "x"), 24, SpielArtEnum.KREUZ),
                spielverlauf.get(3));

    }

    @Test
    void translateSkatrundeToKlassisch() {
        //Given skatrunde
        LocalDateTime now = LocalDateTime.now();
        Skatrunde skatrunde = new Skatrunde();
        skatrunde.setId((long) 1);
        skatrunde.setSpieler(new TreeMap<>(Map.of(0, "Thomas", 1, "Niclas", 2, "Dennis")));
        skatrunde.setAbrechnungsForm(AbrechnungsFormEnum.KLASSISCH);
        skatrunde.setTageszeit(now);
        skatrunde.setSpielrunden(List.of(
                new Spielrunde(1, "Dennis", false, 54, SpielArtEnum.KARO),
                new Spielrunde(2, "Niclas", true, 22, SpielArtEnum.PIEK),
                new Spielrunde(3, "Dennis", true, 20, SpielArtEnum.HERZ)));

        //When translate it to DocSkatrunde in klassisch
        DocSkatrunde docSkatrunde = SkatrundenTranslateService.translateSkatrunde(skatrunde, AbrechnungsFormEnum.KLASSISCH);

        //Then DocSkatrunde should look like this#
        assertEquals(1, docSkatrunde.getId());
        assertEquals(AbrechnungsFormEnum.KLASSISCH, docSkatrunde.getAbrechnungsForm());

        assertEquals("Thomas", docSkatrunde.getSpielerReihenfolge().get(0));
        assertEquals("Niclas", docSkatrunde.getSpielerReihenfolge().get(1));
        assertEquals("Dennis", docSkatrunde.getSpielerReihenfolge().get(2));

        assertEquals("Niclas", docSkatrunde.getPlazierung().get(0));
        assertEquals("Thomas", docSkatrunde.getPlazierung().get(1));
        assertEquals("Dennis", docSkatrunde.getPlazierung().get(2));

        assertEquals(new DocSpiel(1, Map.of("Thomas", "", "Niclas", "", "Dennis", "-54"), -54, SpielArtEnum.KARO),
                docSkatrunde.getSpielverlauf().get(0));
        assertEquals(new DocSpiel(2, Map.of("Thomas", "", "Niclas", "22", "Dennis", ""), 22, SpielArtEnum.PIEK),
                docSkatrunde.getSpielverlauf().get(1));
        assertEquals(new DocSpiel(3, Map.of("Thomas", "", "Niclas","", "Dennis", "-34"), 20, SpielArtEnum.HERZ),
                docSkatrunde.getSpielverlauf().get(2));
    }
}