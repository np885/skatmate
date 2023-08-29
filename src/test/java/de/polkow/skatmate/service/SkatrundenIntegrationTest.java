package de.polkow.skatmate.service;

import de.polkow.skatmate.model.AbrechnungsFormEnum;
import de.polkow.skatmate.model.DocSkatrunde;
import de.polkow.skatmate.model.DocSpiel;
import de.polkow.skatmate.model.SpielArtEnum;
import de.polkow.skatmate.persistence.entity.Skatrunde;
import org.junit.jupiter.api.Test;

import java.time.OffsetDateTime;
import java.util.List;
import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

public class SkatrundenIntegrationTest {

    @Test
    public void translateDocSkatrundeBierlachsToSkatrundeAndBack() {
        //Given DocSkatrunde in bierlachs Notation
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

        //When import to Skatrunde & translate to Bierlachs
        Skatrunde skatrunde = SkatrundenImportService.importDocSkatrunde(docSkatrunde);
        DocSkatrunde docSkatrundeTranslated = SkatrundenTranslateService.translateSkatrunde(skatrunde, AbrechnungsFormEnum.BIERLACHS);

        //Then it should be the same as given
        assertEquals(docSkatrunde.getSpielverlauf(), docSkatrundeTranslated.getSpielverlauf());
        assertEquals(docSkatrunde.getSpielerReihenfolge(), docSkatrundeTranslated.getSpielerReihenfolge());
        assertEquals(docSkatrunde.getId(), docSkatrundeTranslated.getId());
        assertEquals(docSkatrunde.getTageszeit(), docSkatrundeTranslated.getTageszeit());
        assertEquals(docSkatrunde.getAbrechnungsForm(), docSkatrundeTranslated.getAbrechnungsForm());
        assertNotNull(docSkatrundeTranslated.getPlazierung());

    }

    @Test
    public void translateDocSkatrundeKlassischToSkatrundeAndBack() {
        //Given DocSkatrunde in klassisch Notation
        OffsetDateTime now = OffsetDateTime.now();
        DocSkatrunde docSkatrunde = new DocSkatrunde();
        docSkatrunde.setId((long) 1);
        docSkatrunde.setSpielerReihenfolge(List.of("Thomas", "Niclas", "Dennis"));
        docSkatrunde.setAbrechnungsForm(AbrechnungsFormEnum.KLASSISCH);
        docSkatrunde.setTageszeit(now);
        docSkatrunde.setSpielverlauf(List.of(
                new DocSpiel(1, Map.of("Thomas", "", "Niclas", "", "Dennis", "-54"), -54, SpielArtEnum.KARO),
                new DocSpiel(2, Map.of("Thomas", "", "Niclas", "22", "Dennis", ""), 22, SpielArtEnum.PIEK),
                new DocSpiel(3, Map.of("Thomas", "", "Niclas","", "Dennis", "-34"), 20, SpielArtEnum.HERZ)));

        //When import to Skatrunde & translate to klassisch
        Skatrunde skatrunde = SkatrundenImportService.importDocSkatrunde(docSkatrunde);
        DocSkatrunde docSkatrundeTranslated = SkatrundenTranslateService.translateSkatrunde(skatrunde, AbrechnungsFormEnum.KLASSISCH);

        //Then it should be the same as given
        assertEquals(docSkatrunde.getSpielverlauf(), docSkatrundeTranslated.getSpielverlauf());
        assertEquals(docSkatrunde.getSpielerReihenfolge(), docSkatrundeTranslated.getSpielerReihenfolge());
        assertEquals(docSkatrunde.getId(), docSkatrundeTranslated.getId());
        assertEquals(docSkatrunde.getTageszeit(), docSkatrundeTranslated.getTageszeit());
        assertEquals(docSkatrunde.getAbrechnungsForm(), docSkatrundeTranslated.getAbrechnungsForm());
        assertNotNull(docSkatrundeTranslated.getPlazierung());
    }

    @Test
    public void translateDocSkatrundeKlassischToBierlachs() {
        //Given DocSkatrunde in klassisch Notation
        OffsetDateTime now = OffsetDateTime.now();
        DocSkatrunde docSkatrunde = new DocSkatrunde();
        docSkatrunde.setId((long) 1);
        docSkatrunde.setSpielerReihenfolge(List.of("Thomas", "Niclas", "Dennis"));
        docSkatrunde.setAbrechnungsForm(AbrechnungsFormEnum.KLASSISCH);
        docSkatrunde.setTageszeit(now);
        docSkatrunde.setSpielverlauf(List.of(
                new DocSpiel(1, Map.of("Thomas", "", "Niclas", "", "Dennis", "-54"), -54, SpielArtEnum.KARO),
                new DocSpiel(2, Map.of("Thomas", "", "Niclas", "22", "Dennis", ""), 22, SpielArtEnum.PIEK),
                new DocSpiel(3, Map.of("Thomas", "", "Niclas","", "Dennis", "-34"), 20, SpielArtEnum.HERZ)));

        //When import to Skatrunde & translate to bierlachs
        Skatrunde skatrunde = SkatrundenImportService.importDocSkatrunde(docSkatrunde);
        DocSkatrunde docSkatrundeBierlachs = SkatrundenTranslateService.translateSkatrunde(skatrunde, AbrechnungsFormEnum.BIERLACHS);

        //Then
        assertEquals(1, docSkatrundeBierlachs.getId());
        assertEquals(AbrechnungsFormEnum.BIERLACHS, docSkatrundeBierlachs.getAbrechnungsForm());
        assertEquals(docSkatrunde.getSpielerReihenfolge(), docSkatrundeBierlachs.getSpielerReihenfolge());

        assertEquals("Niclas", docSkatrundeBierlachs.getPlazierung().get(0));
        assertEquals("Thomas", docSkatrundeBierlachs.getPlazierung().get(1));
        assertEquals("Dennis", docSkatrundeBierlachs.getPlazierung().get(2));

        List<DocSpiel> spielverlauf = docSkatrundeBierlachs.getSpielverlauf();
        assertEquals(new DocSpiel(1, Map.of("Thomas", "", "Niclas", "", "Dennis", "54"), 54, SpielArtEnum.KARO),
                spielverlauf.get(0));
        assertEquals(new DocSpiel(2, Map.of("Thomas", "22", "Niclas", "", "Dennis", "76"), 22, SpielArtEnum.PIEK),
                spielverlauf.get(1));
        assertEquals(new DocSpiel(3, Map.of("Thomas", "42", "Niclas","20", "Dennis", ""), 20, SpielArtEnum.HERZ),
                spielverlauf.get(2));

    }

    @Test
    public void translateDocSkatrundeBierlachsToKlassisch() {
        //Given DocSkatrunde in bierlachs Notation

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

        //When import to Skatrunde & translate to klassisch
        Skatrunde skatrunde = SkatrundenImportService.importDocSkatrunde(docSkatrunde);
        DocSkatrunde docSkatrundeKlassisch = SkatrundenTranslateService.translateSkatrunde(skatrunde, AbrechnungsFormEnum.KLASSISCH);

        //Then
        assertEquals(docSkatrunde.getId(), docSkatrundeKlassisch.getId());
        assertEquals(AbrechnungsFormEnum.KLASSISCH, docSkatrundeKlassisch.getAbrechnungsForm());
        assertEquals(docSkatrunde.getSpielerReihenfolge(), docSkatrundeKlassisch.getSpielerReihenfolge());

        assertEquals("Stefan", docSkatrundeKlassisch.getPlazierung().get(0));
        assertEquals("Niclas", docSkatrundeKlassisch.getPlazierung().get(1));
        assertEquals("Dennis", docSkatrundeKlassisch.getPlazierung().get(2));
        assertEquals("Thomas", docSkatrundeKlassisch.getPlazierung().get(3));

        assertEquals(new DocSpiel(1, Map.of("Dennis", "x", "Thomas", "-40", "Niclas", "", "Stefan", ""), -40, SpielArtEnum.HERZ),
                docSkatrundeKlassisch.getSpielverlauf().get(0));
        assertEquals(new DocSpiel(2, Map.of("Dennis", "", "Thomas", "x", "Niclas", "", "Stefan", "33"), 33, SpielArtEnum.PIEK),
                docSkatrundeKlassisch.getSpielverlauf().get(1));
        assertEquals(new DocSpiel(3, Map.of("Dennis", "", "Thomas","-13", "Niclas", "x", "Stefan", ""), 27, SpielArtEnum.KARO),
                docSkatrundeKlassisch.getSpielverlauf().get(2));
        assertEquals(new DocSpiel(4, Map.of("Dennis", "", "Thomas","", "Niclas", "24", "Stefan", "x"), 24, SpielArtEnum.KREUZ),
                docSkatrundeKlassisch.getSpielverlauf().get(3));

    }

}
