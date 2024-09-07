package utils

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"regexp"
	"strings"
)

func HandleParsingError(err error) (errParsing fiber.Map, errOther error) {
	if strings.HasPrefix(err.Error(), "failed to decode: schema: error converting value for") {
		// Compile the regex pattern
		regex := regexp.MustCompile(`failed to decode: schema: error converting value for "(.*?)"`)

		// Find the submatches using the regex
		matches := regex.FindStringSubmatch(err.Error())

		// Check if there is a match and print the captured value
		if len(matches) >= 2 {
			capturedValue := matches[1]
			return fiber.Map{capturedValue: err.Error()}, nil
		}
	}

	return nil, err
}

func GenerateErrorMessage(err error) (errValidation fiber.Map) {
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
			default:
				{
					errValidation[fieldName] = fmt.Sprint("Error on tag ", fieldError.Tag(), " on field ", fieldName, " with error ", fieldError.Error())
				}
			}

		}
	}

	return errValidation
}
