package validatorfx

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator = validator.Validate

// New creates a new validator instance with required struct field validation enabled.
// WithRequiredStructEnabled ensures the "required" family of tags (e.g., required, required_if)
// apply to non-pointer struct fields, treating their zero values as invalid.
func New() *Validator {
	v := validator.New(validator.WithRequiredStructEnabled())
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		//nolint:mnd //fixed length
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		if name == "" {
			return fld.Name
		}
		return name
	})

	return v
}
