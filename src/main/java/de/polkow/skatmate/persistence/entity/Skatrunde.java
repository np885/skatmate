package de.polkow.skatmate.persistence.entity;

import de.polkow.skatmate.model.AbrechnungsFormEnum;
import lombok.Data;

import java.time.LocalDateTime;
import java.util.List;
import java.util.TreeMap;

@Data
public class Skatrunde {

    private Long id;
    private LocalDateTime tageszeit;
    private AbrechnungsFormEnum abrechnungsForm;
    private TreeMap<Integer, String> spieler;
    private List<Spielrunde> spielrunden;
}
