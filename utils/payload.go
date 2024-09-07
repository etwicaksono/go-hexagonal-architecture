package utils

import (
	"github.com/gofiber/fiber/v2"
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
