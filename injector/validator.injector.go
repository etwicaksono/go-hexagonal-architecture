package injector

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/validation_util"
	"github.com/go-playground/validator/v10"
)

var validatorInstance *validator.Validate

// validatorProvider returns an instance of validator.Validate, which is used to
// validate structs.  The first time this is called, it creates a new instance
// of validator.Validate and registers a tag name function that sets the name
// of each field to the value of its `json` tag.  Subsequent calls return the
// same instance.
func validatorProvider() *validator.Validate {
	if validatorInstance != nil {
		return validatorInstance
	}

	validatorInstance = validation_util.NewValidator()
	return validatorInstance
}
