package valueobject

import (
	"fmt"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/error_util"
	"github.com/etwicaksono/go-hexagonal-architecture/internal/utils/string_util"
	"github.com/gofiber/fiber/v2"
)

type MultimediaStorage int32

const (
	MultimediaStorage_LOCAL   MultimediaStorage = 1
	MultimediaStorage_MINIO   MultimediaStorage = 2
	MultimediaStorage_INVALID MultimediaStorage = 0
)

func (value MultimediaStorage) ToInt32() int32 {
	return int32(value)
}

func (value MultimediaStorage) ToString() string {
	switch value {
	case MultimediaStorage_LOCAL:
		return "LOCAL"
	case MultimediaStorage_MINIO:
		return "MINIO"
	default:
		return "INVALID"
	}
}

func MultimediaStorageFromInt32(value int32) MultimediaStorage {
	switch value {
	case 1:
		return MultimediaStorage_LOCAL
	case 2:
		return MultimediaStorage_MINIO
	default:
		return MultimediaStorage_INVALID
	}

}
func MultimediaStorageFromString(value string) MultimediaStorage {
	switch value {
	case "LOCAL":
		return MultimediaStorage_LOCAL
	case "MINIO":
		return MultimediaStorage_MINIO
	default:
		return MultimediaStorage_INVALID
	}
}
func ValidateMultimediaStorageString(value string) error {
	if MultimediaStorageFromString(value) == MultimediaStorage_INVALID {
		return error_util.ErrorValidation(
			fiber.Map{"storage": fmt.Sprintf(
				"Invalid storage type. Available types are: %s",
				string_util.Implode(
					[]string{
						MultimediaStorage_LOCAL.ToString(),
						MultimediaStorage_MINIO.ToString(),
					},
					", ",
				))},
		)
	}
	return nil
}
