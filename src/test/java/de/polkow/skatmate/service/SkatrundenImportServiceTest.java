package de.polkow.skatmate.service;

import de.polkow.skatmate.model.AbrechnungsFormEnum;
import de.polkow.skatmate.model.DocSkatrunde;
import de.polkow.skatmate.model.DocSpiel;
import de.polkow.skatmate.model.SpielArtEnum;
import de.polkow.skatmate.persistence.entity.Skatrunde;
import de.polkow.skatmate.persistence.entity.Spielrunde;
import org.junit.jupiter.api.Test;

import java.time.OffsetDateTime;
import java.util.List;
import java.util.Map;

import static org.junit.jupiter.api.Assertions.*;

class SkatrundenImportServiceTest {

    @Test
    public void importDocSkatrundeBierlachs() {
        //Given documented Skatrunde in bierlachs
        OffsetDateTime now = OffsetDateTime.now();
        DocSkatrunde docSkatrunde = new DocSkatrunde();
        docSkatrunde.setId((long) 1);
        docSkatrunde.setSpielerReihenfolge(List.of("Dennis", "Thomas", "Niclas", "Stefan"));
        docSkatrunde.setAbrechnungsForm(AbrechnungsFormEnum.BIERLACHS);
        docSkatrunde.setTageszeit(now);
        docSkatrunde.setSpielverlauf(List.of(
                new DocSpiel(1, Map.of("Dennis", "x", "Thomas", "40", "Niclas", "", "Stefan", ""), 40, SpielArtEnum.HERZ),
                new DocSpiel(2, Map.of("Dennis", "33", "Thomas", "x", "Niclas", "33", "Stefan", ""), 33, SpielArtEnum.PIEK),
                new DocSpiel(3, Map.of("Dennis", "60", "Thomas","", "Niclas", "x", "Stefan", "27"), 27, SpielArtEnum.KARO),
                new DocSpiel(4, Map.of("Dennis", "84", "Thomas", "64", "Niclas", "", "Stefan", "x"), 24, SpielArtEnum.KREUZ)));

        //When imported
        Skatrunde skatrunde = SkatrundenImportService.importDocSkatrunde(docSkatrunde);

        //Then skatrunde should look like this
        assertEquals(1, skatrunde.getId());
        assertEquals(now.toLocalDateTime(), skatrunde.getTageszeit());
        assertEquals(AbrechnungsFormEnum.BIERLACHS, skatrunde.getAbrechnungsForm());

        assertEquals("Dennis", skatrunde.getSpieler().get(0));
        assertEquals("Thomas", skatrunde.getSpieler().get(1));
        assertEquals("Niclas", skatrunde.getSpieler().get(2));
        assertEquals("Stefan", skatrunde.getSpieler().get(3));

        Spielrunde spielrunde1 = skatrunde.getSpielrunden().get(0);
        assertEquals(1, spielrunde1.getRunde());
        assertEquals("Thomas", spielrunde1.getReizGewinner());
        assertFalse(spielrunde1.isGewonnen());
        assertEquals(40, spielrunde1.getPunkte());
        assertEquals(SpielArtEnum.HERZ, spielrunde1.getSpielArt());

        Spielrunde spielrunde2 = skatrunde.getSpielrunden().get(1);
        assertEquals(2, spielrunde2.getRunde());
        assertEquals("Stefan", spielrunde2.getReizGewinner());
        assertTrue(spielrunde2.isGewonnen());
        assertEquals(33, spielrunde2.getPunkte());
        assertEquals(SpielArtEnum.PIEK, spielrunde2.getSpielArt());

        Spielrunde spielrunde3 = skatrunde.getSpielrunden().get(2);
        assertEquals(3, spielrunde3.getRunde());
        assertEquals("Thomas", spielrunde3.getReizGewinner());
        assertTrue(spielrunde3.isGewonnen());
        assertEquals(27, spielrunde3.getPunkte());
        assertEquals(SpielArtEnum.KARO, spielrunde3.getSpielArt());

        Spielrunde spielrunde4 = skatrunde.getSpielrunden().get(3);
        assertEquals(4, spielrunde4.getRunde());
        assertEquals("Niclas", spielrunde4.getReizGewinner());
        assertTrue(spielrunde4.isGewonnen());
        assertEquals(24, spielrunde4.getPunkte());
        assertEquals(SpielArtEnum.KREUZ, spielrunde4.getSpielArt());

    }

    @Test
    public void importDocSkatrundeKlassisch() {
        //Given documented Skatrunde in klassisch

        //When imported

        //Then skatrunde should look like this

    }
}