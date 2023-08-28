package de.polkow.skatmate.model;

import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.List;

import lombok.Data;

/**
 * Skatrunde
 */
@Data
public class Skatrunde {

  private Long id;

  private OffsetDateTime tageszeit;

  private List<String> spielerReihenfolge = new ArrayList<>();

  private AbrechnungsFormEnum abrechnungsForm;

  private List<String> plazierung;

  private List<Spiel> spielverlauf = new ArrayList<>();

}

