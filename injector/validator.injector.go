package injector

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var validatorInstance *validator.Validate

func validatorProvider() *validator.Validate {
	if validatorInstance != nil {
		return validatorInstance
	}

	validatorInstance = validator.New()
	validatorInstance.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	return validatorInstance
}
