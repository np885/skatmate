package de.polkow.skatmate.model;

import lombok.Data;

/**
 * Spiel
 */
@Data
public class Spiel {

  private Integer nr;

  private String spieler1Wert;

  private String spieler2Wert;

  private String spieler3Wert;

  private String spieler4Wert;

  private Integer punkte;

  private SpielArtEnum spielArt;

  public Spiel(Integer nr, String spieler1Wert, String spieler2Wert, String spieler3Wert, String spieler4Wert,
      Integer punkte, SpielArtEnum spielArt) {
    this.nr = nr;
    this.spieler1Wert = spieler1Wert;
    this.spieler2Wert = spieler2Wert;
    this.spieler3Wert = spieler3Wert;
    this.spieler4Wert = spieler4Wert;
    this.punkte = punkte;
    this.spielArt = spielArt;
  }

  public Spiel(Integer nr, String spieler1Wert, String spieler2Wert, String spieler3Wert, Integer punkte, SpielArtEnum spielArt) {
    this.nr = nr;
    this.spieler1Wert = spieler1Wert;
    this.spieler2Wert = spieler2Wert;
    this.spieler3Wert = spieler3Wert;
    this.punkte = punkte;
    this.spielArt = spielArt;
  }

}

