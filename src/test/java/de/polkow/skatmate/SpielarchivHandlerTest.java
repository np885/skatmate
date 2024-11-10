package de.polkow.skatmate;

import de.polkow.skatmate.model.AbrechnungsFormEnum;
import de.polkow.skatmate.model.DocSkatrunde;
import de.polkow.skatmate.model.DocSpiel;
import de.polkow.skatmate.model.SpielArtEnum;
import org.junit.jupiter.api.Test;

import java.time.OffsetDateTime;
import java.util.List;
import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class SpielarchivHandlerTest {

    SpielarchivHandler handler = new SpielarchivHandler();

/*    @Test
    public void testBerechnePlazierungBierlachs() {
        //Given
        DocSkatrunde docSkatrunde = new DocSkatrunde();
        docSkatrunde.setId((long) 1);
        docSkatrunde.setSpielerReihenfolge(List.of("Dennis", "Thomas", "Niclas", "Stefan"));
        docSkatrunde.setAbrechnungsForm(AbrechnungsFormEnum.BIERLACHS);
        docSkatrunde.setTageszeit(OffsetDateTime.now());
        docSkatrunde.setSpielverlauf(List.of(
            new DocSpiel(1, Map.of("Dennis", "x", "Thomas", "40", "Niclas", "", "Stefan", ""), 40, SpielArtEnum.HERZ),
            new DocSpiel(2, Map.of("Dennis", "33", "Thomas", "x", "Niclas", "33", "Stefan", ""), 33, SpielArtEnum.PIEK),
            new DocSpiel(3, Map.of("Dennis", "60", "Thomas","", "Niclas", "x", "Stefan", "27"), 27, SpielArtEnum.KARO),
            new DocSpiel(4, Map.of("Dennis", "84", "Thomas", "64", "Niclas", "", "Stefan", "x"), 24, SpielArtEnum.KREUZ)));

        //When
        DocSkatrunde result = handler.handleRequest(docSkatrunde, null);

        //Then
        assertEquals("Stefan", result.getPlazierung().get(0));
        assertEquals("Niclas", result.getPlazierung().get(1));
        assertEquals("Thomas", result.getPlazierung().get(2));
        assertEquals("Dennis", result.getPlazierung().get(3));
    }*/

/*    @Test
    public void testBerechnePlazierungKlassisch() {
        //Given
        DocSkatrunde docSkatrunde = new DocSkatrunde();
        docSkatrunde.setId((long) 1);
        docSkatrunde.setSpielerReihenfolge(List.of("Thomas", "Niclas", "Dennis"));
        docSkatrunde.setAbrechnungsForm(AbrechnungsFormEnum.KLASSISCH);
        docSkatrunde.setTageszeit(OffsetDateTime.now());
        docSkatrunde.setSpielverlauf(List.of(
                new DocSpiel(1, Map.of("Thomas", "", "Niclas", "", "Dennis", "-54"),  -54, SpielArtEnum.KARO),
                new DocSpiel(2, Map.of("Thomas", "", "Niclas", "22", "Dennis",""),  22, SpielArtEnum.PIEK),
                new DocSpiel(3, Map.of("Thomas", "",  "Niclas", "",  "Dennis", "-34"), 20, SpielArtEnum.HERZ)
                        ));
        //When
        DocSkatrunde result = handler.handleRequest(docSkatrunde, null);

        //Then
        assertEquals("Niclas", result.getPlazierung().get(0));
        assertEquals("Thomas", result.getPlazierung().get(1));
        assertEquals("Dennis", result.getPlazierung().get(2));
    }*/
}
