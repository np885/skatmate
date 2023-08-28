package de.polkow.skatmate.persistence.entity;

import de.polkow.skatmate.model.SpielArtEnum;
import lombok.Data;

@Data
public class Spielrunde {

    private Integer runde;
    private String reizGewinner;
    private boolean gewonnen;
    private Integer punkte;
    private SpielArtEnum spielArt;

}
