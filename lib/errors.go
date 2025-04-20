package lib

import "errors"

var (
	// lib.jwt_claims
	ErrFailedToParseJWTClaimsInContext = errors.New("failed to parse jwt claims in context")
	ErrJWTClaimsNotFoundInContext      = errors.New("jwt claims not found in context")
	ErrUnknownError                    = errors.New("unknown error")

	// user.service
	ErrUserNotFound         = errors.New("user not found")
	ErrSigningMethodInvalid = errors.New("signing method invalid")
	ErrUnauthorizedRequest  = errors.New("unauthorized request")

	// attachment.service
	ErrFailedToGetAttachments = errors.New("failed to get attachments")
	ErrAttachmentNotFound     = errors.New("attachment not found")

	// handler
	ErrFileCannotBeParsed  = errors.New("file cannot be parsed")
	ErrInvalidColumnLength = errors.New("invalid column length")
)
