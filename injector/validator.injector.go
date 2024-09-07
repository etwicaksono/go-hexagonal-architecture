package injector

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var vld *validator.Validate

func validatorInit() *validator.Validate {
	if vld == nil {
		vld = validator.New()
		vld.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}

			return name
		})

	}
	return vld
}
