package de.polkow.skatmate.model;

import lombok.Getter;

@Getter
public enum SpielArtEnum {
    KREUZ("Kreuz"),
    
    PIEK("Piek"),
    
    HERZ("Herz"),
    
    KARO("Karo"),
    
    GRAND("Grand"),
    
    NULLSPIEL("Nullspiel")/*,
    TODO: Ramsch aktuell noch nicht unterstützt
    //RAMSCH("Ramsch")*/;

    private final String value;

    SpielArtEnum(String value) {
      this.value = value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static SpielArtEnum fromValue(String value) {
      for (SpielArtEnum b : SpielArtEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }
}
