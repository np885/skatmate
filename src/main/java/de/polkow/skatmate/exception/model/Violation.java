package de.polkow.skatmate.exception.model;

import lombok.Data;

@Data
public class Violation {

    private String field;
    private String message;

    public Violation(String field, String message) {
        this.field = field;
        this.message = message;
    }
}
