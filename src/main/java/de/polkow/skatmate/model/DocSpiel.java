package de.polkow.skatmate.model;

import lombok.Data;

import java.util.Map;

/**
 * Spiel
 */
@Data
public class DocSpiel {

  private Integer nr;

  private Map<String, String> spielerWerteMap;

  private Integer punkte;

  private SpielArtEnum spielArt;

  public DocSpiel(Integer nr, Map<String, String> spielerWerteMap, Integer punkte, SpielArtEnum spielArt) {
    this.nr = nr;
    this.spielerWerteMap = spielerWerteMap;
    this.punkte = punkte;
    this.spielArt = spielArt;
  }

  public DocSpiel(Integer nr, Integer punkte, SpielArtEnum spielArt) {
    this.nr = nr;
    this.punkte = punkte;
    this.spielArt = spielArt;
  }

}

