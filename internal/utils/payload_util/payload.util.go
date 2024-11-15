package payload_util

import (
	"context"
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/constants"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/error_util"
	"github.com/gofiber/fiber/v2"
	"io"
	"log/slog"
	"mime/multipart"
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

func MultipartFormParser(ctx *fiber.Ctx, fields ...string) (multimediaFilesMap map[string][]entity.MultimediaFile, err error) {
	// Parse the multipart form containing the files
	multimediaFilesMap = make(map[string][]entity.MultimediaFile)
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, error_util.ErrorValidation(fiber.Map{"files": "Failed to parse multipart form"})
	}

	for _, fld := range fields {
		if form.File[fld] == nil {
			return nil, error_util.ErrorValidation(fiber.Map{fld: "File is required"})
		}

		// Get the files from the field in the form
		files := form.File[fld]
		for _, file := range files {
			// Open the file
			openedFile, err := file.Open()
			if err != nil {
				slog.ErrorContext(ctx.UserContext(), fmt.Sprintf("Failed to open file (%s)", file.Filename), slog.String(constants.Error, err.Error()))
				return nil, error_util.ErrorValidation(fiber.Map{fld: fmt.Sprintf("Failed to open file (%s)", file.Filename)})
			}

			// Read file content into []byte
			fileBytes, err := io.ReadAll(openedFile)
			if err != nil {
				closeOpenedFile(ctx.UserContext(), openedFile, file.Filename)
				return nil, error_util.ErrorValidation(fiber.Map{"files": fmt.Sprintf("Failed to read file (%s)", file.Filename)})
			}

			multimediaFilesMap[fld] = append(multimediaFilesMap[fld], entity.MultimediaFile{
				Filename:    file.Filename,
				ContentType: file.Header.Get("Content-Type"),
				Data:        fileBytes,
			})
			closeOpenedFile(ctx.UserContext(), openedFile, file.Filename)
		}
	}

	return
}

func closeOpenedFile(ctx context.Context, openedFile multipart.File, filename string) {
	if err := openedFile.Close(); err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("Failed to close file (%s)", filename), slog.String(constants.Error, err.Error()))
	}

}
