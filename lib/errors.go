package lib

import "errors"

var (
	// lib.jwt_claims
	ErrFailedToParseJWTClaimsInContext = errors.New("failed to parse jwt claims in context")
	ErrJWTClaimsNotFoundInContext      = errors.New("jwt claims not found in context")

	// user.service
	ErrIncorrectPassword    = errors.New("incorrect password")
	ErrSigningMethodInvalid = errors.New("signing method invalid")
	ErrUnauthorizedRequest  = errors.New("unauthorized request")
)
