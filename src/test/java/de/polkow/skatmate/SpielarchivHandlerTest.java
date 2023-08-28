package de.polkow.skatmate;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.time.OffsetDateTime;
import java.util.List;

import org.junit.jupiter.api.Test;

import de.polkow.skatmate.model.AbrechnungsFormEnum;
import de.polkow.skatmate.model.Skatrunde;
import de.polkow.skatmate.model.Spiel;
import de.polkow.skatmate.model.SpielArtEnum;

public class SpielarchivHandlerTest {

    SpielarchivHandler handler = new SpielarchivHandler();

    @Test
    public void testBerechnePlazierungBierlachs() {
        //Given
        Skatrunde skatrunde = new Skatrunde();
        skatrunde.setId((long) 1);
        skatrunde.setSpielerReihenfolge(List.of("Dennis", "Thomas", "Niclas", "Stefan"));
        skatrunde.setAbrechnungsForm(AbrechnungsFormEnum.BIERLACHS);
        skatrunde.setTageszeit(OffsetDateTime.now());
        skatrunde.setSpielverlauf(List.of(
            new Spiel(1, "x", "40", "", "", 40, SpielArtEnum.HERZ), 
            new Spiel(2, "33", "x", "33", "", 33, SpielArtEnum.PIEK), 
            new Spiel(3, "60", "", "x", "27", 27, SpielArtEnum.KARO), 
            new Spiel(4, "84", "64", "", "x", 24, SpielArtEnum.KREUZ)));

        //When
        Skatrunde result = handler.handleRequest(skatrunde, null);

        //Then
        assertEquals("Stefan", result.getPlazierung().get(0));
        assertEquals("Niclas", result.getPlazierung().get(1));
        assertEquals("Thomas", result.getPlazierung().get(2));
        assertEquals("Dennis", result.getPlazierung().get(3));
    }

    @Test
    public void testBerechnePlazierungKlassisch() {
        //Given
        Skatrunde skatrunde = new Skatrunde();
        skatrunde.setId((long) 1);
        skatrunde.setSpielerReihenfolge(List.of("Thomas", "Niclas", "Dennis"));
        skatrunde.setAbrechnungsForm(AbrechnungsFormEnum.KLASSISCH);
        skatrunde.setTageszeit(OffsetDateTime.now());
        skatrunde.setSpielverlauf(List.of(
                new Spiel(1, "", "", "-54",  -54, SpielArtEnum.KARO),
                new Spiel(2, "", "22", "", "", 22, SpielArtEnum.PIEK),
                new Spiel(3, "", "",  "-34", 20, SpielArtEnum.HERZ)
                        ));
        //When
        Skatrunde result = handler.handleRequest(skatrunde, null);

        //Then
        assertEquals("Niclas", result.getPlazierung().get(0));
        assertEquals("Thomas", result.getPlazierung().get(1));
        assertEquals("Dennis", result.getPlazierung().get(2));
    }
}
