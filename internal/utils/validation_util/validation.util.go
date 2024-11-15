package validation_util

import (
	"errors"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/string_util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sagikazarmark/slog-shim"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

const (
	isUsernameTag = "is-username"
)

func NewValidator() *validator.Validate {
	validatorInstance := validator.New()
	validatorInstance.RegisterTagNameFunc(func(fld reflect.StructField) string { // This code sets up a validator instance and configures it to use the json tag from struct fields to determine the validation tag name. If the json tag is set to "-", it will ignore that field during validation.
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	// Register the custom validations
	registerCustomValidations(validatorInstance)

	return validatorInstance
}

func registerCustomValidations(validatorInstance *validator.Validate) {
	err := validatorInstance.RegisterValidation(isUsernameTag, isUsernameValid)
	if err != nil {
		slog.Error("Failed to register is-username validation", slog.String(constants.Error, err.Error()))
	}
}

func translateErrorMessage(err error) (errValidation fiber.Map) {
	// make error map
	errValidation = make(fiber.Map)
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			switch fieldError.Tag() {
			case "required":
				{
					errValidation[fieldName] = fmt.Sprint(fieldName, " field is required")
				}
			case "email":
				{
					errValidation[fieldName] = fmt.Sprint(fieldName, " field must be valid email format")
				}
			case "min":
				{
					if fieldError.Kind() == reflect.String {
						errValidation[fieldName] = fmt.Sprint(fieldName, " field must be longer than ", fieldError.Param(), " characters")
					} else {
						errValidation[fieldName] = fmt.Sprint(fieldName, " field must be greater than ", fieldError.Param())
					}
				}
			case "max":
				{
					if fieldError.Kind() == reflect.String {
						errValidation[fieldName] = fmt.Sprint(fieldName, " field must not be longer than ", fieldError.Param(), " characters")
					} else {
						errValidation[fieldName] = fmt.Sprint(fieldName, " field must not be greater than ", fieldError.Param())
					}
				}
			case "required_with":
				{
					fieldParam := fieldError.Param()
					fieldSlice := strings.Split(fieldParam, " ")
					for i, field := range fieldSlice {

						// Convert the first character to lowercase
						firstCharLower := strings.ToLower(string(field[0]))

						// Convert the last character to lowercase
						lastCharLower := strings.ToLower(string(field[len(field)-1]))

						// Combine the modified first and last characters with the rest of the string
						fieldSlice[i] = firstCharLower + field[1:len(field)-1] + lastCharLower
					}
					errValidation[fieldName] = fmt.Sprint(fieldName, " field is required when ", strings.Join(fieldSlice, ", "), " is filled")
				}
			case isUsernameTag:
				{
					errValidation[fieldName] = fmt.Sprint(fieldName, " username must contain only letters and numbers.")
				}
			default:
				{
					errValidation[fieldName] = fmt.Sprint("Error on tag ", fieldError.Tag(), " on field ", fieldName, " with error ", fieldError.Error())
				}
			}

		}
	}

	return errValidation
}

func IsValidExtension(allowedExtension []string, fileName string) bool {
	// Extract the file extension
	ext := strings.ToLower(filepath.Ext(fileName))
	for _, validExt := range allowedExtension {
		if ext == validExt {
			return true
		}
	}
	return false
}

func ValidateMultimediaFileExtension(files []entity.MultimediaFile, allowedTypes []string) error {
	for _, requestFile := range files {
		//Validate extension
		if !IsValidExtension(allowedTypes, requestFile.Filename) {
			return error_util.ErrorValidation(
				fiber.Map{
					"files": fmt.Sprintf(
						"invalid file type (%s). Allowed types are %s",
						requestFile.Filename,
						string_util.Implode(allowedTypes, ", "),
					),
				},
			)
		}
	}
	return nil
}

func isUsernameValid(fl validator.FieldLevel) bool {
	// Define regex for a valid username
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	// Return whether the field matches the regex
	return re.MatchString(fl.Field().String())
}

func ValidateStruct[T any](validator *validator.Validate, s T) (err error) {
	err = validator.Struct(s)
	if err != nil {
		errValidation := translateErrorMessage(err)
		return error_util.ErrorValidation(errValidation)
	}
	return
}
