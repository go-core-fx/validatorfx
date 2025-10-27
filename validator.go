package validatorfx

import "github.com/go-playground/validator/v10"

type Validator = validator.Validate

// New creates a new validator instance with required struct field validation enabled.
// WithRequiredStructEnabled ensures the "required" family of tags (e.g., required, required_if)
// apply to non-pointer struct fields, treating their zero values as invalid.
func New() *Validator {
	return validator.New(validator.WithRequiredStructEnabled())
}
