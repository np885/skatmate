package de.polkow.skatmate.persistence.entity;

import de.polkow.skatmate.model.SpielArtEnum;
import lombok.Data;

@Data
public class Spielrunde {

    private Integer rundeNr;
    private String reizGewinner;
    private boolean gewonnen;
    private Integer punkte;
    private SpielArtEnum spielArt;

    public Spielrunde(Integer rundeNr, String reizGewinner, boolean gewonnen, Integer punkte, SpielArtEnum spielArt) {
        this.rundeNr = rundeNr;
        this.reizGewinner = reizGewinner;
        this.gewonnen = gewonnen;
        this.punkte = punkte;
        this.spielArt = spielArt;
    }

    public Spielrunde() {
    }
}
