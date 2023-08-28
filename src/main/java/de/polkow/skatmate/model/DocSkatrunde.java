package de.polkow.skatmate.model;

import lombok.Data;

import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.List;

/**
 * Skatrunde
 */
@Data
public class DocSkatrunde {

  private Long id;

  private OffsetDateTime tageszeit;

  private List<String> spielerReihenfolge = new ArrayList<>();

  private AbrechnungsFormEnum abrechnungsForm;

  private List<String> plazierung;

  private List<DocSpiel> spielverlauf = new ArrayList<>();

}

