package de.polkow.skatmate.model;

import lombok.Getter;

@Getter
public enum AbrechnungsFormEnum {
    BIERLACHS("bierlachs"),
    
    KLASSISCH("klassisch");

    private final String value;

    AbrechnungsFormEnum(String value) {
      this.value = value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static AbrechnungsFormEnum fromValue(String value) {
      for (AbrechnungsFormEnum b : AbrechnungsFormEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }
}
