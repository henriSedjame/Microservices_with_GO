package models

// GenericError represents the generic error of application
//
// swagger:model
type GenericError struct {
	// the error message
	Message string
}

// ValidationError represents the error occured when request body is not valid
//
// swagger:model
type ValidationError struct {
	// the error message
	Messages string
}
