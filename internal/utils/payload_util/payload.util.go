package payload_util

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/error_util"
	"github.com/gofiber/fiber/v2"
	"regexp"
	"strings"
)

func handleParsingError(err error) (errParsing fiber.Map, errOther error) {
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

func Slugify(input string) string {
	// Convert the input string to lowercase
	slug := strings.ToLower(input)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove all non-alphanumeric characters (except hyphens)
	re := regexp.MustCompile(`[^a-z0-9-]+`)
	slug = re.ReplaceAllString(slug, "")

	// Return the resulting slug
	return slug
}

func BodyParser[T any](ctx *fiber.Ctx, payload *T) (err error) {
	err = ctx.BodyParser(payload)
	if err != nil {
		errParsing, errOther := handleParsingError(err)
		if errOther != nil {
			return errOther
		}
		return error_util.ErrorValidation(errParsing)
	}
	return
}
