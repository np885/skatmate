package de.polkow.skatmate.exception;

import de.polkow.skatmate.exception.model.Violation;

import java.util.List;

public class ValidationException extends RuntimeException {

    public ValidationException(String message, List<Violation> violationList) {
        super(message + " " + violationList.toString());
    }

}
